package lib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	stdOutFile   string
	stdErrFile   string
	coverageFile string
	summary      *TestSummary
}

var testCases = []TestCase{
	{
		stdOutFile:   "01-stdout",
		stdErrFile:   "01-stderr",
		coverageFile: "01-coverage",
		summary: &TestSummary{
			TotalTests:  3,
			BuildErrors: "",
			Results: Results{
				PASS: []*Test{
					{
						PackageName: "cobra/doc",
						TestName:    "TestBashCompletions",
						Time:        10,
						Output:      "",
					},
					{
						PackageName: "cobra/doc",
						TestName:    "TestBashCompletionHiddenFlag",
						Time:        0,
						Output:      "",
					},
					{
						PackageName: "cobra/doc",
						TestName:    "TestBashCompletionDeprecatedFlag",
						Time:        0,
						Output:      "",
					},
				},
				FAIL: []*Test{},
				SKIP: []*Test{},
			},
			TotalCoverage: 2.58,
			Coverages:     []*Coverage{
				{
					PackageName: "treverLiqgo",
					Coverage:    0,
				},
				{
					PackageName: "treverLiqgo/common",
					Coverage:    8.3,
				},
				{
					PackageName: "treverLiqgo/common/influx",
					Coverage:    0,
				},
				{
					PackageName: "treverLiqgo/toms",
					Coverage:    4.1,
				},
				{
					PackageName: "treverLiqgo/toms/db",
					Coverage:    0.5,
				},
			},
		},
	},
	{
		stdOutFile:   "02-stdout",
		stdErrFile:   "02-stderr",
		coverageFile: "02-coverage",
		summary: &TestSummary{
			TotalTests:  3,
			BuildErrors: "",
			Results: Results{
				PASS: []*Test{
					{
						PackageName: "cobra/doc",
						TestName:    "TestGenManNoGenTag",
						Time:        0,
						Output:      "",
					},
				},
				FAIL: []*Test{
					{
						PackageName: "cobra/doc",
						TestName:    "TestGenManDoc",
						Time:        0,
						Output:      "cmd_test.go:144: Line: 59 Unexpected response.",
					},
				},
				SKIP: []*Test{
					{
						PackageName: "cobra/doc",
						TestName:    "TestGenMdNoTag",
						Time:        0,
						Output:      "md_docs_test.go:72: yolo",
					},
				},
			},
			TotalCoverage: 0,
			Coverages:     nil,
		},
	},
	{
		stdOutFile:   "03-stdout",
		stdErrFile:   "03-stderr",
		coverageFile: "03-coverage",
		summary: &TestSummary{
			TotalTests:  1,
			BuildErrors: "build-error",
			Results: Results{
				PASS: []*Test{},
				FAIL: []*Test{
					{
						PackageName: "cobra/doc",
						TestName:    "TestGenManDoc",
						Time:        0,
						Output:      "cmd_test.go:144: Line: 59 Unexpected response.",
					},
				},
				SKIP: []*Test{},
			},
			TotalCoverage: 0.,
			Coverages:     nil,
		},
	},
}

func TestSummaryParser(t *testing.T) {
	for _, testCase := range testCases {
		t.Logf("Running: (%s, %s)", testCase.stdOutFile, testCase.stdErrFile)

		stdoutFile, err := os.Open("tests/" + testCase.stdOutFile)
		require.NoError(t, err)

		stdoutErr, err := os.Open("tests/" + testCase.stdErrFile)
		require.NoError(t, err)

		coverage, err := os.Open("tests/" + testCase.coverageFile)
		require.NoError(t, err)

		actual, err := Parse(stdoutFile, stdoutErr, coverage)
		require.NoError(t, err)
		require.NotNil(t, actual)
		assert.EqualValues(t, testCase.summary, actual)
	}
}
