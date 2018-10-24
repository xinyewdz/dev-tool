package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tool"
)

func main() {
	//testMd5()
	//testRandom()
	//testDate()
	//testString()

	//testFmt()
	for {
		var typeName string
		fmt.Printf("\n\n\ninput type:\n")
		fmt.Printf("1 date2unix \n")
		fmt.Printf("2 unix2date \n")
		fmt.Printf("3 base64encode \n")
		fmt.Printf("4 base64decode \n")
		fmt.Printf("5 md5 \n")
		fmt.Printf("6 SHA1 \n")
		fmt.Printf("7 randow string \n")
		fmt.Printf("8 rename file\n")
		fmt.Scanln(&typeName)
		run(typeName)
	}

}

func run(typeName string) {
	switch typeName {
	case "1":
		{
			fmt.Printf("date(yyyy-MM-dd HH:mm:ss):")
			data := readData()
			result := tool.Date2Unix(data)
			fmt.Printf("%s to timestamp:%d\n", data, result)
		}
	case "2":
		{
			fmt.Printf("timestamp(second):")
			data := readData()
			num, _ := strconv.Atoi(data)
			result := tool.Unix2Date(int64(num))
			fmt.Printf("%s to date:%s\n", data, result)
		}
	case "3":
		{
			fmt.Printf("origin string:")
			data := readData()
			result := tool.Base64Encode(data)
			fmt.Printf("%s base64Encode:%s\n", data, result)
		}
	case "4":
		{
			fmt.Printf("encode string:")
			data := readData()
			result := tool.Base64Decode(data)
			fmt.Printf("%s Base64Decode:%s\n", data, result)
		}
	case "5":
		{
			fmt.Printf("MD5 string:")
			data := readData()
			result := tool.Md5(data, true)
			fmt.Printf("%s MD5:%s\n", data, result)
		}
	case "6":
		{
			fmt.Printf("SHA1 string:")
			data := readData()
			result := tool.Sha1(data, true)
			fmt.Printf("%s SHA1:%s\n", data, result)
		}
	case "7":
		{
			fmt.Printf("Random string,input length:")
			data := readData()
			len, _ := strconv.Atoi(data)
			result := tool.RandomStr(len)
			fmt.Printf("%s Random:%s\n", data, result)
		}
	case "8":
		{
			fmt.Printf("input path:\n")
			path := readData()
			fmt.Printf("input newNamePrefix:\n")
			prefix := readData()
			tool.Rename(path, prefix)
			fmt.Printf("finish.\n")
		}
	case "exit":
		{
			os.Exit(0)
		}
	}
}

func readData() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = data[:len(data)-2]
	return data
}

func testFmt() {
	var name string
	fmt.Scanf("%s%s", &name)
	fmt.Printf("input:%s\n", name)
}

func testString() {
	str := "nihao"
	str = str[:2]
	fmt.Printf("%s\n", str)
}

func testDate() {
	str := "2015-12-23 12:50:52"
	ts := tool.Date2Unix(str)
	fmt.Printf("%s=%d\n", str, ts)
}

func testBase() {
	data := "hello"
	str := tool.Base64Encode(data)
	newData := tool.Base64Decode(str)
	fmt.Printf("data:%s,encode:%s\n", data, str)
	fmt.Printf("encode:%s,data:%s\n", str, newData)
}

func testMd5() {
	data := "nihao"
	result := tool.Md5(data, false)
	fmt.Printf("data:%s,cry:%s\n", data, result)
	sh1Result := tool.Sha1(data, true)
	fmt.Printf("data:%s,cry:%s\n", data, sh1Result)
}

func testRandom() {
	str := tool.RandomStr(10)
	fmt.Printf("str:%s\n", str)
}
