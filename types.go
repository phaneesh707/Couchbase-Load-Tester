package main


type T = interface{}

type doc struct {
	Key1 string   `json:"key1"`
	Key2 string   `json:"key2"`
	Key3 string   `json:"key3"`
}

type WorkerPool interface{
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker int
	queuedTaskC chan func()
}
