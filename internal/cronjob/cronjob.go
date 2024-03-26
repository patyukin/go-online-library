package cronjob

import (
	"context"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	c  *cron.Cron
	uc UseCase
}

type UseCase interface {
	GetAllPromotions(ctx context.Context) error
}

func NewCronJob(uc UseCase) *CronJob {
	return &CronJob{
		c:  cron.New(),
		uc: uc,
	}
}

func (cj *CronJob) Stop() {
	cj.c.Stop()
}

func (cj *CronJob) Run(ctx context.Context, errCh chan error) error {
	//_, err := cj.c.AddFunc(schedule, f)
	//err = cj.uc.GetAllPromotions(ctx)
	//if err != nil {
	//	logrus.Error(err)
	//	return err
	//}
	//cj.c.Start()

	/// Todo
	//err = uc.GetAllPromotions(ctx)
	//if err != nil {
	//	logrus.Errorf("error occured while adding cron job: %v", err)
	//}
	//return err

	return nil
}
