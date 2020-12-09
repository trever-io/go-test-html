package main

import (
	"./lib"
	"fmt"
	"log"
	"os"

	"io/ioutil"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Incorrect command line arguments")
		fmt.Println("Usage: go-test-html [gotest_stdout_file] [gotest_stderr_file] [gotest_coverage] [output_file]")
		os.Exit(1)
	}

	//"Trever Liqgo version: %v-%v", GitTag, GitCommit

	gotestStdoutFile := os.Args[1]
	gotestStderrFile := os.Args[2]
	gotestCoverageFile := os.Args[3]
	outputFile := os.Args[4]

	gotestStdout, err := os.Open(gotestStdoutFile)
	check(err)

	gotestStderr, err := os.Open(gotestStderrFile)
	check(err)

	gotestCoverage, err := os.Open(gotestCoverageFile)
	check(err)

	summary, err := lib.Parse(gotestStdout, gotestStderr, gotestCoverage)
	check(err)

	templateBox := rice.MustFindBox("template")
	html, err := lib.GenerateHTML(templateBox.MustString("template.html"), summary)
	check(err)

	err = ioutil.WriteFile(outputFile, []byte(html), 0644)
	check(err)

	outputFilePath, err := filepath.Abs(outputFile)
	check(err)

	fmt.Printf("Test results written to '%s'\n", outputFilePath)
}
