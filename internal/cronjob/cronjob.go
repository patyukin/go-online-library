package cronjob

import "github.com/robfig/cron/v3"

type CronJob struct {
	c *cron.Cron
}

func NewCronJob() *CronJob {
	return &CronJob{
		c: cron.New(),
	}
}

func (cj *CronJob) Start() {
	cj.c.Start()
}

func (cj *CronJob) Stop() {
	cj.c.Stop()
}

func (cj *CronJob) AddFunc(schedule string, f func()) error {
	_, err := cj.c.AddFunc(schedule, f)
	return err
}
