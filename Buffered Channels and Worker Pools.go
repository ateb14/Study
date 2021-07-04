package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct{
	id int
}
type Result struct{
	task Task
}//completed tasks

var tasks = make(chan Task,10)
var results = make(chan Result,10)

func DoingTheTask() {
	time.Sleep(time.Second)//simulation of the time that is needed to finish the task
}
func distribute(NumberOfTask int){
	for i:=0;i<NumberOfTask;i++{
		id := i
		task := Task{id}
		tasks <- task//distributing the tasks to the channel
		fmt.Println("Task ",id," is distributed")
	}
	close(tasks)
}

func PrintResult(flag chan bool){
	for result := range results{
		fmt.Println("Task ",result.task.id," is finished")
	}
	flag<-true
}

func worker(wg *sync.WaitGroup){
	for task := range tasks{
		DoingTheTask()
		ret := Result{task}
		results <- ret
	}
	wg.Done()
}//It watches the tasks channel.
func CreatWorkerPools(NumberOfWorkers int){
	var wg sync.WaitGroup
	for i:=0;i<NumberOfWorkers;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func main(){
	StartTime:=time.Now()
	NumberOfTasks:=100
	NumberOfWorkers:=20
	flag:=make(chan bool)
	go distribute(NumberOfTasks)
	go CreatWorkerPools(NumberOfWorkers)
	go PrintResult(flag)
	<-flag
	EndTime:=time.Now()
	CostTime:=EndTime.Sub(StartTime)
	fmt.Println("Total time cost",CostTime)
}
