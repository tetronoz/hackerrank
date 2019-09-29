package main

func staircase (n int32) {
	c := "#"
	for i := n-1; i > 0; i-- {
		fmt.Printf("%*s\n", n ,c)
		c += "#"

	}
	fmt.Println(c)
}

func main() {
	reader := buffio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader)), 10, 64)
	CheckErr(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *buffio.Reader) string {
	str, _, err: = reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}