package searchController

import (
	"MD5-Hash-Computer/datastructures/inputParser"
	"MD5-Hash-Computer/datastructures/searchController"
)

type hashMD5 = inputParser.HashMD5Struct
type metaData = inputParser.Metadata
type wrapper = searchController.WrapperMD5

//main function that is in control of the control flow of the searchcontroller
func SearchControllerMain( input []hashMD5, metadata metaData) (wrappedOut []wrapper, data searchController.GeneratorData){
	//create a channel such that other threads can return values
	ch := make(chan []wrapper)
	//start the wrapMD5 function in a seperate thread
	go wrapMD5(input, ch)
	//start the createGeneratorData in this thread
	data = createGeneratorData(metadata)
	//listen to the channel to receive the result of the WrapMD5 function
	wrappedOut = <- ch
	return
}

//create the generatordata object based on the metadata received from the input
func createGeneratorData( data metaData) (generatorObject searchController.GeneratorData){
	//copy the relevant dat to the new struct
	generatorObject.SetMaxWordLength(data.GetMaxWordLength())
	generatorObject.SetMaxAscii(data.GetMaxAscii())
	generatorObject.SetMinAscii(data.GetMinAscii())
	generatorObject.SetNrofworkers(data.GetNrofworkers())
	return
}

//wraps the MD5 values in the new struct with the corresponding field data
func wrapMD5 ( data []hashMD5, ch chan []wrapper) {
	//create a an array of wrappers and a temporary wrap object
	var wrap []wrapper
	var tempWrap wrapper

	for _, object := range data{
		//add each value found by the inputParser to the new array
		tempWrap.SetHashMD5(object)
		wrap = append(wrap, tempWrap)
	}
	//send the new wrap back to the main thread
	ch <- wrap
	close(ch)
}
