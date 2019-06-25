// Packages that contains all datastructures related to the inputParser software component
package inputParser

type MD5hash string

//HashMD5 struct definition. It is used to store data from the input .txt file
type HashMD5 struct {
	md5 MD5hash
}

func (hashMD5 *HashMD5) SetMd5(md5 MD5hash) {
	hashMD5.md5 = md5
}

//Returns the MD5 hash value from the struct
func (hashMD5 *HashMD5) getHash() MD5hash {
	return (*hashMD5).md5
}
