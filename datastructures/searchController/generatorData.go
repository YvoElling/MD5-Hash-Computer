//this packages handles all datastructures for the search controller software component
package searchController

//typedefs for readability
type Ascii int
type Workers int
type MaxWordLength int

//GeneratorData struct definition
type GeneratorData struct {
	minAscii Ascii
	maxAscii Ascii
	nrofworkers Workers
	maxWordLength MaxWordLength
}

//getter for amount of workers
func (generatorData *GeneratorData) GetNrofworkers() Workers {
	return (*generatorData).nrofworkers
}

//getter for minimal ascii number
func (generatorData *GeneratorData) GetMinAscii() Ascii {
	return (*generatorData).minAscii
}


//getter for maximum ascii number
func (generatorData *GeneratorData) GetMaxAscii() Ascii {
	return (*generatorData).maxAscii
}


//getter for maximum length of word
func (generatorData *GeneratorData) GetMaxWordLength() MaxWordLength {
	return (*generatorData).maxWordLength
}