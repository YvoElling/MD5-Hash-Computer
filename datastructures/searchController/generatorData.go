//this packages handles all datastructures for the search controller software component
package searchController

//typedefs for readability
type Ascii int
type Workers int
type MaxWordLength int

//GeneratorData struct definition
type GeneratorData struct {
	MinAscii Ascii
	MaxAscii Ascii
	Nrofworkers Workers
	MaxWordLength MaxWordLength
}

//getter for amount of workers
func (generatorData *GeneratorData) GetNrofworkers() Workers {
	return (*generatorData).Nrofworkers
}

//getter for minimal ascii number
func (generatorData *GeneratorData) GetMinAscii() Ascii {
	return (*generatorData).MinAscii
}


//getter for maximum ascii number
func (generatorData *GeneratorData) GetMaxAscii() Ascii {
	return (*generatorData).MaxAscii
}


//getter for maximum length of word
func (generatorData *GeneratorData) GetMaxWordLength() MaxWordLength {
	return (*generatorData).MaxWordLength
}