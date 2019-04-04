package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()
	//make sure connect is avaliable
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	router := gin.Default()

	router.POST("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")
		stmt, err := db.Prepare("insert into person(first_name,last_name) values(?,?);")
		if err != nil {
			fmt.Println("BugHere")
		}
		res, err := stmt.Exec(first_name, last_name)
		//
		lid, err := res.LastInsertId()   //เอา Last id ที่เพิ่มล่าสุดไป Return เป็น int 64
		id := strconv.FormatInt(lid, 10) // Convert int 64 to base10 string
		//Reference https://yourbasic.org/golang/convert-int-to-string/#int64-to-string
		buffer.WriteString(id) //เขียน
		//
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s Succesfully Created ", name),
		})
	})
	router.GET("/persons", func(c *gin.Context) {
		var (
			person  Person
			persons []Person
		)
		rows, err := db.Query("select * from person")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
			persons = append(persons, person)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(200, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})
	router.DELETE("/person/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		stmt, err := db.Prepare("DELETE FROM PERSON where id=?")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer stmt.Close()
		c.JSON(200, gin.H{
			"result": "Delete Success",
		})
	})
	router.GET("/person", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("Select * from	person where id = ?")
		if err != nil {
			fmt.Println(err.Error())
		}
		rows, err := stmt.Exec(id)

		//In Complete
		//digio place for study
	})
	router.Run(":3000")
}
