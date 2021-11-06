package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	Progress = "Progress"
	Queued = "Queued"
	Completed = "Comleted"
)

var ReposTask []*Task
var ChanTask = make(chan *Task, 1)
var StartTuskId = 0

func PutTask(w http.ResponseWriter, req *http.Request) {

	log.Printf("handling put a task in a queue at %s\n", req.URL.Path)

	w.Header().Set("Content-Type", "application/json")

	var tasks Task
	err := json.NewDecoder(req.Body).Decode(&tasks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks.StatementTime = time.Now().Format(time.RFC822)
	tasks.ExecStatus = Progress

	StartTuskId++
	fmt.Println(StartTuskId)
	tasks.Id = StartTuskId
	log.Printf("task %v added", tasks.Id)

	ReposTask = append(ReposTask, &tasks)

	http.Redirect(w, req, "/add", http.StatusOK)

	var queueCounter int // updating the numInQueue
	for _, repTask := range ReposTask {
		if repTask.ExecStatus == Progress {
			repTask.IdQueue = queueCounter
			queueCounter++
		}
		if repTask.ExecStatus == Progress && repTask.IdQueue == 0 {
			// log.Println("w = ", w, "worker = ", worker)
			go func(tasks *Task) {
				log.Printf("worker %v sending to chan\n", tasks.Id)
				ChanTask <- tasks
			}(&tasks)
		}
	}
	//Поставить задачу в очередь. Параметры:
	//n - количество элементов (целочисленное)
	//d - дельта между элементами последовательности (вещественное)
	//n1 - Стартовое значение (вещественное)
	//I - интервал в секундах между итерациями (вещественное)
	//TTL - время хранения результата в секундах (вещественное)
}

func GetListAndStatus(w http.ResponseWriter, req *http.Request) {

	log.Printf("handling get a sorted list and the statuses at %s\n", req.URL.Path)

	//Получить отсортированный список задач и статусы выполнения этих задач. Поля результата для каждой задачи:
	//Номер в очереди (целочисленное)
	//Статус: В процессе/В очереди/Завершена
	//n
	//d
	//n1
	//I
	//TTL
	//Текущая итерация
	//Время постановки задачи
	//Время старта задачи
	//Время окончания задачи (в случае если задача завершена)
}
