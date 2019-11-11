# 使用Golang发送http请求

最近做了一个回调url的功能，即数据库里面存在一个url字段(可能是http的url或者是grpc的url)，一个urlproto字段(这个字段的作用是判断url的类型，0为http，1为grpc这种)，要实现的功能为如果这个url是http类型，我就通过httpclient发送请求，如果url是grpc，那就通过grpcclient发送grpc请求，非常简单的一个功能，由于grpc请求涉及到api结构，所以这里只记录一下golang发送http请求。

## 正文

```go
func Post() {

	url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1/data"

	payload := strings.NewReader("a=111")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Date", "Tue, 11 Sep 2018 10:57:09 GMT")
	req.Header.Add("Authorization", "oqSBNbmgAAGI155F")
	req.Header.Add("Content-Type", "text/plain")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}


//Get携带各种参数、数组
func Get() {

	url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "BNbmgAAGI155")

	res, _ := http.DefaultClient.Do(req)
	urlWithParameter := req.URL.Query()
	urlWithParameter.Add("doamin", domain)
	for i := 0; i < len(onlyIncludeIds); i++ {
		urlWithParameter.Add("array", array[i])
	}
		req.URL.RawQuery = urlWithParameter.Encode()
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}



func Put() {

	url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1"

	payload := strings.NewReader("a=111")
    
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bmgAAGI155F6MJ3N2T")
	req.Header.Add("Date", "Wed, 12 Sep 2018 02:10:09 GMT")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}



func Delete() {

	url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1"

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("Authorization", "5F6MJ3N2Tk9ruL_6XQpx-uxkkg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
```

## 总结

上面这四个请求，每一个都有defer res.Body.Close()，看了一下函数解释，说是发送请求时收到的body必须手动关闭，否则会造成GC无法清理的情况。