package supports

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

// 生成指定数量的联系的字符串
func NChars(b string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += b
	}
	return s
}

func PBar(lineI int, lineLen int) {
	countFinish := NChars("#", int(lineI*100/lineLen))
	countWait := NChars("-", int(100-lineI*100/lineLen))
	print("\r", countFinish, countWait, ":", lineI, ":")
}

func Mimei(s string) bool {
	b, _ := regexp.MatchString(`^\w{15}$`, s)
	return s != "NULL" && b
}

func Midfa(s string) bool {
	b, _ := regexp.MatchString(`^[A-Z\d]{8}-[A-Z\d]{4}-[A-Z\d]{4}-[A-Z\d]{4}-[A-Z\d]{12}$`, s)
	return s != "NULL" && b
}

func Moaid(s string) bool {
	b, _ := regexp.MatchString(`^[a-z\d]{8}-[a-z\d]{4}-[a-z\d]{4}-[a-z\d]{4}-[a-z\d]{12}$`, s)
	return s != "NULL" && b
}

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close() // 延时处理
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 读取buf
	buf := bufio.NewReader(f)

	// 循环处理
	i := 1
	for {
		if i%1000 == 0 {
			fmt.Println("No", "-", i, "-", time.Now().Format("15:04:05"), ";", "	")
		}
		line, _, err := buf.ReadLine()
		if (err == io.EOF) {
			print("is end")
			return nil
		}

		line = []byte(strings.TrimSpace(string(line)))
		handle(string(line))
		i++
	}
}
