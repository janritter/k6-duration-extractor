prepare:
	go mod download

build: prepare
	go build -o ./bin/k6-duration-extractor -v -ldflags "-X github.com/janritter/k6-duration-extractor/cmd.gitSha=`git rev-parse HEAD` -X github.com/janritter/k6-duration-extractor/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/janritter/k6-duration-extractor/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v  --cover

run:
	go run main.go
