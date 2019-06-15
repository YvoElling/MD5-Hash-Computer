#General Purpose of this project
This project is created by Yvo Elling on May 18th 2019. It is intended to be used to get more familiar with the 
Go programming language and moreover working with version control software altogether. Therefore, this can be considered
as a test of both application of the Golang programming language and integration with GitHub. 

##What does this program do?
This project is an implementation of a multi-threaded worker-master MD5 hasher. It needs to compute all possible words of
a certain range and output all the answers that have the same MD5 hash value. To make the program faster and to work with
concurrency, this is done using seperate worker thread, that all take one string, compute its value and return whether or 
not it is indeed the same as the provided value

##Is this program interactive?
Yes, the program should accept input from stdout or file input to make the program more dynamically. Also, the amount of
characters and the length of the word should be computed. 

##What are extra features?
It would be great if the program could provide some estimation upon how long it would take to find the right answer. It
would also be a nice inclusion if the program would show its progress, either by a loading bar or by simply displaying 
what string are currently being handled. 