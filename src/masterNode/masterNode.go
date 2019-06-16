package masterNode

import (
	"MD5-Hash-Computer/datastructures/masterNode"
	"MD5-Hash-Computer/datastructures/searchController"
)

type wrapper *searchController.WrapperMD5
type genData *searchController.GeneratorData
type job *masterNode.WorkerJob

//receives input from the SearchController
//Forwards received data to MasterNodeMain (the master node control flow)
func ToMasterNode( wrapperArray []wrapper, generatorData genData) bool{

	return true
}

//Receives input data from 'ToMasterNode' function
//start the control flow from the master node
func masterNodeMain( wrapperArray []wrapper, generatorData genData) bool{

	return true
}

//returns the amount of workers of type integer
func getNrofWorkers ( generatorData genData) int {

	return 0
}

//Returns a string that was generated
func generateString(generatorData genData) (result string) {

	return
}


//Starts the string generation procedure
//function should be called first in the string generation method
func generateStringStart() {

}

//This function is called maxWordLengthCount times
func generateStringRecurse() {

}

//This function is used to terminate the string generation process once maxWordLengthCount has been reached
func generateStringBase() {

}

//Initialize workers
func createWorkers( nrofworker int) bool{

}

//Send a job to a node
//receive an integer with the amount of hits the worker made
func sendJob( job job) int{

}