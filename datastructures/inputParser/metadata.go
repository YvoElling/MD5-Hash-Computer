// Packages that contains all datastructures related to the inputParser software component
package inputParser


// typedefs for code readability
type day int
type month int
type year int
type name string
type filename string
type workers int
type ascii int
type maxWordLength int

// date struct that is used for printing out the correct date in the output file
type date struct {
	Day day
	Month month
	Year year
}

// metadata struct that is used for the output file and the masterNode
type Metadata struct {
	date date
	name name
	nrofworkers workers
	fileName filename
	minAscii ascii
	maxAscii ascii
	maxWordLength maxWordLength
}

//getter for date
func (metadata *Metadata) getDate() date {
	return (*metadata).date
}

//getter for name
func (metadata *Metadata) getName() name {
	return (*metadata).name
}

//getter for amount of workers
func (metadata *Metadata) getNrofworkers() workers {
	return (*metadata).nrofworkers
}

//getter for output file name
func (metadata *Metadata) getFilename() filename {
	return (*metadata).fileName
}

//getter for minimal ascii number
func (metadata *Metadata) getMinAscii() ascii {
	return (*metadata).minAscii
}


//getter for maximum ascii number
func (metadata *Metadata) getMaxAscii() ascii {
	return  (*metadata).maxAscii
}


//getter for maximum length of word
func (metadata *Metadata) getMaxWordLength() maxWordLength {
	return  (*metadata).maxWordLength
}