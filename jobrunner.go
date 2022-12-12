package jobrunner

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

const UNNAMED = "(unnamed)"

type Job struct {
	Name    string
	inner   cron.Job
	status  uint32
	Status  string
	Latency string
	running sync.Mutex
}

type JobStatus struct {
	Name      string
	Status    string
	StartTime time.Time
	EndTime   time.Time
}

func (js *JobStatus) String() {
	fmt.Printf("[Name: %s, createTime: %s, endTime: %s\n]", js.Name, js.StartTime.Format(time.RFC3339), js.EndTime.Format(time.RFC3339))
}

func New(job cron.Job, n string) *Job {
	name := UNNAMED
	if n != "" {
		name = n
	} else {
		name := reflect.TypeOf(job).Name()
		if name == "Func" {
			name = UNNAMED
		}
	}

	return &Job{
		Name:  name,
		inner: job,
	}
}

func (j *Job) StatusUpdate() string {
	if atomic.LoadUint32(&j.status) > 0 {
		j.Status = "RUNNING"
		return j.Status
	}
	j.Status = "IDLE"
	return j.Status

}

func (j *Job) Run() {
	start := time.Now()
	// If the job panics, just print a stack trace.
	// Don't let the whole process die.
	defer func() {
		if err := recover(); err != nil {
			var buf bytes.Buffer
			logger := log.New(&buf, "JobRunner Log: ", log.Lshortfile)
			logger.Panic(err, "\n", string(debug.Stack()))
		}
	}()

	if !selfConcurrent {
		j.running.Lock()
		defer j.running.Unlock()
	}

	if workPermits != nil {
		workPermits <- struct{}{}
		defer func() { <-workPermits }()
	}

	atomic.StoreUint32(&j.status, 1)
	j.StatusUpdate()

	defer j.StatusUpdate()
	defer atomic.StoreUint32(&j.status, 0)

	j.inner.Run()

	end := time.Now()
	j.Latency = end.Sub(start).String()

	if IsStoreExecutionStatus {
		JobsExecutionStatusChan <- &JobStatus{
			Name:      j.Name,
			StartTime: start,
			EndTime:   end,
		}
	}
}
