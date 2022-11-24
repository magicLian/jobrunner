package examples

import (
	"fmt"
	"time"

	"github.com/magiclian/jobrunner"
	"github.com/sirupsen/logrus"
)

func StartWithLocation() {
	loc, err := time.LoadLocation("Asia/Tykyo")
	if err != nil {
		logrus.Fatalf("init loc failed:[%s]", err.Error())
	}

	jobrunner.Start(loc)
	jobrunner.Schedule("@every 5s", ReminderEmails{})
}

// Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}
