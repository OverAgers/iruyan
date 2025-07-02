package usecases

import (
	"errors"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/repositories"
	"iruyan-api/responses"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkTimeUsecase interface {
	GetWorkLogForLastWeek(iruyanID string) (responses.WorkLogForLastWeekResponse, error)
	UpdateTask(iruyanID, task string) error
	GetRecentLogs(iruyanID string, limit int) ([]responses.WorkTimeLog, error)
	EnterRoom(iruyanID string, roomID uuid.UUID, task string) (*models.WorkTime, error)
	GetLatestEntry(iruyanID string, roomID uuid.UUID) (*models.WorkTime, error)
	GetWeeklyLogs(iruyanID string) ([]responses.WorkTimeResponse, error)
	GetAllWorkTimes() ([]models.WorkTime, error)
	GetWorkTimeByIruyanID(iruyanID string) ([]models.WorkTime, error)
}

type workTimeUsecase struct {
	UserRepo     repositories.UserRepositoryInterface
	WorkTimeRepo repositories.WorkTimeRepositoryInterface
}

// [DI] UserRepositoryInterface
func NewWorkTimeUsecase(userRepo repositories.UserRepositoryInterface, workTimeRepo repositories.WorkTimeRepositoryInterface) WorkTimeUsecase {
	return &workTimeUsecase{
		UserRepo:     userRepo,
		WorkTimeRepo: workTimeRepo,
	}
}

func (u *workTimeUsecase) GetWorkLogForLastWeek(iruyanID string) (responses.WorkLogForLastWeekResponse, error) {
	// 直近1週間の作業ログ取得
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		if errors.Is(err, errdefs.ErrUserNotFound) {
			return responses.WorkLogForLastWeekResponse{}, errdefs.ErrUserNotFound
		}
		return responses.WorkLogForLastWeekResponse{}, err
	}

	workLogs, err := u.WorkTimeRepo.GetLogForLastWeek(user.ID, iruyanID)
	if err != nil {
		return responses.WorkLogForLastWeekResponse{}, err
	}

	// 日ごとの合計時間に集計
	dailyWorkHours := make(map[string]time.Duration)
	for _, log := range workLogs {
		dateStr := log.EntryTime.Format("2006-01-02")
		dailyWorkHours[dateStr] += log.Duration
	}

	dailyLogs := []responses.DailyWorkLogResponse{}
	for date, hours := range dailyWorkHours {
		dailyLogs = append(dailyLogs, responses.DailyWorkLogResponse{
			Date:  date,
			Hours: hours,
		})
	}

	return responses.WorkLogForLastWeekResponse{
		IruyanID:  iruyanID,
		DailyLogs: dailyLogs,
	}, nil
}

func (u *workTimeUsecase) UpdateTask(iruyanID, task string) error {
	// ユーザー存在確認
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return errdefs.ErrUserNotFound
	}

	// 入室中の WorkTime を取得
	workTime, err := u.WorkTimeRepo.GetActiveEntry(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errdefs.ErrNotInRoom
		}
		return err
	}

	// Task を更新
	workTime.Task = task
	if err := u.WorkTimeRepo.Update(workTime); err != nil {
		return err
	}
	return nil
}

func (u *workTimeUsecase) GetRecentLogs(iruyanID string, limit int) ([]responses.WorkTimeLog, error) {
	// ユーザーが存在するか確認
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	// 最新のWorkTimeログ取得
	records, err := u.WorkTimeRepo.GetLatestEntriesByUser(user.ID, limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdefs.ErrNotInRoom
		}
		return nil, err
	}

	// レスポンス整形
	result := make([]responses.WorkTimeLog, len(records))
	for i, w := range records {
		result[i] = responses.WorkTimeLog{
			EntryTime:   w.EntryTime,
			LeavingTime: w.LeavingTime,
			Duration:    w.Duration,
		}
	}

	return result, nil
}

func (u *workTimeUsecase) EnterRoom(iruyanID string, roomID uuid.UUID, task string) (*models.WorkTime, error) {
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	inRoom, err := u.WorkTimeRepo.IsUserAlreadyInRoom(user.ID, roomID)
	if err != nil {
		return nil, err
	}
	if inRoom {
		return nil, errdefs.ErrAlreadyInRoom
	}

	workTime := &models.WorkTime{
		UserID:      user.ID,
		RoomID:      roomID,
		Task:        task,
		EntryTime:   time.Now(),
		LeavingTime: time.Time{},
		Duration:    0,
	}
	if err := u.WorkTimeRepo.Create(workTime); err != nil {
		return nil, err
	}
	return workTime, nil
}

func (u *workTimeUsecase) GetLatestEntry(iruyanID string, roomID uuid.UUID) (*models.WorkTime, error) {
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	workTime, err := u.WorkTimeRepo.FindLatestEntry(user.ID, roomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdefs.ErrNoActiveSession
		}
		return nil, err
	}
	workTime.User = *user
	return workTime, nil
}

func (u *workTimeUsecase) GetWeeklyLogs(iruyanID string) ([]responses.WorkTimeResponse, error) {
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	workTimes, err := u.WorkTimeRepo.GetLogsSince(user.ID, time.Now().AddDate(0, 0, -6))
	if err != nil {
		return nil, err
	}

	var result []responses.WorkTimeResponse
	for _, wt := range workTimes {
		result = append(result, responses.WorkTimeResponse{
			IruyanID:    user.IruyanID,
			RoomID:      wt.RoomID,
			Task:        wt.Task,
			EntryTime:   wt.EntryTime,
			LeavingTime: wt.LeavingTime,
			Duration:    wt.Duration,
		})
	}
	return result, nil
}

func (u *workTimeUsecase) GetAllWorkTimes() ([]models.WorkTime, error) {
	return u.WorkTimeRepo.FindAllWorkTimes()
}

func (u *workTimeUsecase) GetWorkTimeByIruyanID(iruyanID string) ([]models.WorkTime, error) {
	user, err := u.UserRepo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, errdefs.ErrUserNotFound
	}

	workTimes, err := u.WorkTimeRepo.GetByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	if len(workTimes) == 0 {
		return nil, errdefs.ErrNotInRoom
	}

	return workTimes, nil
}
