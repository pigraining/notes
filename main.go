package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var mapStr  = make(map[float64]string)
var slice []float64
var pathStr string = "/home/s/logs/lanxin/lx_app/"

func loadLog(path string) {
	file ,err := os.Open(path)
	if err != nil {
		fmt.Println("打开log文件失败",err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//是否有下一行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "msg:rpc-call"){
			extractInformation(scanner.Text())
		}
	}
}
func extractInformation(text string){
	strArr := strings.Split(text, "real:")
	strArr2 := strings.Split(strArr[1], "reqid:")
	str := strings.TrimSpace(strArr2[0])

	num,_ := strconv.ParseFloat(str,3)
	mapStr[num] = text
	slice = append(slice,num)
}

func operation(){
	sort.Float64s(slice)
}

func printMaxs(num int) {
	if num > len(slice){
		num = len(slice)
	}
	for a := 0; a <num; a++ {
		fmt.Println(mapStr[slice[len(slice)-a-1]])
	}
}

func main(){
	fmt.Println("请输入绝对路径(默认路径为/home/s/logs/lanxin/lx_app/)")
	var str string
	fmt.Scanln(&str)
	if !strings.Contains(str,"/D"){
		str = pathStr + str
	}
	loadLog(str)
	operation()
	fmt.Println("请输入提取行数")
	var in int
	fmt.Scanln(&in)
	printMaxs(in)
	for {
		time.Sleep(time.Second*10000)
	}
}

