# Go Concurrent Couchbase Loader

Loader which is used to insert, update, and delete docs into couchbase server and also does  memory and CPU profiling.


## Code Structure

The project consists of three files:

- `main.go`:  sets up the Couchbase server connection, creates the worker pool, and manages  tasks that are added to the worker pool.
- `operations.go`: contains functions to perform CRUD operations on the Couchbase database and also operations related to worker pool.
- `types.go`: defines custom types for workerpool

## Memory & CPU Usage Monitoring

Memory usage is monitored using the `runtime/debug` package. A goroutine is created in the main function that continuously reads memory statistics using `runtime.ReadMemStats` and logs the memory usage in a text file named `memory-usage.txt`. The program will exit automatically if the memory usage exceeds 2 GB.

CPU usage is monitored using the `runtime` package. The main function logs the number of goroutines and CPU usage at regular intervals.


## Worker Pool

The `WorkerPool` type has two methods: `Run` and `AddTask`. The `Run` method creates a fixed number of goroutines, and the `AddTask` method adds tasks to a channel that is monitored by the goroutines.

The main function creates a worker pool with a fixed number of workers and manages tasks by calling the `AddTask` method. The main function reads data from a JSON file named `doc.json` and performs the CRUD operations on the data using the worker pool.



