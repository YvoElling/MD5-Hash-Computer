package workerNodes

import (
	"MD5-Hash-Computer/datastructures/masterNode"
	"hash"
)

//typealiases for readability
type Job = *masterNode.WorkerJob

//Declares the state of the worker
const (
	active = iota
	readyForJob
	terminating
)

//function that request a new Job from the masterNode
func getNewJob( job Job) {

}

//function that is used to iterate over all MD5 values provided
func iterateOverArray(job Job) {

}

//function that is used to generate the MD5 value of the provided string
func generateMD5( string string) (md5hash hash.Hash){

	return
}

//function that compares the computed MD5 hash value with the hash value in the wrapper
//returns true if the values are the same and returns false if they are not.
func compare( wrapperValue hash.Hash, computedValue hash.Hash) bool{

	return true
}

//function that sets the worker node into 'terminating mode'
//kill the worker node after the completion of the current job.
func isLast( job Job) bool {

	return false
}