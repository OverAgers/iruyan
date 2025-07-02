package usecases

import (
	"errors"
	"fmt"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/repositories"
	"iruyan-api/responses"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomUsecase interface {
	CreateRoom(roomName string) (*models.Room, error)
	EnterRoom(iruyanID string, roomID uuid.UUID, task string) (*models.WorkTime, *models.Room, error)
	GetRooms() ([]responses.RoomDetail, error)
	GetRoomByID(roomID uuid.UUID) (*models.Room, error)
	LeaveRoom(iruyanID string, roomID uuid.UUID, duration time.Duration) (*responses.LeaveRoomResponse, error)
	TakeSeat(iruyanID string, roomID uuid.UUID, seatNumber int) error
	LeaveSeat(iruyanID string, roomID uuid.UUID, seatNumber int) error
	GetSeatedUsers(roomID uuid.UUID) ([]responses.SeatStatusResponse, error)
	CreateRoomWithSeats(name string, seatCount int) error
}

type roomUsecase struct {
	UserRepo     repositories.UserRepositoryInterface
	RoomRepo     repositories.RoomRepositoryInterface
	SeatRepo     repositories.SeatRepositoryInterface
	WorkTimeRepo repositories.WorkTimeRepositoryInterface
	DB           *gorm.DB
}

func (u *roomUsecase) CreateRoomWithSeats(name string, seatCount int) error {
	// 既存ルーム確認
	exists, err := u.RoomRepo.ExistsByName(name)
	if err != nil {
		return fmt.Errorf("failed to check existing room: %w", err)
	}
	if exists {
		fmt.Printf("Room '%s' already exists. Skipping creation.\n", name)
		return nil
	}

	// トランザクション開始
	tx := u.DB.Begin()

	room := models.Room{Name: name}
	if err := room.ValidateName(); err != nil {
		tx.Rollback()
		return fmt.Errorf("room name invalid: %w", err)
	}

	if err := u.RoomRepo.CreateWithTx(tx, &room); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create room: %w", err)
	}

	for i := 0; i < seatCount; i++ {
		if _, err := u.SeatRepo.CreateSeat(tx, room.ID); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create seat: %w", err)
		}
	}

	return tx.Commit().Error
}

func (u *roomUsecase) CreateRoom(name string) (*models.Room, error) {
	room := &models.Room{Name: name}

	// 名前のバリデーション
	if err := room.ValidateName(); err != nil {
		return nil, err
	}

	// [💡]複数のリポジトリを横断するビジネス処理の場合は、usecaeでトランザクションを制御する
	tx := u.DB.Begin()

	// 部屋作成
	if err := u.RoomRepo.CreateWithTx(tx, room); err != nil {
		tx.Rollback()
		return nil, err
	}

	// シートを10個自動生成
	for i := 0; i < 10; i++ {
		if _, err := u.SeatRepo.CreateSeat(tx, room.ID); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to create seat: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return room, nil
}

func (u *roomUsecase) EnterRoom(iruyanID string, roomID uuid.UUID, task string) (*models.WorkTime, *models.Room, error) {
	// ユーザー取得
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, nil, errdefs.ErrUserNotFound
	}

	// ルーム取得
	room, err := u.RoomRepo.FindByID(roomID)
	if err != nil {
		return nil, nil, errdefs.ErrRoomNotFound
	}

	// すでに入室していないか確認
	exists, err := u.WorkTimeRepo.IsUserAlreadyInRoom(user.ID, roomID)
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errdefs.ErrAlreadyInRoom
	}

	// WorkTimeを作成
	workTime := &models.WorkTime{
		UserID:    user.ID,
		RoomID:    roomID,
		Task:      task,
		EntryTime: time.Now(),
	}

	if err := u.WorkTimeRepo.Create(workTime); err != nil {
		return nil, nil, err
	}

	return workTime, room, nil
}

func (u *roomUsecase) GetRoomByID(roomID uuid.UUID) (*models.Room, error) {
	room, err := u.RoomRepo.FindByIDWithSeats(roomID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (u *roomUsecase) LeaveRoom(iruyanID string, roomID uuid.UUID, duration time.Duration) (*responses.LeaveRoomResponse, error) {
	// ユーザー確認
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	// ルーム確認
	room, err := u.RoomRepo.FindByID(roomID)
	if err != nil {
		return nil, errdefs.ErrRoomNotFound
	}

	// 入室記録取得
	workTime, err := u.WorkTimeRepo.GetLatestEntry(user.ID, roomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdefs.ErrNotInRoom
		}
		return nil, err
	}

	leavingTime := time.Now()

	// 不正なduration判定
	if duration < leavingTime.Sub(workTime.EntryTime) {
		return nil, errdefs.ErrInvalidDuration
	}

	// 更新
	workTime.LeavingTime = leavingTime
	workTime.Duration = duration

	if err := u.WorkTimeRepo.Update(workTime); err != nil {
		return nil, err
	}

	return &responses.LeaveRoomResponse{
		RoomID:      room.ID.String(),
		RoomName:    room.Name,
		EntryTime:   workTime.EntryTime,
		LeavingTime: leavingTime,
		Duration:    duration,
	}, nil
}

func (u *roomUsecase) TakeSeat(iruyanID string, roomID uuid.UUID, seatNumber int) error {
	// ユーザー確認
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return errdefs.ErrUserNotFound
	}

	// 座席が部屋に存在するか
	exists, err := u.SeatRepo.ExistsInRoom(roomID, seatNumber)
	if err != nil {
		return err
	}
	if !exists {
		return errdefs.ErrSeatNotFound
	}

	// 座席がすでに使用されているか
	taken, err := u.WorkTimeRepo.IsSeatTaken(roomID, seatNumber)
	if err != nil {
		return err
	}
	if taken {
		return errdefs.ErrSeatAlreadyTaken
	}

	// 入室記録を取得
	workTime, err := u.WorkTimeRepo.GetLatestEntry(user.ID, roomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errdefs.ErrNotInRoom
		}
		return err
	}

	// 座席番号を更新
	if err := u.WorkTimeRepo.UpdateSeatNumber(workTime.ID, seatNumber); err != nil {
		return err
	}
	return nil
}

func (u *roomUsecase) LeaveSeat(iruyanID string, roomID uuid.UUID, seatNumber int) error {
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return errdefs.ErrUserNotFound
	}

	workTime, err := u.WorkTimeRepo.GetLatestEntry(user.ID, roomID)
	if err != nil {
		return err
	}

	if workTime.SeatNumber == 0 {
		return errdefs.ErrNotSeated
	}
	if workTime.SeatNumber != seatNumber {
		return fmt.Errorf("%w: current seat number is %d", errdefs.ErrSeatMismatch, workTime.SeatNumber)
	}

	return u.WorkTimeRepo.UpdateSeatNumber(workTime.ID, 0)
}

func (u *roomUsecase) GetSeatedUsers(roomID uuid.UUID) ([]responses.SeatStatusResponse, error) {
	workTimes, err := u.WorkTimeRepo.FindSeatedUsersByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	// 整形してレスポンスを生成
	response := make([]responses.SeatStatusResponse, 0, len(workTimes))
	for _, wt := range workTimes {
		response = append(response, responses.SeatStatusResponse{
			SeatNumber: wt.SeatNumber,
			IruyanID:   wt.User.IruyanID,
			UserName:   wt.User.Name,
			Email:      wt.User.Email,
			AvatarUrl:  "",
			Task:       wt.Task,
			Note:       "",
			Status:     "",
			StartTime:  wt.EntryTime.Unix(),
		})
	}
	return response, nil
}
