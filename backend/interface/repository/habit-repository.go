package repository

import "Hackathon/backend/domain"

type HabitRepository struct {
	handler DatabaseHandler
}

func NewHabitRepository(handler DatabaseHandler) *HabitRepository {
	return &HabitRepository{handler}
}

func (repo HabitRepository) NewHabit(habit domain.Habit) error {
	return repo.handler.SaveHabit(habit)
}

func (repo HabitRepository) ChangeHabit(id, listId int, habit domain.Habit) error {
	return repo.handler.ChangeHabit(id, listId, habit)
}

func (repo HabitRepository) LoadHabit(id, listId int) (*domain.Habit, error) {
	return repo.handler.LoadHabit(id, listId)
}

func (repo HabitRepository) LoadHabitList(listId int) error {
	return repo.handler.LoadHabitList(listId)
}

func (repo HabitRepository) LoadAllHabit() (*[]domain.Habit, error) {
	return repo.handler.LoadAllHabit()
}
