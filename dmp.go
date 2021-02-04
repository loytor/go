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
	// 生成文件前缀
	fPrefix := "dmp/csv-" + time.Now().Format("150405")

	fIdfa, _ := os.OpenFile("./"+fPrefix+"-idfa.txt", os.O_RDWR|os.O_CREATE, 0766);
	fImei, _ := os.OpenFile("./"+fPrefix+"-imei.txt", os.O_RDWR|os.O_CREATE, 0766);
	//fAdid, _ := os.OpenFile("./"+fPrefix+"-androidid.txt", os.O_RDWR|os.O_CREATE, 0766);
	fOaid, _ := os.OpenFile("./"+fPrefix+"-oaid.txt", os.O_RDWR|os.O_CREATE, 0766);
	//fMd5, _ := os.OpenFile("./"+fPrefix+"-md5.txt", os.O_RDWR|os.O_CREATE, 0766);

	// 读取本地文件
	originFile := os.Args[1]
	buf, _ := ioutil.ReadFile(originFile)
	originates := strings.Split(string(buf), "\n")
	lineLen := len(originates)
	lineI := 0
	for _, txt := range originates {
		// 显示进度条
		lineI++
		supports.PBar(lineI, lineLen)

		sli := strings.Split(txt, "\t")

		// 写入imei
		mImei, _ := regexp.MatchString(`^[0-9]{15}$`, sli[3])
		if sli[3] != "NULL" && mImei {
			_, _ = fImei.WriteString(sli[3] + "\n")
		}

		mIdfa, _ := regexp.MatchString(`^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`, sli[6])
		if sli[6] != "NULL" && mIdfa {
			_, _ = fIdfa.WriteString(sli[6] + "\n")
		}

		if sli[4] != "NULL" {
			_, _ = fOaid.WriteString(sli[4] + "\n")
		}
	}

}
