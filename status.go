package jobrunner

import (
	"time"

	"github.com/robfig/cron/v3"
)

type StatusData struct {
	Id        cron.EntryID
	JobRunner *Job
	Next      time.Time
	Prev      time.Time
}

// Return detailed list of currently running recurring jobs
// to remove an entry, first retrieve the ID of entry
func Entries() []cron.Entry {
	return MainCron.Entries()
}

func StatusPage() []StatusData {
	ents := MainCron.Entries()
	statuses := make([]StatusData, len(ents))
	
	for k, v := range ents {
		statuses[k].Id = v.ID
		statuses[k].JobRunner = AddJob(v.Job)
		statuses[k].Next = v.Next
		statuses[k].Prev = v.Prev
	}
	return statuses
}

func StatusJson() map[string]interface{} {

	return map[string]interface{}{
		"jobrunner": StatusPage(),
	}

}

func AddJob(job cron.Job) *Job {
	return job.(*Job)
}
