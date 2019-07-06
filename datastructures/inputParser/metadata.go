// Packages that contains all datastructures related to the inputParser software component
package inputParser


// typedefs for code readability
type Day int
type Month int
type Year int
type Name string
type Filename string
type Workers int
type Ascii int
type MaxWordLength int

// Date struct that is used for printing out the correct Date in the output file
type Date struct {
	Day Day
	Month Month
	Year Year
}

// metadata struct that is used for the output file and the masterNode
type Metadata struct {
	date          Date
	name          Name
	nrofworkers   Workers
	fileName      Filename
	minAscii      Ascii
	maxAscii      Ascii
	maxWordLength MaxWordLength
}

func (metadata *Metadata) SetMaxWordLength(maxWordLength MaxWordLength) {
	metadata.maxWordLength = maxWordLength
}

func (metadata *Metadata) SetMaxAscii(maxAscii Ascii) {
	metadata.maxAscii = maxAscii
}

func (metadata *Metadata) SetMinAscii(minAscii Ascii) {
	metadata.minAscii = minAscii
}

func (metadata *Metadata) SetFileName(fileName Filename) {
	metadata.fileName = fileName
}

func (metadata *Metadata) SetNrofworkers(nrofworkers Workers) {
	metadata.nrofworkers = nrofworkers
}

func (metadata *Metadata) SetName(name Name) {
	metadata.name = name
}

func (metadata *Metadata) SetDate(date Date) {
	metadata.date = date
}

//getter for Date
func (metadata *Metadata) GetDate() Date {
	return (*metadata).date
}

//getter for name
func (metadata *Metadata) GetName() Name {
	return (*metadata).name
}

//getter for amount of workers
func (metadata *Metadata) GetNrofworkers() Workers {
	return (*metadata).nrofworkers
}

//getter for output file name
func (metadata *Metadata) GetFilename() Filename {
	return (*metadata).fileName
}

//getter for minimal ascii number
func (metadata *Metadata) GetMinAscii() Ascii {
	return (*metadata).minAscii
}


//getter for maximum ascii number
func (metadata *Metadata) GetMaxAscii() Ascii {
	return  (*metadata).maxAscii
}


//getter for maximum length of word
func (metadata *Metadata) GetMaxWordLength() MaxWordLength {
	return  (*metadata).maxWordLength
}