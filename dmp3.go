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
		// userinfo_id,device_id,identity,os,device_type,dt
		sli := strings.Split(s, ",")
		if sli[4] == "imei" {
			_, _ = imei.WriteString(sli[1] + "\n")
		}
		if sli[4] == "idfa" {
			_, _ = idfa.WriteString(sli[1] + "\n")
		}
		if sli[4] == "oaid" {
			_, _ = oaid.WriteString(sli[1] + "\n")
		}
	})
}
