// Packages that contains all datastructures related to the inputParser software component
package inputParser

import "hash"

//HashMD5 struct definition. It is used to store data from the input .txt file
type HashMD5 struct {
	md5 hash.Hash
}

func (hashMD5 *HashMD5) SetMd5(md5 hash.Hash) {
	hashMD5.md5 = md5
}

//Returns the MD5 hash value from the struct
func (hashMD5 *HashMD5) getHash() hash.Hash{
	return (*hashMD5).md5
}
