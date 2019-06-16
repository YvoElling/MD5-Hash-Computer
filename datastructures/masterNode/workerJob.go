//package contains all datastructures for the masterNode
package masterNode

import "MD5-Hash-Computer/datastructures/searchController"

//typdefs for readability
type job string
type wrapperMD5 *searchController.WrapperMD5

//workerJob struct
//This struct includes all data that is sent from the masterNode to the workerNode
type WorkerJob struct {
	job job
	wrapperArray []wrapperMD5
	lastJob bool
}

//retuns the string that is to be hashed by the worker
func (workerJob *WorkerJob) getJob() job {
	return (*workerJob).job
}

//returns an array of wrapperMD5 objects
func (workerJob *WorkerJob) getWrapperArray() []wrapperMD5{
	return (*workerJob).wrapperArray
}

//returns whether or not this was the final job
func (WorkerJob *WorkerJob) isLastJob() bool  {
	return (*WorkerJob).lastJob
}