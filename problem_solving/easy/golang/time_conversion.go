package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	if strings.HasSuffix(s, "AM") {
		if s == "12:00:00AM" {
			return "00:00:00"
		}
		hour, _ := strconv.ParseInt(s[0:2], 0, 64)
		if hour == 12 {
			return "00" + ":" + s[3:8]
		} else {
			return s[0:2] + ":" + s[3:8]
		}
	}

	if strings.HasSuffix(s, "PM") {
		if s == "12:00:00PM" {
			return "12:00:00"
		}

		hour, _ := strconv.ParseInt(s[0:2], 0, 64)
		if hour == 12 {
			return s[0:2] + ":" + s[3:8]
		} else {
			hour += 12
			return strconv.Itoa(int(hour)) + ":" + s[3:8]
		}
	}
	return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
