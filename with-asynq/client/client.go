package main

import (
	"github.com/hibiken/asynq"
	"log"
	"quickstart/task"
	"time"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	t1, err := task.NewWelcomeEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	// Process the task immediately.
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)

	//for i := 3; i < 7; i++ {
	//	info, err = client.Enqueue(t2, asynq.ProcessIn(time.Duration(i)*time.Second))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Printf(" [*] Successfully enqueued task: %+v", info)
	//}

	t2, err := task.NewReminderEmailTask(555)
	if err != nil {
		log.Fatal(err)
	}
	info, err = client.Enqueue(t2, asynq.ProcessIn(4*time.Second))

	t2_2, err := task.NewReminderEmailTask_2(666)
	if err != nil {
		log.Fatal(err)
	}
	info, err = client.Enqueue(t2_2, asynq.ProcessIn(6*time.Second))

	t2_3, err := task.NewReminderEmailTask_3(777)
	if err != nil {
		log.Fatal(err)
	}
	info, err = client.Enqueue(t2_3, asynq.ProcessIn(8*time.Second))

	t2_4, err := task.NewReminderEmailTask_4(888)
	if err != nil {
		log.Fatal(err)
	}
	info, err = client.Enqueue(t2_4, asynq.ProcessIn(10*time.Second))
}
