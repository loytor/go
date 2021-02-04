package main

import (
	"bufio"
	"demo01/supports"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// 获取文件地址
	originFile := os.Args[1]

	// -1
	idfaD1, _ := os.OpenFile("./dmp/idfaD1.txt", os.O_RDWR|os.O_CREATE, 0766);
	imeiD1, _ := os.OpenFile("./dmp/imeiD1.txt", os.O_RDWR|os.O_CREATE, 0766);
	oaidD1, _ := os.OpenFile("./dmp/oaidD1.txt", os.O_RDWR|os.O_CREATE, 0766);
	// -2
	idfaD2, _ := os.OpenFile("./dmp/idfaD2.txt", os.O_RDWR|os.O_CREATE, 0766);
	imeiD2, _ := os.OpenFile("./dmp/imeiD2.txt", os.O_RDWR|os.O_CREATE, 0766);
	oaidD2, _ := os.OpenFile("./dmp/oaidD2.txt", os.O_RDWR|os.O_CREATE, 0766);
	// +1
	idfaI1, _ := os.OpenFile("./dmp/idfaI1.txt", os.O_RDWR|os.O_CREATE, 0766);
	imeiI1, _ := os.OpenFile("./dmp/imeiI1.txt", os.O_RDWR|os.O_CREATE, 0766);
	oaidI1, _ := os.OpenFile("./dmp/oaidI1.txt", os.O_RDWR|os.O_CREATE, 0766);
	// +2
	idfaI2, _ := os.OpenFile("./dmp/idfaI2.txt", os.O_RDWR|os.O_CREATE, 0766);
	imeiI2, _ := os.OpenFile("./dmp/imeiI2.txt", os.O_RDWR|os.O_CREATE, 0766);
	oaidI2, _ := os.OpenFile("./dmp/oaidI2.txt", os.O_RDWR|os.O_CREATE, 0766);

	// 循环处理文件
	_ = ReadFile(originFile, func(s string) {
		// user_type, userinfo_id, device_id, imei, oaid, idfa, os
		sli := strings.Split(s, ",")
		if sli[0] == "-1" {
			if supports.Mimei(sli[3]) {
				_, _ = imeiD1.WriteString(sli[3] + "\n")
			}

			if supports.Midfa(sli[5]) {
				_, _ = idfaD1.WriteString(sli[5] + "\n")
			}

			if supports.Moaid(sli[4]) {
				_, _ = oaidD1.WriteString(sli[4] + "\n")
			}

			// 测试
			if sli[3] == "NULL" && sli[4] == "NULL" && sli[5] == "NULL" {
				if supports.Mimei(sli[2]) {
					_, _ = imeiD1.WriteString(sli[2] + "\n")
				}

				if supports.Midfa(sli[2]) {
					_, _ = idfaD1.WriteString(sli[2] + "\n")
				}
			}
		}

		if sli[0] == "-2" {
			if supports.Mimei(sli[3]) {
				_, _ = imeiD2.WriteString(sli[3] + "\n")
			}

			if supports.Midfa(sli[5]) {
				_, _ = idfaD2.WriteString(sli[5] + "\n")
			}

			if supports.Moaid(sli[4]) {
				_, _ = oaidD2.WriteString(sli[4] + "\n")
			}

			// 测试
			if sli[3] == "NULL" && sli[4] == "NULL" && sli[5] == "NULL" {
				if supports.Mimei(sli[2]) {
					_, _ = imeiD2.WriteString(sli[2] + "\n")
				}

				if supports.Midfa(sli[2]) {
					_, _ = idfaD2.WriteString(sli[2] + "\n")
				}
			}
		}

		if sli[0] == "+1" {
			if supports.Mimei(sli[3]) {
				_, _ = imeiI1.WriteString(sli[3] + "\n")
			}

			if supports.Midfa(sli[5]) {
				_, _ = idfaI1.WriteString(sli[5] + "\n")
			}

			if supports.Moaid(sli[4]) {
				_, _ = oaidI1.WriteString(sli[4] + "\n")
			}
			// 测试
			if sli[3] == "NULL" && sli[4] == "NULL" && sli[5] == "NULL" {
				if supports.Mimei(sli[2]) {
					_, _ = imeiI1.WriteString(sli[2] + "\n")
				}

				if supports.Midfa(sli[2]) {
					_, _ = idfaI1.WriteString(sli[2] + "\n")
				}
			}
		}

		if sli[0] == "+2" {
			if supports.Mimei(sli[3]) {
				_, _ = imeiI2.WriteString(sli[3] + "\n")
			}

			if supports.Midfa(sli[5]) {
				_, _ = idfaI2.WriteString(sli[5] + "\n")
			}

			if supports.Moaid(sli[4]) {
				_, _ = oaidI2.WriteString(sli[4] + "\n")
			}
			// 测试
			if sli[3] == "NULL" && sli[4] == "NULL" && sli[5] == "NULL" {
				if supports.Mimei(sli[2]) {
					_, _ = imeiI2.WriteString(sli[2] + "\n")
				}

				if supports.Midfa(sli[2]) {
					_, _ = idfaI2.WriteString(sli[2] + "\n")
				}
			}
		}
	})
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
		if i%10000 == 0 {
			fmt.Print("No", "-", i, "-", time.Now().Format("15:04:05"), ";", "	")
		}
		line, _, err := buf.ReadLine()
		fmt.Println(line, err)
		if (err == io.EOF) {
			print("is end")
			return nil
		}

		line = []byte(strings.TrimSpace(string(line)))
		handle(string(line))
		i++
	}
}
