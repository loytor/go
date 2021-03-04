package main

import (
	"demo01/supports"
	"os"
	"strings"
)

func main() {
	// 获取文件地址
	originFile := os.Args[1]

	idfa, _ := os.OpenFile("./dmp/idfa.txt", os.O_RDWR|os.O_CREATE, 0766);
	imei, _ := os.OpenFile("./dmp/imei.txt", os.O_RDWR|os.O_CREATE, 0766);
	oaid, _ := os.OpenFile("./dmp/oaid.txt", os.O_RDWR|os.O_CREATE, 0766);

	// 循环处理文件
	_ = supports.ReadFile(originFile, func(s string) {
		// userinfo_id	os	imei	oaid	android_id	idfa	idfv	caid
		sli := strings.Split(s, "	")

		if len(sli) >= 3 && supports.Mimei(sli[2]) {
			_, _ = imei.WriteString(sli[2] + "\n")
		}
		if len(sli) >= 4 && supports.Moaid(sli[3]) {
			_, _ = oaid.WriteString(sli[3] + "\n")
		}
		if len(sli) >= 6 && supports.Midfa(sli[5]) {
			_, _ = idfa.WriteString(sli[5] + "\n")
		}
	})
}
