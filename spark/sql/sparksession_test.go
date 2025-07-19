// Package utils @author: Violet-Eva @date  : 2025/2/11 @notes :
package sql

import (
	"context"
	"fmt"
	"github.com/apache/arrow-go/v18/arrow"
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	//a := "2025-07-17 15:49:59.330808 +0800 CST"
	//parse, err := time.Parse("2006-01-02 15:04:05.999999 Z0700 MST", a)
	a := "2025-07-17 15:49:59"
	parse, err := time.Parse("2006-01-02 15:04:05", a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(parse)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	fromString, _, err2 := arrow.TimestampFromStringInLocation(parse.Format("2006-01-02 15:04:05"), arrow.Second, location)
	if err2 != nil {
		log.Println(err2)
	}
	fmt.Println(fromString)
	fromString1, err3 := arrow.TimestampFromTime(parse, arrow.Second)
	if err3 != nil {
		log.Println(err3)
	}
	fmt.Println(fromString1)
	toTime := fromString1.ToTime(arrow.Second)
	fmt.Println(toTime)
	unix := parse.Unix()
	fmt.Println(unix)
}

func TestSQL(t *testing.T) {
	params := make(map[string]string)
	params["username"] = "violet-eva"
	params["username_auth"] = "5TRyuMpZClX4bSiZ2eAapg"
	// 127.0.0.1
	params["address_auth"] = "liIgrWj6TMHTW9hTiWLYNQ"
	sql, err := SparkConnServer("127.0.0.1", 15002, params)
	if err != nil {
		t.Fatal(err)
	}
	frame, err := sql.Sql(context.Background(), "select cast(\"2025-07-17 15:49:59\" as timestamp) as col")
	if err != nil {
		t.Fatal(err)
	}
	err = frame.Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
	err = frame.Describe(context.Background()).Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
}

type Water struct {
	Word  string    `json:"word"  spark:"word_name"`
	Sale  float32   `json:"sale"  spark:"sale_name" `
	Count int64     `json:"count" spark:"count_name" sparkType:"int8"`
	Date  time.Time `json:"date"  spark:"date_name"  sparkType:"date,2006-01-02 15:04:05"`
	Times time.Time `json:"times" spark:"times_name" sparkType:"timestamp_us,2006-01-02 15:04:05"`
}

func TestDDL(t *testing.T) {
	params := make(map[string]string)
	params["username"] = "violet-eva"
	params["username_auth"] = "5TRyuMpZClX4bSiZ2eAapg"
	params["address_auth"] = "liIgrWj6TMHTW9hTiWLYNQ"
	sql, err := NewSparkSQL("127.0.0.1", 15002, params, 3, 10)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	fmt.Println(1, now)
	var ws []Water
	for i := 0; i < 10; i++ {
		var w Water
		if i%2 == 0 {
			w.Word = fmt.Sprintf("w%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Times = now.AddDate(0, 0, i) //.Format("2006-01-02 15:04:05.000")
		} else {
			w.Word = fmt.Sprintf("t%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Date = now.AddDate(0, 0, i) //.Format("2006-01-02 15:04:05") //.AddDate(0, 0, -i)
		}
		ws = append(ws, w)
	}
	fromStruct, err := sql.CreateDataFrameFromStruct(context.Background(), ws, true)
	if err != nil {
		t.Fatal(err)
	}
	err = fromStruct.Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
	if err = fromStruct.CreateTempView(context.Background(), "test", true, false); err != nil {
		t.Fatal(err)
	}
	dataFrame, err := sql.Sql(context.Background(), "desc test")
	if err != nil {
		t.Fatal(err)
	}
	if err = dataFrame.Show(context.Background(), 100, false); err != nil {
		t.Fatal(err)
	}
	frame, err := sql.Sql(context.Background(), "select cast(times_name as timestamp) as col from test")
	if err != nil {
		t.Fatal(err)
	}
	err = frame.CreateTempView(context.Background(), "test2", true, false)
	if err != nil {
		t.Fatal(err)
	}
	frame, err = sql.Sql(context.Background(), "desc test2")
	if err != nil {
		t.Fatal(err)
	}
	err = frame.Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
}

type W struct {
	Word  string    `json:"word"  spark:"word_name"`
	Sale  float32   `json:"sale"  spark:"sale_name" `
	Count int64     `json:"count" spark:"count_name" sparkType:"int8"`
	Date  time.Time `json:"date"  spark:"date_name"  sparkType:"date,2006-01-02"`
	Times string    `json:"times" spark:"times_name" `
}

func TestDFToMap(t *testing.T) {
	params := make(map[string]string)
	params["username"] = "violet-eva"
	params["username_auth"] = "5TRyuMpZClX4bSiZ2eAapg"
	// 127.0.0.1
	params["address_auth"] = "liIgrWj6TMHTW9hTiWLYNQ"
	sql, err := NewSparkSQL("127.0.0.1", 15002, params, 3, 10)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = sql.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}()
	var ws []W
	for i := 0; i < 10; i++ {
		var w W
		if i%2 == 0 {
			w.Word = fmt.Sprintf("w%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Date = time.Now()
			w.Times = time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
		} else {
			w.Word = fmt.Sprintf("t%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Date = time.Now().AddDate(0, 0, -1)
			w.Times = time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
		}
		ws = append(ws, w)
	}
	frame, err := sql.CreateDataFrameFromStruct(context.Background(), ws, true)
	if err != nil {
		t.Fatal(err)
	}

	err = frame.Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
}
