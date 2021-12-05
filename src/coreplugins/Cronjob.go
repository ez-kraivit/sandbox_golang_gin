package coreplugins

import (
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func CronJobFunc(){
    log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { 
        log.Info("[Job 1]Every minute job\n")
    })
	log.Info("Start cron")
	c.Start()
	CronJobprintCronEntries(c.Entries())
	time.Sleep(2 * time.Minute)
}

func CronJobprintCronEntries(cronEntries []*cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}