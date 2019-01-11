package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "./go-sql-driver"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

const batch_size int = 1000

func Insert(begin int, end int) {
	if begin > end {
		fmt.Println("@Insert -> begin lager than end. function will exit.")
		return
	}

	var err error
	for begin < end+1 {
		start_time := time.Now()
		tx, _ := db.Begin()
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		for i := 1; i < batch_size && begin < end+1; i++ {
			//每次循环用的都是tx内部的连接，没有新建连接，效率高
			str := GetRandomSalt()
			_, err = tx.Exec("INSERT INTO tbl_user(uid, username, age, memo, remark) values(?, ?, ?, ?, ?)", begin, "user"+strconv.Itoa(begin), begin-1000, str, MD5(str))
			if err != nil {
				log.Println("db insert execute error: ", err)
				return
			}
			begin++
		}
		//最后释放tx内部的连接
		tx.Commit()
		end_time := time.Now()
		fmt.Println("insert total time:", end_time.Sub(start_time).Seconds())
	}
}

func Query() {
	var err error
	start_time := time.Now()
	rows, err := db.Query("SELECT uid, username, age, memo, remark FROM tbl_user")
	if err != nil {
		log.Println("db query execute error: ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		var memo string
		var remark string
		if err := rows.Scan(&id, &name, &age, &memo, &remark); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id:%d, name:%s, age:%d, memo:%s, remark: %s.\n", id, name, age, memo, remark)
	}
	end_time := time.Now()
	fmt.Println("query total time:", end_time.Sub(start_time).Seconds())
}

func Update(begin int, end int) {
	if begin > end {
		fmt.Println("@Update -> begin lager than end. function will exit.")
		return
	}

	var err error
	for begin < end+1 {
		start_time := time.Now()
		tx, _ := db.Begin()
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		for i := 1; i < batch_size && begin < end+1; i++ {
			str := GetRandomString(10)
			_, err = tx.Exec("UPDATE tbl_user SET memo=?, remark = ? WHERE uid = ?", str, MD5(str), begin)
			if err != nil {
				log.Println("db update execute error: ", err)
				return
			}
			begin++
		}
		//最后释放tx内部的连接
		tx.Commit()
		end_time := time.Now()
		fmt.Println("update total time:", end_time.Sub(start_time).Seconds())
	}
}

func Delete() {
	var err error
	start_time := time.Now()
	tx, _ := db.Begin()
	_, err = tx.Exec("TRUNCATE TABLE tbl_user")
	if err != nil {
		log.Println("db delete execute error: ", err)
		return
	}
	tx.Commit()
	end_time := time.Now()
	fmt.Println("delete total time:", end_time.Sub(start_time).Seconds())
}
