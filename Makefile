build:
	go build -o logger .\cmd\main.go

run:
	.\logger .\test_file.txt

rebuild: build run

docker:
	docker build -t yadro-logger .
	docker run yadro-logger

test:
	go test ./tests