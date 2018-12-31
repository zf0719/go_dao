package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "./go-sql-driver"
)

func Update() {
	var err error
	start := time.Now()
	// Begin 内部会去获取连接
	tx, err := db.Begin()
	if err != nil {
		log.Println("db update begin error: ", err)
		return
	}
	for i := 1301; i <= 1400; i++ {
		_, err = tx.Exec("update user set age=? where uid=?", i, i)
		if err != nil {
			log.Println("db update execute error: ", err)
			return
		}
	}
	tx.Commit()
	end := time.Now()
	fmt.Println("update total time:", end.Sub(start).Seconds())
}

func Delete() {
	var err error
	start := time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		_, err = tx.Exec("DELETE FROM USER WHERE uid=?", i)
		if err != nil {
			log.Println("db delete execute error: ", err)
			return
		}
	}
	tx.Commit()
	end := time.Now()
	fmt.Println("delete total time:", end.Sub(start).Seconds())
}

func Query() {
	var err error
	start := time.Now()
	rows, err := db.Query("SELECT uid,username FROM USER")
	if err != nil {
		log.Println("db query execute error: ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end := time.Now()
	fmt.Println("query total time:", end.Sub(start).Seconds())
}

func Insert() {
	var err error
	start := time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		_, err = tx.Exec("INSERT INTO user(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
		if err != nil {
			log.Println("db insert execute error: ", err)
			return
		}
	}
	//最后释放tx内部的连接
	tx.Commit()
	end := time.Now()
	fmt.Println("insert total time:", end.Sub(start).Seconds())
}
