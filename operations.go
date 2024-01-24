package main

import(
	"log"
	"github.com/couchbase/gocb/v2"
)


// Insert/Upsert a document
func insertDocument(col *gocb.Collection,testKey string) {
	
	testDoc := doc{
		Key1:"value1",
		Key2:"value2",
		Key3:"value3",
	}

	_, err := col.Upsert(testKey,testDoc, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Update a document

// Deleta a document


// creating new WorkerPool
func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func()),
	}
	return wp
}

func (wp *workerPool) Run(){
	for i := 0 ; i < wp.maxWorker ;i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}


func (wp *workerPool) AddTask(task func()){
	wp.queuedTaskC <- task
}