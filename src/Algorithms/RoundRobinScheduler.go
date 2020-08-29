package algorithms

import "fmt"
/**
Question 1:
Given a scheduler arrangement with some jobs having unique job ids and every job id has a number of tasks with unique task ids for that job.
Example:-

 Job id : 500    Task id : 700
      Task id : 300
      Task id : 350
 Job id : 600    Task id : 400
      Task id : 350
      Task id : 600
 Job id : 1000    Task id : 800
      Task id : 100
 Job id : 200    Task id : 650
Write a Code for listing the job ids + task id combination in round robin fashion.
Example:-

Job id : 500    Task id : 700
Job id : 600    Task id : 400
Job id : 1000    Task id : 800
Job id : 200    Task id : 650
Job id : 500    Task id : 300
Job id : 600    Task id : 350
Job id : 1000    Task id : 100
Job id : 500    Task id : 350
Job id : 600    Task id : 600
Choose any data structure of the input (array of structures or linked lists or queue).
Design Test cases for the same.
 */
func main(){
	// Input Data is that we will get a list of job ID and corresponding list of task id associated
	// with that job id. We need to print the data in round robin fashion
	inputDataSet := make(map[int][]int)
	inputDataSet[500] = []int{700, 300, 350}
	inputDataSet[600] = []int{400, 350, 600}
	inputDataSet[1000] = []int{800, 100}
	inputDataSet[200] = []int{650}

	printDataInRoundRobin(&inputDataSet)
}

func printDataInRoundRobin(dataSetMap *map[int][]int) {
	globalIndex := 0

	jobExecuted := true
	for ;jobExecuted == true;  {
		jobExecuted = false
		for key, jobList := range *dataSetMap {
			if globalIndex < len(jobList) {
				fmt.Printf("Job id : %d   Task id : %d \n ", key, jobList[globalIndex])
				jobExecuted = true
			}
		}
		globalIndex++
	}
}
