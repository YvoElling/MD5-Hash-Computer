package searchController

import (
	"MD5-Hash-Computer/datastructures/inputParser"
)

//wrapperMD5 structs
//it is used to embed the hash value with its corresponding strings
type WrapperMD5 struct {
	hashMD5 inputParser.HashMD5Struct
	correspondingStrings []string
}

func (wrapperMD5 *WrapperMD5) SetHashMD5(hashMD5 inputParser.HashMD5Struct) {
	wrapperMD5.hashMD5 = hashMD5
}

func (WrapperMD5 *WrapperMD5) AddString(newString string) {
	WrapperMD5.correspondingStrings = append(WrapperMD5.correspondingStrings, newString)
}

//returns the hash value
func (wrapperMD5 *WrapperMD5) getHashMD5Object() inputParser.HashMD5Struct {
	return (*wrapperMD5).hashMD5
}

//returns the list of strings tha hash to the same hash value
func (wrapperMD5 *WrapperMD5) getStringArray() []string {
	return (*wrapperMD5).correspondingStrings
}