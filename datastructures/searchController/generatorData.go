//this packages handles all datastructures for the search controller software component
package searchController

import "MD5-Hash-Computer/datastructures/inputParser"

//GeneratorData struct definition
type GeneratorData struct {
	minAscii inputParser.Ascii
	maxAscii inputParser.Ascii
	nrofworkers inputParser.Workers
	maxWordLength inputParser.MaxWordLength
}

//setter
func (generatorData *GeneratorData) SetMaxWordLength(maxWordLength inputParser.MaxWordLength) {
	generatorData.maxWordLength = maxWordLength
}

//setter
func (generatorData *GeneratorData) SetNrofworkers(nrofworkers inputParser.Workers) {
	generatorData.nrofworkers = nrofworkers
}

//setter
func (generatorData *GeneratorData) SetMaxAscii(maxAscii inputParser.Ascii) {
	generatorData.maxAscii = maxAscii
}

//setter
func (generatorData *GeneratorData) SetMinAscii(minAscii inputParser.Ascii) {
	generatorData.minAscii = minAscii
}

//getter for amount of workers
func (generatorData *GeneratorData) GetNrofworkers() inputParser.Workers {
	return (*generatorData).nrofworkers
}

//getter for minimal ascii number
func (generatorData *GeneratorData) GetMinAscii() inputParser.Ascii {
	return (*generatorData).minAscii
}


//getter for maximum ascii number
func (generatorData *GeneratorData) GetMaxAscii() inputParser.Ascii {
	return (*generatorData).maxAscii
}


//getter for maximum length of word
func (generatorData *GeneratorData) GetMaxWordLength() inputParser.MaxWordLength {
	return (*generatorData).maxWordLength
}