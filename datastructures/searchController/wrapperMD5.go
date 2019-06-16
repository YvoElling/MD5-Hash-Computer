package searchController

import (
	"MD5-Hash-Computer/datastructures/inputParser"
)

//wrapperMD5 structs
//it is used to embed the hash value with its corresponding strings
type WrapperMD5 struct {
	hashMD5 inputParser.HashMD5
	correspondingStrings []string
}

//returns the hash value
func (wrapperMD5 *WrapperMD5) getHashMD5Object() inputParser.HashMD5 {
	return (*wrapperMD5).hashMD5
}

//returns the list of strings tha hash to the same hash value
func (wrapperMD5 *WrapperMD5) getStringArray() []string {
	return (*wrapperMD5).correspondingStrings
}