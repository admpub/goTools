package main

import (
	Bock "./sqlBock/Mysql"
	"fmt"
)

func main()  {
	table := Bock.Bock{
		Table: "role",
	}
	var params map[string] interface{}
	args := make(map[string] interface{})
	//db := make([]DB.IBock, 1)
	//db[0] = &table
	//rs, _ := db[0].Retrieve("", params, nil, nil)
	//db := &table
	//rs, _ := db.Retrieve("afataea", values)
	fields := []string {"id", "name"}
	args["fields"] = fields
	db := &table
	rs := db.Retrieve(params, args)
	fmt.Println(rs)
}
