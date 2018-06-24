# Go Survey Tool

My attempt using Go

## Running

```
go run src/survey-tool/question.go src/survey-tool/response.go \
       src/survey-tool/main.go \
       --survey src/survey-tool/example-data/survey-1.csv \
       --responses src/survey-tool/example-data/survey-1-responses.csv
```

## Build

```
go build src/survey-tool/*.go
```

## Todo

- [] Use `goroutine` to do the CSV parsing
- [] Add tests
- [] Setup travis
