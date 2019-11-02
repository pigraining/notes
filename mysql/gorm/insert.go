package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	url := "root:sql@tcp(10.249.168.92:3306)/zeus_test1?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	db, err := gorm.Open("mysql", url)
	//sql := "INSERT INTO `action_t` (`domain`,`category`,`action`) VALUES "
	sql := "INSERT INTO `action_t` (`domain`,`category`,`action`,`create_time`,`update_time`) VALUES "
	fmt.Println(err)

	T := GetCSTTime()
	//sql += fmt.Sprintf("(%s,'%s','%s'),", "1", "user", "edit")
	//sql += fmt.Sprintf("(%s,'%s','%s');", "1", "user", "read")
	sql += fmt.Sprintf("(%s,'%s','%s','?','?');", "1", "user", "edit")
	//sql += fmt.Sprintf("(%s,'%s','%s','?','?');", "1", "user", "read")
	fmt.Println(sql)
	//valus := []time.Time{T, T, T, T}
	rew := db.Exec(sql, T, T)
	if err := rew.Error; err != nil {
		fmt.Println(err)
	}
}

func GetCSTTime() time.Time {
	now := time.Now()
	var local, _ = time.LoadLocation("Asia/Shanghai")
	return now.In(local)
}
