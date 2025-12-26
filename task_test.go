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
	types.UpdateTask(1, "updated to 1", "updated the description")
}
