package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

/* Bruteforce */
func getTotalX_bruteforce(a []int32, b []int32) int32 {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	var count int32 = 0

	for i := a[len(a)-1]; i <= b[0]; i++ {
		if is_all_factors(i, a) && is_factor_of(i, b) {
			count += 1
		}

	}
	fmt.Println(count)
	return count
}

/* Use GCD and LCM */
func getTotalX(a []int32, b []int32) int32 {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	var count int32 = 0
	var lcm_a int32 = 0
	var gcd_b int32 = 0

	len_a := len(a)
	len_b := len(b)

	if len_a < 2 {
		lcm_a = a[0]
	} else {
		lcm_a = lcm(a[0], a[1])
		if len_a > 2 {
			for i := 2; i < len_a; i++ {
				lcm_a = lcm(lcm_a, a[i])
			}
		}
	}

	if len_b < 2 {
		gcd_b = b[0]
	} else {
		gcd_b = gcd(b[0], b[1])
		if len_b > 2 {
			for i := 2; i < len_b; i++ {
				gcd_b = gcd(gcd_b, b[i])
			}
		}
	}

	incr := lcm_a

	for lcm_a <= gcd_b {
		if gcd_b%lcm_a == 0 {
			count += 1
			lcm_a += incr
		} else {
			lcm_a += incr
		}
	}

	return count
}

func gcd(a int32, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a int32, b int32) int32 {
	return a * b / gcd(a, b)
}

func is_all_factors(n int32, a []int32) bool {
	for i := 0; i < len(a); i++ {
		if n%a[i] != 0 {
			return false
		}
	}
	return true
}

func is_factor_of(n int32, b []int32) bool {
	for i := 0; i < len(b); i++ {
		if b[i]%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
