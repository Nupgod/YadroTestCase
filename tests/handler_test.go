package tests

import (
	"bufio"
	"os"
	"fmt"
	"testing"
	"YadroTestCase/pkg/logger"


)
const layout string = "15:04"
func TestClub_HandleEvents_Result(t *testing.T) {
	// Путь к файлу с ожидаемыми результатами
	validResultFile := "testdata/valid_result.txt"

	// Чтение ожидаемых результатов из файла
	validResult, err := readLinesFromFile(validResultFile)
	if err != nil {
		t.Errorf("Failed to read valid result file: %v", err)
	}

	// Создание экземпляра Club и обработка событий
	club, err:= logger.ParseFile("testdata/valid.txt") 
	if err != nil {
		t.Errorf("Parser error: %s", err)
	}

	club.HandleEvents()

	res := club.StartTime.Format(layout)
	if  res != validResult[0] {
		t.Errorf("Different start time: %s\nExpected:  %s\n", res, validResult[0])
	}

	for i, event := range club.Result {
		if event.TableNum == 0 {
			res = fmt.Sprintf("%s %v %s", event.Time.Format(layout), event.ID, event.Client)
		} else {
			res = fmt.Sprintf("%s %v %s %v", event.Time.Format(layout), event.ID, event.Client, event.TableNum)
		}
		if res != validResult[i+1] {
			t.Errorf("Different event: %s\nExpected:  %s\n", res, validResult[i+1])
		}
	}

	res = club.EndTime.Format(layout)
	if res != validResult[len(validResult)-4] {
		t.Errorf("Different end time: %s\nExpected: %s\n", res, validResult[len(validResult)-4])
	}

	for i, table := range club.Tables {
		res = fmt.Sprintf("%d %d %s", i+1, table.Income, logger.FmtDuration(club.Tables[i].Duration))

		if res != validResult[len(validResult)-3+i] {
			t.Errorf("Different end time: %s\nExpected: %s\n", res, validResult[len(validResult)-3+i])
		}
	}
}

// Функция для чтения строк из файла
func readLinesFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}