package controllers

import (
	"Hackathon/backend/domain"
	"encoding/json"
	"io"
	"net/http"
)

type HabitController struct {
	habitRepository domain.HabitRepository
}

func NewHabitController(habitRepository domain.HabitRepository) *HabitController {
	return &HabitController{habitRepository}
}

func (controller HabitController) AddHabit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var habit domain.Habit

	err = json.Unmarshal(body, &habit)
	if err != nil {
		http.Error(w, "Error parsing request Body", http.StatusBadRequest)
		return
	}

	err = controller.habitRepository.NewHabit(habit)
	if err != nil {
		http.Error(w, "Error saving to Database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Habit added successfully"})
	if err != nil {
		return
	}
}
