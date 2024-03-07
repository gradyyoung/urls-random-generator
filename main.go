package main

import (
	"encoding/json"
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
	fileName := flag.String("f", "urls-baidu.txt", "生成文件名称")
	searchType := flag.String("t", "baidu", "生成哪个搜索引擎的数据-baidu,bing")
	url := flag.String("u", "https://www.ygang.top/urls.txt", "链接文件地址")
	bingKey := flag.String("bingKey", "989***************b6", "Bing提交链接时的key")
	flag.Parse()
	fmt.Printf("Count:%d\nFileName:%s\nSearchType:%s\nUrl:%s\n", *count, *fileName, *searchType, *url)
	urls := getRandomUrls(*count, *url)
	var str string
	switch *searchType {
	case "baidu":
		str = strings.Join(urls, "\n")
	case "bing":
		m := map[string]any{
			"host":        "www.ygang.top",
			"key":         *bingKey,
			"keyLocation": fmt.Sprintf("https://www.ygang.top/%s.txt", *bingKey),
			"urlList":     urls,
		}
		j, _ := json.Marshal(m)
		str = string(j)
	}
	writeToFile(*fileName, str)
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
	if count >= len(urls) {
		return urls
	}
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
