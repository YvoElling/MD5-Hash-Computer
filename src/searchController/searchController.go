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
	ch := make(chan []wrapper)
	go wrapMD5(input, ch)
	data = createGeneratorData(metadata)
	wrappedOut = <- ch
	return
}

//create the generatordata object based on the metadata received from the input
func createGeneratorData( data metaData) (generatorObject searchController.GeneratorData){
	generatorObject.SetMaxWordLength(data.GetMaxWordLength())
	generatorObject.SetMaxAscii(data.GetMaxAscii())
	generatorObject.SetMinAscii(data.GetMinAscii())
	generatorObject.SetNrofworkers(data.GetNrofworkers())
	return
}

//wraps the MD5 values in the new struct with the corresponding field data
func wrapMD5 ( data []hashMD5, ch chan []wrapper) {
	var wrap []wrapper
	var tempWrap wrapper

	for _, object := range data{
		tempWrap.SetHashMD5(object)
		wrap = append(wrap, tempWrap)
	}
	ch <- wrap
	close(ch)
}
