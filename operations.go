package main

import (
	"fmt"
	// "time"

	"github.com/couchbase/gocb/v2"
)

// Insert/Upsert a document
func insertDocument(col *gocb.Collection,testKey string) {
	
	testDoc := doc{
		Key1:"value1",
		Key2:"value2",
		Key3:"value3",
		Key4:"value4",
		Key5:"value5",
		Key6:"value6",
		Key7:"value7",
		Key8:"value8",
		Key9:"value9",
		Key0:"value0",
	}

	_, err := col.Insert(testKey,testDoc, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// update document
func updateDocument(col *gocb.Collection,testKey string) {
	
	testDoc := doc{
		Key1:"value1-modified",
		Key2:"value2-modified",
		Key3:"value3-modified",
		Key4:"value4",
		Key5:"value5",
		Key6:"value6",
		Key7:"value7",
		Key8:"value8",
		Key9:"value9",
		Key0:"value0",
	}

	_, err := col.Upsert(testKey,testDoc, nil)
	if err != nil {
		fmt.Println(err)
	}
}


// Deleta a document
func deleteDocument(col *gocb.Collection,testKey string){
	_, err := col.Remove(testKey,nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted key " + testKey)
}

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