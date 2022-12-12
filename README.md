# ![](https://raw.githubusercontent.com/magicLian/jobrunner/master/views/runclock.jpg) JobRunner

##### This Repo is forked from ```github.com/bamzi/jobrunner```

※This is an enhanced version※

JobRunner is framework for performing work asynchronously, outside of the request flow. It comes with cron to schedule and queue job functions for processing at specified time. 

It includes a live monitoring of current schedule and state of active jobs that can be outputed as JSON and it support get job execution history by go channel. 

## Install

`go get github.com/magicLian/jobrunner`

### Supported Featured
*All jobs are processed outside of the request flow*

* Support time location in scheduling job
* Support user-defined job name
* Support task status real-time query 
* Support to obtain task execution history records through Go `channel`

### Live Monitoring
![](https://raw.githubusercontent.com/magicLian/jobrunner/master/views/jobrunner-html.png)

## Basics

```go
    jobrunner.Schedule("* */5 * * * *", DoSomething{}, "DoSomethingJob") // every 5min do something
    jobrunner.Schedule("@every 1h30m10s", ReminderEmails{}, "ReminderEmailsJob")
    jobrunner.Schedule("@midnight", DataStats{}, "DataStatsJob") // every midnight do this..
    jobrunner.Every(16*time.Minute, CleanS3{}, "CleanS3Job") // evey 16 min clean...
    jobrunner.In(10*time.Second, WelcomeEmail{}, "WelcomeEmailJob") // one time job. starts after 10sec
    jobrunner.Now(NowDo{}, "NowJob") // do the job as soon as it's triggered
```
[**More Detailed CRON Specs**](https://github.com/robfig/cron/blob/v2/doc.go)

### Setup

#### standalone
```go
package main

import "github.com/magicLian/jobrunner"

func main() {
    if err := os.Setenv("ZONEINFO", "/tzdata/data.zip"); err != nil {
	panic(err)
    }
    loc,err := time.loadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }
    jobrunner.Start(loc, false)
    jobrunner.Schedule("@every 5s", ReminderEmails{}, "reminderEmails")
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
```

#### Integrate with Gin
```go

package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magicLian/jobrunner"
)

func main() {
	routes := gin.Default()

	// Resource to return the JSON data
	routes.GET("/jobrunner/json", JobJson)

	jobrunner.Start(nil, false)
	jobrunner.Every(10*time.Minute, DoSomeThing{}, "DoSomeThing")

	routes.Run(":8080")
}

type DoSomeThing struct {
}

func (d DoSomeThing) Run() {
	fmt.Printf("Start to do something")
}

func JobJson(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

```

#### Get job execution result through go channel
```go

func main() {
	...
	
	jobrunner.Start(nil, true)
	go getCronExecutionResults()
	jobrunner.Every(10*time.Minute, DoSomeThing{}, "DoSomeThing")

	


	...
}

func getCronExecutionResults(){
	for {
		select {
		case record := <-jobrunner.JobsExecutionStatusChan:
			fmt.Printf("Job exection status result:[%v]", record)
		}
	}
}

```

## Credits
- [revel jobs module](https://github.com/revel/modules/tree/master/jobs) - Origin of JobRunner
- [robfig cron v3](https://github.com/robfig/cron/tree/v3) - github.com/robfig/cron/v3

#### License
MIT
