package inputParser

import (
	"MD5-Hash-Computer/datastructures/inputParser"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var path, _ = filepath.Abs("input.txt")

//Start the input parser process
func InputParserMain() {
	md5array := getFileData()
	for _, value := range md5array {
		fmt.Println(string(value.GetHash()))
	}
}

//retrieve file data from input file
func getFileData() (md5Array []inputParser.HashMD5Struct){
	//opens a connection to the input file
	var file, e =  os.Open(path)
	if e != nil {
		log.Fatal(e)
	}

	//stores the data from the file in the data string slice
	data := storeInObject(file)

	//closes the file after operation
	err := file.Close()
	if err != nil {
		log.Fatal(e)
	}

	//evaluate the content
	md5Array = evalFileData(data)

	return md5Array
}

//evaluates file data
func evalFileData( data []string) (md5Array []inputParser.HashMD5Struct){
	//create metadata object to forward to outputHandler
	var MetaData inputParser.Metadata

	//iterate over the created string slice
	for i, stringData := range data {
		//if the index value is less than 6, this means that this data is metadat
		if i < 7 {
			switch i {
			case 0:
				//index nr 1: The date
				day := inputParser.Day(stringData[0] + stringData[1])
				month := inputParser.Month(stringData[3] + stringData[4])
				year := inputParser.Year(stringData[6] + stringData[7] + stringData[8] + stringData[9])
				date := inputParser.Date{Day: day, Month: month, Year: year}
				MetaData.SetDate(date)
			case 1:
				//index nr 2: Name of the owner
				name := inputParser.Name(stringData)
				MetaData.SetName(name)
			case 2:
				//index nr 3: The amount of workers used in the computations
				temp, err := strconv.Atoi(stringData)
				if err != nil {
					log.Fatal(err)
				}
				nrofworkers := inputParser.Workers(temp)
				MetaData.SetNrofworkers(nrofworkers)
			case 3:
				//index nr 4: The filename for the generated output file
				fileName := inputParser.Filename(stringData)
				MetaData.SetFileName(fileName)
			case 4:
				//index nr 5: The minimal ascii value used to generate strings in the masterNode
				temp, err := strconv.Atoi(stringData)
				if err != nil {
					log.Fatal(err)
				}
				minAscii := inputParser.Ascii(temp)
				MetaData.SetMinAscii(minAscii)
			case 5:
				//index nr 6: the maximal ascii value used to generate strings in the MasterNode
				temp, err := strconv.Atoi(stringData)
				if err != nil {
					log.Fatal(err)
				}
				maxAscii := inputParser.Ascii(temp)
				MetaData.SetMaxAscii(maxAscii)
			case 6:
				//index nr 7: The maximum length of the generated string by the masterNode
				temp, err := strconv.Atoi(stringData)
				if err != nil {
					log.Fatal(err)
				}
				maxWordLength := inputParser.MaxWordLength(temp)
				MetaData.SetMaxWordLength(maxWordLength)
			}

			//If the data is a comment, then we can disregard it
		} else if !isComment(stringData) {
			//create an inputParser.HashM5 object
			var hashMd5Object inputParser.HashMD5Struct
			//store string data in MD5 object.
			hashMd5Object.SetMd5(inputParser.Hash(stringData))
			//Add new MD5 object to MD5array
			md5Array = append(md5Array, hashMd5Object)
		}
	}
	//return the MD5 array
	return md5Array
}

//checks if data is comment
func isComment( input string) bool{
	//if the strings starts with an "#", this is a comment
	if strings.HasPrefix(input, "#") {
		return true
	}
	return false
}

//stores the data in object
func storeInObject(file *os.File) (output []string){
	//creates a input scanner to read file data
	scanner := bufio.NewScanner(file)

	//if there is something to scan (i.e. there is still unread text)
	//read it and store it in the input string slice
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		output = append(output, scanner.Text())
	}

	//after the reading has completed, check if the scanner has reported any errors
	if e := scanner.Err(); e != nil {
		//if so, log and return
		log.Fatal(e)
	}
	return
}

