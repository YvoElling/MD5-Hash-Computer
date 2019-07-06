package main

import (
	"MD5-Hash-Computer/src/io/inputParser"
	searchController2 "MD5-Hash-Computer/src/searchController"
)

//Main control flow function
//calls the run function to start the program
func main() {
	run()
}

func run() {
	data, metadata := inputParser.InputParserMain()
	searchController2.SearchControllerMain(data, metadata)
}