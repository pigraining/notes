package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpDoGet() {
	req, err := http.NewRequest("GET", "www.baidu.com", nil)
	if err != nil {
		//hander err
	}
	urlWithParameter := req.URL.Query()
	urlWithParameter.Add("arg1", "1")
	urlWithParameter.Add("arg2", "2")
	for i := 0; i < 3; i++ {
		urlWithParameter.Add("arg3", strconv.Itoa(i))

	}
	req.URL.RawQuery = urlWithParameter.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//hander err
	}
	defer resp.Body.Close()

	//处理返回的body值
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		//hander err
	}
	fmt.Println("body", body)
}
func main() {
	httpGet()
	httpPost()
	httpPostForm()
	httpDo()
	httpDoGet()
}
