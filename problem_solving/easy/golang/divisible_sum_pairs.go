package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Int32Set struct {
	items map[int32]bool
}

func (s *Int32Set) addToSet(v int32) {
	found := s.isInSet(v)
	if found == false {
		s.items[v] = true
	}
}

func (s *Int32Set) isInSet(v int32) bool {
	_, found := s.items[v]
	return found
}

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	l := len(ar)
	var count int32 = 0
	if l < 2 {
		return 0
	}

	set := Int32Set{}
	set.items = make(map[int32]bool)
	for _, val := range ar {
		set.addToSet(val)
	}

	//fmt.Println(set.items)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count += 1
			}
		}
	}

	//fmt.Println(count)
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := divisibleSumPairs(n, k, ar)

	fmt.Fprintf(writer, "%d\n", result)

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
