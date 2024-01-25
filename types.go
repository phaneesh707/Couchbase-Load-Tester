package main


type T = interface{}

type doc struct {
	Key1 string   `json:"key1"`
	Key2 string   `json:"key2"`
	Key3 string   `json:"key3"`
	Key4 string   `json:"key4"`
	Key5 string   `json:"key5"`
	Key6 string   `json:"key6"`
	Key7 string   `json:"key7"`
	Key8 string   `json:"key8"`
	Key9 string   `json:"key9"`
	Key0 string   `json:"key0"`
}

type WorkerPool interface{
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker int
	queuedTaskC chan func()
}
