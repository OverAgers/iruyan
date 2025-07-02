package repositories

import (
	"iruyan-api/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkTimeRepositoryInterface interface {
	IsUserAlreadyInRoom(userID uint, roomID uuid.UUID) (bool, error)
	Create(workTime *models.WorkTime) error
	GetLatestEntry(userID uint, roomID uuid.UUID) (*models.WorkTime, error)
	Update(workTime *models.WorkTime) error
	IsSeatTaken(roomID uuid.UUID, seatNumber int) (bool, error)
	UpdateSeatNumber(workTimeID int, seatNumber int) error
	FindSeatedUsersByRoomID(roomID uuid.UUID) ([]models.WorkTime, error)
	GetLogForLastWeek(userID uint, iruyanID string) ([]models.WorkTime, error)
	GetActiveEntry(userID uint) (*models.WorkTime, error)
	GetLatestEntriesByUser(userID uint, limit int) ([]models.WorkTime, error)
	FindLatestEntry(userID uint, roomID uuid.UUID) (*models.WorkTime, error)
	GetLogsSince(userID uint, since time.Time) ([]models.WorkTime, error)
	FindAllWorkTimes() ([]models.WorkTime, error)
	GetByUserID(userID uint) ([]models.WorkTime, error)
}

type workTimeRepository struct {
	DB *gorm.DB
}

func NewWorkTimeRepository(db *gorm.DB) WorkTimeRepositoryInterface {
	return &workTimeRepository{
		DB: db,
	}
}

func (r *workTimeRepository) IsUserAlreadyInRoom(userID uint, roomID uuid.UUID) (bool, error) {
	var count int64
	if err := r.DB.Model(&models.WorkTime{}).
		Where("user_id = ? AND room_id = ? AND exit_time IS NULL", userID, roomID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *workTimeRepository) Create(workTime *models.WorkTime) error {
	return r.DB.Create(workTime).Error
}

func (r *workTimeRepository) GetLatestEntry(userID uint, roomID uuid.UUID) (*models.WorkTime, error) {
	var workTime models.WorkTime
	err := r.DB.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		Order("entry_time desc").
		First(&workTime).Error

	return &workTime, err
}

func (r *workTimeRepository) Update(workTime *models.WorkTime) error {
	return r.DB.Save(workTime).Error
}

func (r *workTimeRepository) IsSeatTaken(roomID uuid.UUID, seatNumber int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.WorkTime{}).
		Where("room_id = ? AND seat_number = ? AND leaving_time IS NULL", roomID, seatNumber).
		Count(&count).Error
	return count > 0, err
}

func (r *workTimeRepository) UpdateSeatNumber(workTimeID int, seatNumber int) error {
	return r.DB.Model(&models.WorkTime{}).
		Where("id = ?", workTimeID).
		Update("seat_number", seatNumber).Error
}

func (r *workTimeRepository) FindSeatedUsersByRoomID(roomID uuid.UUID) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	err := r.DB.
		Preload("User").
		Where("room_id = ? AND seat_number > 0", roomID).
		Find(&workTimes).Error
	if err != nil {
		return nil, err
	}
	return workTimes, nil
}

func (r *workTimeRepository) GetLogForLastWeek(userID uint, iruyanID string) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	fiveDaysAgo := time.Now().AddDate(0, 0, -5)

	err := r.DB.
		Where("user_id = ? AND entry_time >= ?", userID, fiveDaysAgo).
		Order("entry_time DESC").
		Find(&workTimes).Error

	if err != nil {
		return nil, err
	}

	return workTimes, nil
}

func (r *workTimeRepository) GetActiveEntry(userID uint) (*models.WorkTime, error) {
	var workTime models.WorkTime
	if err := r.DB.Where("user_id = ? AND leaving_time IS NULL", userID).First(&workTime).Error; err != nil {
		return nil, err
	}
	return &workTime, nil
}

func (r *workTimeRepository) GetLatestEntriesByUser(userID uint, limit int) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	err := r.DB.Where("user_id = ?", userID).
		Order("entry_time DESC").
		Limit(limit).
		Find(&workTimes).Error

	if err != nil {
		return nil, err
	}
	if len(workTimes) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return workTimes, nil
}

func (r *workTimeRepository) FindLatestEntry(userID uint, roomID uuid.UUID) (*models.WorkTime, error) {
	var workTime models.WorkTime
	err := r.DB.
		Where("user_id = ? AND room_id = ? AND leaving_time IS NULL", userID, roomID).
		Order("entry_time DESC").
		First(&workTime).Error

	if err != nil {
		return nil, err
	}
	return &workTime, nil
}

func (r *workTimeRepository) GetLogsSince(userID uint, since time.Time) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	if err := r.DB.
		Where("user_id = ? AND entry_time >= ?", userID, since).
		Order("entry_time desc").
		Find(&workTimes).Error; err != nil {
		return nil, err
	}
	return workTimes, nil
}

func (r *workTimeRepository) FindAllWorkTimes() ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	err := r.DB.
		Preload("User").
		Preload("Room").
		Find(&workTimes).Error

	if err != nil {
		return nil, err
	}
	return workTimes, nil
}

func (r *workTimeRepository) GetByUserID(userID uint) ([]models.WorkTime, error) {
	var workTimes []models.WorkTime
	if err := r.DB.
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Room").
		Find(&workTimes).Error; err != nil {
		return nil, err
	}
	return workTimes, nil
}
