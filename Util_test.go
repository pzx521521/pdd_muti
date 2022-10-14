package main

import (
	"PDD_Muti/data"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"testing"
)

func Test_analsysHtml(t *testing.T) {
	group_order_id := int64(2007363467928373305)
	order := &data.Order{OrderID: group_order_id}
	good, err := GetGoodInfo(order)
	fmt.Printf("%v\n", good)
	fmt.Printf("%v\n", order)
	fmt.Printf("%v\n", err)
}

func Test_DBSaveLoad(t *testing.T) {
	var m = make(map[int64]*data.Good)
	var m2 = make(map[int64]*data.Good)
	m[1] = &data.Good{1, "2", "3", "3",
		[]data.Order{data.Order{99, 99, 99, 99}}}

	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(m); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	os.Remove("./db.db")
	os.WriteFile("./db.db", buf.Bytes(), 0666)

	data, err := os.ReadFile("./db.db")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	// 进行反序列化
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v\n", m)
}
