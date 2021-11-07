package internal

import (
	"fmt"
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

				log.Printf("task %v starts processing", tasks.Id)
				tasks.StartTime = time.Now().Format(time.Stamp)
				tasks.ExecStatus = Queued

				ArithmProg := tasks.N1 + (float64(tasks.N)-1.0)*tasks.D
				fmt.Printf("\nАрифметическая прогрессия по таску: %v \n", ArithmProg)

				for i := 0; i < tasks.N; i++ {
					time.Sleep(time.Millisecond * time.Duration(tasks.I*100))
					tasks.N1 += tasks.D
					tasks.CurIteration ++
				}

				tasks.EndTime = time.Now().Format(time.Stamp)
				tasks.ExecStatus = Completed

				log.Printf("task %v finished", tasks.Id)

				wg.Add(1)
				go func(tasks *Task) {
					time.AfterFunc(time.Millisecond*time.Duration(tasks.TTL*1000), func() {
						DeleteTask(tasks)
						log.Printf("task %v deleted", tasks.Id)
						wg.Done()
					})
				}(tasks)
			}
		}
	}
}
