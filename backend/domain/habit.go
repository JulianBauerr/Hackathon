package domain

import "time"

type Habit struct {
	Name     string    `json:"name"`
	Deadline time.Time `json:"deadline"`
	Status   Status    `json:"status"`
	Id       int       `json:"id"`
	ListId   int       `json:"listid"`
}

type HabitRepository interface {
	NewHabit(habit Habit) error
	ChangeHabit(id, listId int, habit Habit) error
	LoadHabit(id, listId int) (*Habit, error)
	LoadHabitList(listId int) error
	LoadAllHabit() (*[]Habit, error)
}
