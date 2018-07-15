package main

import (
	"test/log/resource"
	"fmt"
	"flag"
	"os"
	"net/url"
	"time"
	"math/rand"
	"strings"
)

func main() {
	total := flag.Int("total", 100, "需要生成多少行日志 ?")
	filePath := flag.String("filePath", "./dog.log", "日志路径")
	flag.Parse() //解析

	//生成 url规则
	result := resource.ReuleResource()
	urlList := resource.BuildUlr(result)
	for i, url := range urlList {
		fmt.Printf("第 %d 次;url = %s\n", i, url)
	}

	//按照要求,生成 total 行日志文件
	f, err := os.Open(*filePath)

	if err != nil {
		f, err = os.Create(*filePath)
		if err != nil {
			panic("创建文件失败")
		}
		panic("打开文件失败")
	}

	defer f.Close()

	var logStr string
	for i := 0; i <= *total; i++ {
		currentUlr := urlList[randInt(0, len(urlList)-1)]
		referUrl := urlList[randInt(0, len(urlList)-1)]
		ua := uaList()[randInt(0, len(uaList())-1)]
		logStr = makeLog(currentUlr, referUrl, ua)
		logStr = logStr + logStr
	}

	f.Write([]byte("1111"))
	fmt.Println("done")
}

/**
 生成随机数.
 */
func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	if min > max {
		return max
	}

	return r.Intn(max-min) + min
}

/**
  拼装日志文件格式.
 */
func makeLog(currentUlr, referUrl, ua string) string {
	u := url.Values{}
	u.Set("time", string(time.Now().Unix()))
	u.Set("url", currentUlr)
	u.Set("refer", referUrl)
	u.Set("ua", ua)
	paramsStr := u.Encode()
	logTemplate := "192.168.0.101 - - [14/Jul/2018:20:14:13 -0400] \"GET /dog?{$paramsStr} HTTP/1.1\" 200 43 \"-\" \"{$ua}\" \"-\""
	log := strings.Replace(logTemplate, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	return log
}

func uaList() (list []string) {
	list = []string{
		"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
		"MQQBrowser/25 (Linux; U; 2.3.3; zh-cn; HTC Desire S Build/GRI40;480*800)",
		"Mozilla/5.0 (Linux; U; Android 2.3.3; zh-cn; HTC_DesireS_S510e Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaE75-1 /110.48.125 Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Mobile/8J2",
		"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.51.22 (KHTML, like Gecko) Version/5.1.1 Safari/534.51.22",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; XBLWP7; ZuneWP7)",
		"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
		"Mozilla/5.0 (Windows NT 5.1; rv:5.0) Gecko/20100101 Firefox/5.0",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
		"Mozilla/4.0 (compatible; MSIE 60; Windows NT 5.1; SV1; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
		"Opera/9.80 (Windows NT 5.1; U; zh-cn) Presto/2.9.168 Version/11.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.04506.648; .NET CLR 3.5.21022; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/533.21.1 (KHTML, like Gecko) Version/5.0.5 Safari/533.21.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 2.0.50727; TheWorld)",
	}

	return list
}
