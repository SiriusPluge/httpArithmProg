package internal

import "sync"

func DeleteTask(t *Task) {
	mu := sync.Mutex{}

	for indx, item := range ReposTask {
		if item == t {
			mu.Lock()
			ReposTask[indx], ReposTask[len(ReposTask)-1] = ReposTask[len(ReposTask)-1], ReposTask[indx]
			ReposTask[len(ReposTask)-1] = nil
			ReposTask = ReposTask[:len(ReposTask)-1]
			mu.Unlock()

			return
		}
	}
}
