QQWry
=====

纯真IP库 golang 版

>有关纯真IP库的文档详见 http://lumaqq.linuxsir.org/article/qqwry_format_detail.html

## 1. 依赖
* mahonia 处理 GBK 编码的地址信息 (请确保先装好 `hg`)
```bash
go get code.google.com/p/mahonia
```

update:

```bash
go get github.com/yinheli/mahonia
```

* 纯真 IP 库

>下载地址 http://update.cz88.net/soft/qqwry.rar

## 2. 使用
* 下载
```bash
go get github.com/yinheli/qqwry
```
* 在项目中引入
```go
import (
	"github.com/yinheli/qqwry"
	"log"
)

func main() {
	q := qqwry.NewQQwry("qqwry.dat")
	q.Find("180.89.94.90")
	log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
	// output: 
	// 2014/02/22 22:10:32 ip:180.89.94.90, Country:北京市, City:长城宽带
}
```

## 3. 使用注意
* qqwry 不是线程安全的
* qqwry 没有使用缓存

## 4. license
Licensed under MIT.