package database

import (
	"Hackathon/backend/domain"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
	"time"
)

type DBHandler struct {
	SqlDB  *sql.DB
	GormDB *gorm.DB
}

var testData = []domain.Habit{
	{
		Id:       1,
		Name:     "Exercise",
		Deadline: time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
		Status:   1,
		ListId:   101,
	},
	{
		Id:       2,
		Name:     "Read a book",
		Deadline: time.Date(2023, 11, 30, 20, 0, 0, 0, time.UTC),
		Status:   0,
		ListId:   102,
	},
	{
		Id:       3,
		Name:     "Meditate",
		Deadline: time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC),
		Status:   1,
		ListId:   103,
	},
}

func NewDBHandler(connectString string) (DBHandler, error) {
	return DBHandler{}, nil
	sqlDB, err := sql.Open("pgx", connectString)
	if err != nil {
		log.Fatal(err)
		return DBHandler{}, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return DBHandler{}, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return DBHandler{}, err
	}

	err = gormDB.AutoMigrate(&domain.Habit{})
	if err != nil {
		return DBHandler{}, err
	}

	return DBHandler{sqlDB, gormDB}, nil
}

func (dbHandler DBHandler) SaveHabit(habit domain.Habit) error {
	fmt.Printf("%+v\n", habit)
	return nil
	if err := dbHandler.GormDB.Find(&domain.Habit{}, "id = ? AND listid = ?", habit.Id, habit.ListId).Error; err != nil {
		return fmt.Errorf("failed to find habit with id:%d listid:%d: %w", habit.Id, habit.ListId, err)
	}

	if err := dbHandler.GormDB.Save(&habit).Error; err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%s: Added a Habit\t\t", time.Now().Format(time.TimeOnly)), habit)
	return nil
}

func (dbHandler DBHandler) ChangeHabit(id, listId int, habit domain.Habit) error {
	//TODO implement me
	panic("implement me")
}

func (dbHandler DBHandler) LoadHabit(id, listId int) (*domain.Habit, error) {
	//TODO implement me
	panic("implement me")
}

func (dbHandler DBHandler) LoadHabitList(listId int) error {
	//TODO implement me
	panic("implement me")
}

func (dbHandler DBHandler) LoadAllHabit() (*[]domain.Habit, error) {
	return &testData, nil
}
