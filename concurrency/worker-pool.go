package main

import (
	"fmt"
	"sync"
	"time"
)

// struct Task is defined with field ID of type int
type Task struct {
	ID int
}

/*
Pointer receiver method on type Task to simulate a running process by sleeping it for 10 seconds.
*/
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	time.Sleep(time.Second * 10)
	fmt.Printf("Processed task %d\n", t.ID)
}

/*
Struct WorkerPool having the following fields:
1. Slice of tasks
2. Concurrency: count of concurrent workers to spawn
3. Buffered channel tasksChan or type Task
4. Waitgroup to wait for all workers to process the tasks
*/
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

/*
Worker method which listens to new tasks pushed on the buffered channel and processes it once arrived. After processing, it calls Done() method of waitgroup to decrement the waitgroup counter
*/
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

/*
Run method on WorkerPool which does the following things:
1. Initialize a buffered channel in WorkerPool of size equal to the length of tasks
2. Spawn worker goroutines equal to the value of concurrency field
3. Add the length of tasks as the counter in waitgroup
4. Push the task on the buffered channel: tasksChan for the workers to pick up the tasks to process
5. After all the tasks sent on the channel, close the channel indicating no more tasks will be sent after this
6. Wait until the counter in waitgroup goes to zero to prevent the function from returning (it is a blocking operation)
*/
func (wp *WorkerPool) Run() {

	wp.tasksChan = make(chan Task, len(wp.Tasks)) // 1

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker() // 2
	}
	wp.wg.Add(len(wp.Tasks)) // 3
	for _, task := range wp.Tasks {
		wp.tasksChan <- task // 4
	}
	close(wp.tasksChan) // 5

	wp.wg.Wait() // 6
}

func workerPoolMain() {
	tasks := make([]Task, 20)
	for i := range 20 {
		tasks[i] = Task{ID: i + 1}
	}
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}
	wp.Run()
	fmt.Println("All tasks are processed!")
}
