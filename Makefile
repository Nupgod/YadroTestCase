build:
	go build .\cmd\main.go

run:
	.\main.exe .\test_file.txt

rebuild: build run

docker:
	docker build -t yadro-logger .
	docker run yadro-logger