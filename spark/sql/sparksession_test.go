// Package utils @author: Violet-Eva @date  : 2025/2/11 @notes :
package sql

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Water struct {
	Word  string    `json:"word" spark:"word_name"`
	Sale  float32   `json:"sale" spark:"sale_name"`
	Count int64     `json:"count" spark:"count_name"`
	Date  time.Time `json:"date" spark:"date_name"`
	Times string    `json:"times" spark:"times_name" type:"timestamp"`
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
	var ws []Water
	for i := 0; i < 10; i++ {
		var w Water
		if i%2 == 0 {
			w.Word = fmt.Sprintf("w%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Date = time.Now()
			w.Times = time.Now().Format("2006-01-02")
		} else {
			w.Word = fmt.Sprintf("t%d", i)
			w.Sale = float32(i)
			w.Count = int64(i)
			w.Date = time.Now().AddDate(0, 0, -1)
			w.Times = time.Now().Format("2006-01-02 15:04:05.000")
		}
		ws = append(ws, w)
	}
	frame, err := sql.CreateDataFrameFromStruct(context.Background(), ws, true)
	if err != nil {
		t.Fatal(err)
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancelFunc()
	/*	name, err := frame.DropByName(timeout, "word_name", "sale_name", "count_name")
		if err != nil {
			t.Fatal(err)
		}*/
	err = frame.Show(timeout, 100, false)
	if err != nil {
		t.Fatal(err)
	}
}
