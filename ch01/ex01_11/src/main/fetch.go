package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<- ch)
	}
	fmt.Printf("fetched %d urls, %.2fs elasped\n", len(os.Args[1:]), time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	url = getURL(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func getURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}

/*
<実行結果>
-----------------------------------
0.44s      81 http://Baidu.com
0.49s   19347 http://Google.com
0.50s   19678 http://Yahoo.co.jp
0.51s   21246 http://Google.co.in
0.55s   79486 http://Bing.com
0.56s       0 http://Instagram.com
0.63s  156362 http://Google.de
0.64s   19247 http://Google.co.jp
0.86s   17297 http://Yandex.ru
1.08s   44209 http://Msn.com
1.34s  173870 http://Ebay.com
1.55s  248655 http://Twitter.com
1.55s   41554 http://Linkedin.com
1.60s    5645 http://Weibo.com
1.66s  526843 http://Sina.com.cn
1.72s  351875 http://Youtube.com
1.72s    6064 http://Vk.com
1.91s  417027 http://Amazon.com
2.16s   71175 http://Taobao.com
2.18s  393804 http://Yahoo.com
2.21s   70019 http://Facebook.com
2.29s   54537 http://Wikipedia.org
2.36s  708128 http://Hao123.com
3.48s    9656 http://Live.com
22.29s  617945 http://Qq.com
fetched 25 urls, 22.29s elasped
-----------------------------------

全てのURLからの返答がないと、チャネルの受信で処理が止まることがある

 */