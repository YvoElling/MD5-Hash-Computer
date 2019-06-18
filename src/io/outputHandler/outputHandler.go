package outputHandler

import (
	"MD5-Hash-Computer/datastructures/inputParser"
	"MD5-Hash-Computer/datastructures/resultHandler"
)

//typealiases for readability
type Metadata = *inputParser.Metadata
type ResultHandlerOutput = *resultHandler.ResultHandlerOutput

//function retrieve Metadata from the resulthandler
func getMetaDataObject () (metadata Metadata){

	return
}

//function writes Metadata to the output file
func writeMetaData( metadata2 Metadata) (success bool){

	return success
}

//functions gets MD5wrapper object from the resultHandler
//returns true when written
func getResultHandlerData ( output ResultHandlerOutput) bool{

	return true
}

//functions write default text and layout on .txt output file
func prepareOutputFile() {

}

//function that writes the MD5 wrapper data to the output .txt file
func writeToFile( output ResultHandlerOutput) bool{

	return true
}
