package main

import (
	"MD5-Hash-Computer/src/io/inputParser"
	"MD5-Hash-Computer/src/masterNode"
	"MD5-Hash-Computer/src/searchController"
)

//Main control flow function
//calls the run function to start the program
func main() {
	run()
}

func run() {
	data, metadata := inputParser.InputParserMain()
	wrapperObject, genObject := searchController.SearchControllerMain(data, metadata)
	masterNode.ToMasterNode(wrapperObject, genObject)
}