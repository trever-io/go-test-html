# go-test-html

Converts `go test` output into a prettified HTML summary.

![Example HTML output](https://cloud.githubusercontent.com/assets/2838283/18885683/8f53b7fa-84e4-11e6-9ca8-f5e3101acf98.png)

## Installation

Go version 1.1 or higher is required. Install or update using the `go get`
command:

```bash
go get -u github.com/trever-io/go-test-html
```

## Usage

The go-test-html command takes in files containing the stdout and stderr from the `go test` command and the location
of the HTML file to write.

```bash
go-test-html [gotest_stdout_file] [gotest_stderr_file] [gotest_coverage] [name] [output_file]
```

To produce the `gotest_stdout_file` and `gotest_stderr_file` without changing the output of your `go test` runs, use
the following command.

```bash
go test ./... -v > ./test.out 2> ./test.err
```

To produce the `gotest_coverage` use the following command.

```bash
go test  -cover ./... > ./coverage.out
```
    
Here is full usage example.

```bash
go test ./... -v > ./test.out 2> ./test.err
go test  -cover ./... > ./coverage.out
go-test-html ./test.out ./test.err ./coverage.out "My Fancy Go Project" test-protocol.html
```
