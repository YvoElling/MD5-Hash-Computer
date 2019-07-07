package masterNode

import (
	"MD5-Hash-Computer/datastructures/masterNode"
	"MD5-Hash-Computer/datastructures/searchController"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Type aliases for code readability
type wrapper = searchController.WrapperMD5
type generatorData = searchController.GeneratorData
type job = masterNode.WorkerJob

//receives input from the SearchController
//Forwards received data to MasterNodeMain (the master node control flow)
func ToMasterNode( wrapperArray []wrapper, generatorData generatorData) {
	if !masterNodeMain(wrapperArray, generatorData) {
		log.Fatal("Error in passing data to MasterNodeMain")
	}
}

//Receives input data from 'ToMasterNode' function
//start the control flow from the master node
func masterNodeMain( wrapperArray []wrapper, generatorData generatorData) bool{
	var counter []int
	stringval, _ := generateStringStart(counter, generatorData)
	fmt.Println(stringval)
	return true
}

//returns the amount of workers of type integer
func getNrofWorkers ( generatorData generatorData) int {
	return int(generatorData.GetNrofworkers())
}

//Returns a string that was generated
func generateString(generatorData generatorData, stringIndexCounter []int) (result string, counter []int) {
	result, counter = generateStringStart(stringIndexCounter, generatorData)
	return
}

//Starts the string generation procedure
//function should be called first in the string generation method
func generateStringStart( counter []int, genData generatorData) (generatedString string, updatedCounter []int){
	var stringValue []string
	if int(genData.GetMinAscii()) + counter[0] < int(genData.GetMaxAscii()) {
		asciiValue := int(genData.GetMinAscii()) + counter[0]
		stringValue = append(stringValue, strconv.Itoa(asciiValue))
		counter[0]++
		generateStringRecurse(&stringValue, &counter, 0, genData)
	}
	return strings.Join(stringValue, ""), counter
}

//This function is called maxWordLengthCount times
func generateStringRecurse(tempString* []string, counter* []int, count int, genData generatorData){
	if count < int(genData.GetMaxWordLength()) {
		if (*counter)[count] < int(genData.GetMaxAscii()){
			asciiValue := int(genData.GetMinAscii()) + (*counter)[count]
			*tempString = append(*tempString, strconv.Itoa(asciiValue))
			(*counter)[count]++
			count++
			generateStringRecurse(tempString, counter, count, genData)
		} else {
			asciiValue := int(genData.GetMinAscii()) + (*counter)[count]
			*tempString = append(*tempString, strconv.Itoa(asciiValue))
			count++
			generateStringRecurse(tempString, counter, count, genData)
		}
	}
	return
}

//Initialize workers
func createWorkers( nrofworker int) bool{

	return true
}

//Send a job to a node
//receive an integer with the amount of hits the worker made
func sendJob( job job) int{

	return 0
}