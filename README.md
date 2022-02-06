# Go Generics

This is just a repository to play around with Generics and Fuzzing, two big things coming with go 1.18!


## To Use

`go mod tidy`

`go build`

## Testing

To run regular unit tests:

`go test -v ./...`

To run Fuzzing:

`go test -v ./genericmath -fuzz=Fuzz`

The above will execute test cases forever, so to end it early you can either `ctrl-c` or run with the additional flag `-fuzztime 30s`


