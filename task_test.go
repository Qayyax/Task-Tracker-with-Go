package main

import (
	"fmt"
	"testing"

	"github.com/Qayyax/Task-Tracker-with-Go/types"
)

func TestGetTaskIndexById(t *testing.T) {
	index := types.GetTaskIndexById(2)
	fmt.Println("The index is:", index)
}

func TestUpdateTask(t *testing.T) {
	types.UpdateTask(1, "New Day 1", "Nobara is looking at me")
}

func TestDeleteTask(t *testing.T) {
	types.DeleteTask(2)
}
