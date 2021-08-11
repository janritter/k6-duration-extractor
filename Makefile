prepare:
	go mod download

build: prepare
	go build -o ./bin/k6-duration-extractor -v

tests:
	go test ./... -v  --cover

run:
	go run main.go result.csv
