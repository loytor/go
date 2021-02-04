package main

import (
	"demo01/supports"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	// 生成随机文件前缀
	fPrefix := "dmp/" + time.Now().Format("150405") + "-chiyt"

	// 打开文件流
	fIdfa, _ := os.OpenFile("./"+fPrefix+"-idfa.txt", os.O_RDWR|os.O_CREATE, 0766);
	fImei, _ := os.OpenFile("./"+fPrefix+"-imei.txt", os.O_RDWR|os.O_CREATE, 0766);
	fAdid, _ := os.OpenFile("./"+fPrefix+"-androidid.txt", os.O_RDWR|os.O_CREATE, 0766);
	fOaid, _ := os.OpenFile("./"+fPrefix+"-oaid.txt", os.O_RDWR|os.O_CREATE, 0766);
	fMd5, _ := os.OpenFile("./"+fPrefix+"-md5.txt", os.O_RDWR|os.O_CREATE, 0766);

	// 获取原始文件 /Users/yanbingti/Desktop/olduser-0115.txt
	originFile := os.Args[1]
	buf, _ := ioutil.ReadFile(originFile)
	originates := strings.Split(string(buf), "\n")
	lineLen := len(originates)
	lineI := 0
	for _, txt := range originates {
		// 显示进度条
		lineI++
		supports.PBar(lineI, lineLen)

		// 匹配idfa
		mIdfa, _ := regexp.MatchString(`^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`, txt)
		if mIdfa {
			_, _ = fIdfa.WriteString(txt + "\n")
		}

		// 匹配imei
		mImei, _ := regexp.MatchString(`^[0-9]{15}$`, txt)
		if mImei {
			_, _ = fImei.WriteString(txt + "\n")
		}

		// 匹配AndroidId - 至少包含
		mAdid, _ := regexp.MatchString(`^[\w\d]{16}$`, txt)
		if mAdid {
			_, _ = fAdid.WriteString(txt + "\n")
		}

		// 匹配md5
		mMd5, _ := regexp.MatchString(`^[\w\d]{32}$`, txt)
		if mMd5 {
			_, _ = fMd5.WriteString(txt + "\n")
		}

		// 啥也不是放入oaid
		if !mIdfa && !mImei && !mAdid && !mMd5 {
			_, _ = fOaid.WriteString(txt + "\n")
		}
	}
}
