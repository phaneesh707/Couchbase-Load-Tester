package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
)


func monitor(file *os.File){
	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		message := fmt.Sprintf("[main] memory useage : Allocated = %v MB \tHeap Allocated : %v MB \n",memStats.Alloc / 1024 /1024,memStats.HeapAlloc / 1024 /1024)

		memAllocated := memStats.Alloc / 1024 /1024
		if memAllocated > 2048 {
			fmt.Println("Memory usage is more than 2 GB , exiting progrma !")
			os.Exit(1)
		}

		log.Printf(message)
		_ ,err := file.WriteString(message)
		if err != nil {
			fmt.Println("Error writing to log file : ",err)
		}
		time.Sleep(1 * time.Second)
	}
}


func main() {
	// Uncomment following line to enable logging
	// gocb.SetLogger(gocb.VerboseStdioLogger())

	connectionString := "localhost"
	bucketName := "test"
	username := "admin"
	password := "password"
	
	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	
	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	col := bucket.Scope("_default").Collection("_default")

	fmt.Println("Connection succesful!")
    
	// ----------------------------------------------------

	// log file for memory profiling
	file , err := os.Create("memory-usage.txt")
	if err != nil {
		fmt.Println("Error creting file")
	}
	defer file.Close()

	// For monitoring purpose
	waitC := make(chan bool)

	go monitor(file)


	totalWorker := 96
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id int
		value int
	}

	totalTask := 1000000
	resultC := make(chan result, totalTask)

	docKeyDelete := ""
	docKeyUpdate := ""

	for i := 0; i < totalTask; i++ {
		id := uuid.New()
		testKey := id.String()
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", i)
			insertDocument(col,testKey)
			resultC <- result{i, i * 2}
		})

		if i % 10 == 0  {
			wp.AddTask(func() {
				updateDocument(col,docKeyUpdate)
			})
			docKeyUpdate = testKey
		}

		if i % 100 == 0 {
			wp.AddTask(func() {
				deleteDocument(col,docKeyDelete)
			})
			docKeyDelete = testKey
		}

	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d", res.id, res.value)
	}

	
	<-waitC
}








