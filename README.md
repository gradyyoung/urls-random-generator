# URLS-RANDOM-GENERATOR

> 根据urls.txt生成搜索引擎需要的数据，并通过`curl`向搜索引擎提交

```shell
./urls-random-generator

-c int
    # 生成链接数量 (default 10)
-f string
    # 生成文件名称 (default "urls-baidu.txt")
-t string
    # 生成哪个搜索引擎的数据-baidu,bing (default "baidu")
-u string
    # 链接文件地址 (default "https://www.ygang.top/urls.txt")
```