package main

import (
	"github.com/hibiken/asynq"
	"log"
	"quickstart/task"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)

	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)
	mux.HandleFunc(task.TypeReminderEmail_2, task.HandleReminderEmailTask_2)
	mux.HandleFunc(task.TypeReminderEmail_3, task.HandleReminderEmailTask_3)
	mux.HandleFunc(task.TypeReminderEmail_4, task.HandleReminderEmailTask_4)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
