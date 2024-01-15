package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 解析命令行
	count := flag.Int("c", 10, "生成链接数量")
	fileName := flag.String("f", "urls.txt", "生成文件名称")
	url := flag.String("u", "https://www.ygang.top/urls.txt", "链接文件地址")
	flag.Parse()
	fmt.Println(*count, *fileName, *url)
	str := getRandomUrls(*count, *url)
	writeToFile(*fileName, strings.Join(str, "\n"))
}

func writeToFile(fileName string, str string) {
	file, err := os.Create("./" + fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(str)
}

// 随机获取count个地址
func getRandomUrls(count int, url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := io.ReadAll(resp.Body)
	urls := strings.Split(string(body), "\n")
	result := make([]string, 0)
	indexs := randomNum(len(urls), count)
	for k, _ := range indexs {
		result = append(result, urls[k])
	}
	return result
}

// 生成count个不重复的随机数0-max
func randomNum(max, count int) map[int]struct{} {
	result := make(map[int]struct{}, count)
	for {
		num := rand.Intn(max) //[0~9]
		if _, exist := result[num]; exist {
			continue
		} else {
			result[num] = struct{}{}
			if count == len(result) {
				break
			}
		}

	}
	return result
}
