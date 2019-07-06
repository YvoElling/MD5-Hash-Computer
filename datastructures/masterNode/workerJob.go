//package contains all datastructures for the masterNode
package masterNode

import "MD5-Hash-Computer/datastructures/searchController"

//typdefs for readability
type wrapperMD5 *searchController.WrapperMD5

//workerJob struct
//This struct includes all data that is sent from the masterNode to the workerNode
type WorkerJob struct {
	job string
	wrapperArray []wrapperMD5
	lastJob bool
}

//retuns the string that is to be hashed by the worker
func (this *WorkerJob) getJob() string {
	return (*this).job
}

//returns an array of wrapperMD5 objects
func (this *WorkerJob) getWrapperArray() []wrapperMD5{
	return (*this).wrapperArray
}

//returns whether or not this was the final job
func (this *WorkerJob) isLastJob() bool  {
	return (*this).lastJob
}