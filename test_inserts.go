package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	url := "root:sql@tcp(10.249.168.92:3306)/zeus_test1?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	db, _ := gorm.Open("mysql", url)
	sql := "INSERT INTO `action` (`domain`,`category`,`action`,`create_time`,`update_time`) VALUES "
	time := GetCSTTime()
	sql += fmt.Sprintf("(%s,'%s','%s','%s','%s');", "1", "user", "read", time, time)
	sql += fmt.Sprintf("(%s,'%s','%s','%s','%s'),", "1", "user", "edit", time, time)
	db.Exec(sql)
}

func GetCSTTime() time.Time {
	now := time.Now()
	var local, _ = time.LoadLocation("Asia/Shanghai")
	return now.In(local)
}
