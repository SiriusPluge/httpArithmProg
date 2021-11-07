package internal

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func TaskCompletion(wg *sync.WaitGroup, i uint) {

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	for {
		select {
		case cancel := <-cancelChan:
			log.Printf("cancelChan received: %v, goroutine %v end", cancel, i)
			wg.Done()
			return
		case tasks := <-ChanTask:
			for _, repTask := range ReposTask {
				if repTask.ExecStatus == Progress && repTask.IdQueue != 0 {
					repTask.IdQueue--
				}

				log.Printf("worker %v starts processing", tasks.Id)
				tasks.StartTime = time.Now().Format(time.Stamp)
				tasks.ExecStatus = Queued

				for i := 0; i < tasks.N; i++ {
					time.Sleep(time.Millisecond * time.Duration(tasks.D*100))
					tasks.N1 += tasks.D
					tasks.I++
				}

				tasks.I = 0
				tasks.EndTime = time.Now().Format(time.Stamp)
				tasks.ExecStatus = Completed

				log.Printf("worker %v finished", tasks.Id)

				wg.Add(1)
				go func(tasks *Task) {
					time.AfterFunc(time.Millisecond*time.Duration(tasks.TTL*1000), func() {
						DeleteTask(tasks)
						log.Printf("worker %v deleted", tasks.Id)
						wg.Done()
					})
				}(tasks)
			}
		}
	}
}
