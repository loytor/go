package supports

import (
	"fmt"
	"regexp"
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
	fmt.Print("\r", countFinish, countWait, ":", lineI, ":")
}

func Mimei(s string) bool {
	b, _ := regexp.MatchString(`^[\w\d]{15}$`, s)
	return s != "NULL" && b
}

func Midfa(s string) bool {
	b, _ := regexp.MatchString(`^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`, s)
	return s != "NULL" && b
}

func Moaid(s string) bool {
	return s != "NULL" && s != ""
}
