package main

import "time"

// - `id`: A unique identifier for the task
// - `description`: A short description of the task
// - `status`: The status of the task (`todo`, `in-progress`, `done`)
// - `createdAt`: The date and time when the task was created
// - `updatedAt`: The date and time when the task was last updated

type Task struct {
	id          int
	description string
	status      string // todo, in-progress, done
	createdAt   time.Time
	updatedAt   time.Time
}

// TODO:
// - [] methods for Task
