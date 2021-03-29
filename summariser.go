// a small command-line application
// - makes multiple calls to an HTTP API
// - generates a simple file for us.
// - have an empty test environment setup that is ready to be run.

// run the command line with -type flag and -date flag
// parse flags
// fail if invalid
// create get request to get type of news on specified date
// unmarshall results
// write results to a file

package main

import (
	"flag"
	"fmt"
	"github.com/winnab/summariser/runner"
)

func main() {
	wordPtr := flag.String("section", "", "what news section would you like to see")
	flag.Parse()
	fmt.Printf("Yes, you can request %v news.", *wordPtr)

	runner.Run()
}
