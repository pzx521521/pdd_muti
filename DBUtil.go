package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

func DBSave() {
	t := time.NewTicker(time.Hour)
	defer t.Stop()
	for {
		<-t.C
		ClearEndTimeMsOrder()
		DoDBSave()
	}
}

func ClearEndTimeMsOrder() {
	now := time.Now().Unix() * 1000
	for _, good := range DB {
		for i2 := len(good.Orders) - 1; i2 >= 0; i2-- {
			order := good.Orders[i2]
			if order.EndTimeMs < now {
				good.Orders = append(good.Orders[:i2], good.Orders[i2+1:]...)
			}
		}
	}
}

func DoDBSave() {
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(DB); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	os.Remove("./db.db")
	os.WriteFile("./db.db", buf.Bytes(), 0666)
}

func DBLoad() {
	data, err := os.ReadFile("./db.db")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	// 进行反序列化
	err = decoder.Decode(&DB)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v\n", DB)
}
