package outputHandler

import (
	"MD5-Hash-Computer/datastructures/inputParser"
	"MD5-Hash-Computer/datastructures/resultHandler"
)

//typealiases for readability
type metadata = *inputParser.Metadata
type resultHandlerOutput = *resultHandler.ResultHandlerOutput

//function retrieve metadata from the resulthandler
func getMetaDataObject () (metadata metadata){

	return
}

//function writes metadata to the output file
func writeMetaData( metadata2 metadata) (success bool){

	return success
}

//functions gets MD5wrapper object from the resultHandler
//returns true when written
func getResultHandlerData ( output resultHandlerOutput ) bool{

	return true
}

//functions write default text and layout on .txt output file
func prepareOutputFile() {

}

//function that writes the MD5 wrapper data to the output .txt file
func writeToFile( output resultHandlerOutput) bool{

	return true
}
