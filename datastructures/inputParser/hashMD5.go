// Packages that contains all datastructures related to the inputParser software component
package inputParser

type Hash string

//HashMD5Struct struct definition. It is used to store data from the input .txt file
type HashMD5Struct struct {
	md5 Hash
}

func (hashMD5 *HashMD5Struct) SetMd5(md5 Hash) {
	hashMD5.md5 = md5
}

//Returns the MD5 hash value from the struct
func (hashMD5 *HashMD5Struct) GetHash() Hash {
	return (*hashMD5).md5
}
