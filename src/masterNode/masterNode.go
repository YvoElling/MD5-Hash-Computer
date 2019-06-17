package masterNode

import (
	"MD5-Hash-Computer/datastructures/masterNode"
	"MD5-Hash-Computer/datastructures/searchController"
)

//Type aliases for code readability
type wrapper = *searchController.WrapperMD5
type generatorData = *searchController.GeneratorData
type job = *masterNode.WorkerJob

//receives input from the SearchController
//Forwards received data to MasterNodeMain (the master node control flow)
func ToMasterNode( wrapperArray []wrapper, generatorData generatorData) bool{

	return true
}

//Receives input data from 'ToMasterNode' function
//start the control flow from the master node
func masterNodeMain( wrapperArray []wrapper, generatorData generatorData) bool{

	return true
}

//returns the amount of workers of type integer
func getNrofWorkers ( generatorData generatorData) int {

	return 0
}

//Returns a string that was generated
func generateString(generatorData generatorData) (result string) {

	return
}

//Starts the string generation procedure
//function should be called first in the string generation method
func generateStringStart() (generatedString string){

	return
}

//This function is called maxWordLengthCount times
func generateStringRecurse(tempString string){

	return
}

//This function is used to terminate the string generation process once maxWordLengthCount has been reached
func generateStringBase( returnString string){

	return
}

//Initialize workers
func createWorkers( nrofworker int) bool{

	return true
}

//Send a job to a node
//receive an integer with the amount of hits the worker made
func sendJob( job job) int{

	return 0
}