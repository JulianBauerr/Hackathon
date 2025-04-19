package repository

import "Hackathon/backend/domain"

type DatabaseHandler interface {
	SaveHabit(habit domain.Habit) error
	ChangeHabit(id, listId int, habit domain.Habit) error
	LoadHabit(id, listId int) (*domain.Habit, error)
	LoadHabitList(listId int) error
	LoadAllHabit() (*[]domain.Habit, error)
}
