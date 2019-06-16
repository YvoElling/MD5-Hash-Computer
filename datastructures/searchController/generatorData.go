//this packages handles all datastructures for the search controller software component
package searchController

//typedefs for readability
type ascii int
type workers int
type maxWordLength int

//GeneratorData struct definition
type GeneratorData struct {
	minAscii ascii
	maxAscii ascii
	nrofworkers workers
	maxWordLength maxWordLength
}

//getter for amount of workers
func (generatorData *GeneratorData) getNrofworkers() workers {
	return (*generatorData).nrofworkers
}

//getter for minimal ascii number
func (generatorData *GeneratorData) getMinAscii() ascii {
	return (*generatorData).minAscii
}


//getter for maximum ascii number
func (generatorData *GeneratorData) getMaxAscii() ascii {
	return (*generatorData).maxAscii
}


//getter for maximum length of word
func (generatorData *GeneratorData) getMaxWordLength() maxWordLength {
	return (*generatorData).maxWordLength
}