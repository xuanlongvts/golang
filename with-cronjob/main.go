package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	_, _ = c.AddFunc("0 */1 * * * *", func() { // every 1 minute
		fmt.Println("aaa")
	})
	c.Start()
	time.Sleep(5 * time.Minute)
}
