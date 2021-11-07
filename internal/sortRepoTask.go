package internal

import "sort"

func sortRepoTask() {
	sortTasks := map[string]int{
		Progress:  0,
		Queued:    1,
		Completed: 2,
	}
	sort.Slice(ReposTask, func(i, j int) bool {
		if ReposTask[i].ExecStatus == ReposTask[j].ExecStatus {
			return ReposTask[i].IdQueue > ReposTask[j].IdQueue
		}
		return sortTasks[ReposTask[i].ExecStatus] < sortTasks[ReposTask[j].ExecStatus]
	})
}