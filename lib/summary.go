package lib

import (
	"bufio"
	"fmt"
	"github.com/improbable-io/go-junit-report/parser"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Results map[string][]*Test

const (
	PASS = "pass"
	FAIL = "fail"
	SKIP = "skip"
)

var jsonTestKeys = map[parser.Result]string{
	parser.PASS: PASS,
	parser.FAIL: FAIL,
	parser.SKIP: SKIP,
}

type Test struct {
	PackageName string `json:"package_name"`
	TestName    string `json:"test_name"`
	Time        int    `json:"time"`
	Output      string `json:"output"`
}

type Coverage struct {
	PackageName string  `json:"package_name"`
	Coverage    float64 `json:"coverage"`
}

type TestSummary struct {
	TotalTests    int         `json:"total_tests"`
	BuildErrors   string      `json:"build_errors"`
	Results       Results     `json:"results"`
	TotalCoverage float64     `json:"total_coverage"`
	Coverages     []*Coverage `json:"coverages"`
}

func Parse(stdoutReader, stderrReader, coverageReader io.Reader) (*TestSummary, error) {
	results := Results{
		PASS: []*Test{},
		FAIL: []*Test{},
		SKIP: []*Test{},
	}

	res, err := parser.Parse(stdoutReader, "")
	if err != nil {
		return nil, err
	}

	totalTests := 0
	for _, pkg := range res.Packages {
		for _, t := range pkg.Tests {
			key, _ := jsonTestKeys[t.Result]

			jsonTest := &Test{
				PackageName: pkg.Name,
				TestName:    t.Name,
				Time:        t.Time,
				Output:      strings.Join(t.Output, "\n"),
			}

			results[key] = append(results[key], jsonTest)
			totalTests += 1
		}
	}

	var coverages []*Coverage
	numCoverages, totalCoverage := 0., 0.
	fmt.Println(totalCoverage)
	scanner := bufio.NewScanner(coverageReader)
	for scanner.Scan() {
		line := scanner.Text()

		cov := 0.
		matches := regexp.MustCompile("\\d+\\.?\\d*%").FindAllString(line, -1)
		if len(matches) == 1 {
			no := strings.TrimRight(matches[0], "%")
			if cov, err = strconv.ParseFloat(no, 64); err != nil {
				panic(err)
			}
		}

		name := strings.Split(line, "\t")[1]
		c := &Coverage{
			PackageName: name,
			Coverage:    cov,
		}
		coverages = append(coverages, c)
		numCoverages += 1
	}
	for _, c := range coverages {
		totalCoverage += c.Coverage
	}
	if numCoverages != 0 {
		totalCoverage /= numCoverages
	}

	buildErrorBytes, err := ioutil.ReadAll(stderrReader)
	if err != nil {
		return nil, err
	}

	fmt.Println(numCoverages, totalCoverage)

	summary := &TestSummary{
		TotalTests:    totalTests,
		Results:       results,
		BuildErrors:   string(buildErrorBytes),
		TotalCoverage: totalCoverage,
		Coverages:     coverages,
	}

	return summary, nil
}
