# Computer Club Logger
## About
This is a prototype system that tracks the operation of a computer club, processes events, and calculates revenue and time spent at each table. The solution is implemented in Golang.

## Limitations and Assumptions
- The program is implemented using only the standard Golang library, without using third-party libraries.
- Input data must match the described format; otherwise, the program will output an error.
- Events are processed sequentially in time, checking their correctness and impact on the current state of the computer club.
  
## Testing
The program includes test examples to verify the correctness of event processing and revenue calculation.

## Installation Guide
### 1. Downloading the Project
You can download the project by cloning the repository from GitHub:
``` git clone git@github.com:Nupgod/YadroTestCase.git ```
### 2. Move to repository
``` cd YadroTestCase ```
### 3. Build App
#### Build by Makefile
```make build```
#### Build by GO
```go build -o logger .\cmd\main.go```
#### Build by Docker
```	
docker build -t yadro-logger .
docker run yadro-logger
```
### 4. Run app
```logger <filepath>```
