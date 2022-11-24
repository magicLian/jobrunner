package examples

import (
	"fmt"
	"time"

	"github.com/magicLian/jobrunner"
	"github.com/sirupsen/logrus"
)

func getJobExecutionStatus() {
	loc, err := time.LoadLocation("Asia/Tykyo")
	if err != nil {
		logrus.Fatalf("init loc failed:[%s]", err.Error())
	}

	jobrunner.Start(loc)
	jobrunner.Schedule("@every 5s", JobFunc1{}, "JobFunc1")

	for {
		select {
		case record := <-jobrunner.JobsExecutionStatusChan:
			fmt.Printf("record:[%v]\n", record)
		}
	}
}

// Job Specific Functions
type JobFunc1 struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e JobFunc1) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}
