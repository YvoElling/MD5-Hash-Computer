//this package contains all the datastructures for the resultHandler software comoponent
package resultHandler

import (
	"MD5-Hash-Computer/datastructures/searchController"
)

//typedef for readability
type output searchController.WrapperMD5

//resultHandler object class
//This is an MD5 wrapper as declared before
type ResultHandlerOutput struct {
	MD5output output
}

//returns the data from the MD5 wrapper
func ( resultHandlerOutput *ResultHandlerOutput) getOutputData() output {
	return (*resultHandlerOutput).MD5output
}