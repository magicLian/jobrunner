package examples

import (
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
}
