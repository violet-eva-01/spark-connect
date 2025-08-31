// Package utils @author: Violet-Eva @date  : 2025/2/11 @notes :
package sql

import (
	"context"
	"fmt"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	// 创建内存分配器
	alloc := memory.NewGoAllocator()

	// 1. 定义 schema - 包含两个字符串数组字段
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "names", Type: arrow.BinaryTypes.String, Nullable: true},
			{Name: "hobbies", Type: arrow.BinaryTypes.String, Nullable: true},
		},
		nil,
	)

	// 2. 创建字符串数组构建器
	namesBuilder := array.NewStringBuilder(alloc)
	defer namesBuilder.Release()

	hobbiesBuilder := array.NewStringBuilder(alloc)
	defer hobbiesBuilder.Release()

	// 3. 向构建器添加数据 (模拟[]string数据)
	names := []string{"Alice", "Bob", "", "Dave"} // 空字符串将被视为null
	hobbies := []string{"reading", "sports", "music", ""}

	for _, name := range names {
		if name == "" {
			namesBuilder.AppendEmptyValue()
		} else {
			namesBuilder.Append(name)
		}
	}

	for _, hobby := range hobbies {
		if hobby == "" {
			hobbiesBuilder.AppendEmptyValue()
		} else {
			hobbiesBuilder.Append(hobby)
		}
	}

	// 4. 构建Arrow数组
	namesArr := namesBuilder.NewStringArray()
	defer namesArr.Release()

	hobbiesArr := hobbiesBuilder.NewStringArray()
	defer hobbiesArr.Release()

	// 5. 创建Record (一条记录包含多个字段的数组)
	record := array.NewRecord(schema, []arrow.Array{namesArr, hobbiesArr}, -1)
	defer record.Release()

	fmt.Println("创建的Record:")
	printRecord(record)

	// 6. 创建Table (可以包含多个Record)
	tbl := array.NewTableFromRecords(schema, []arrow.Record{record})
	defer tbl.Release()

	fmt.Println("\n创建的Table:")
	fmt.Printf("表包含 %d 个字段, %d 行数据\n", tbl.NumCols(), tbl.NumRows())

	// 7. 从表中读取数据
	readTableData(tbl)
}

// 打印Record内容
func printRecord(record arrow.Record) {
	fmt.Printf("Schema: %v\n", record.Schema().Fields())
	fmt.Printf("行数: %d\n", record.NumRows())

	for i, col := range record.Columns() {
		field := record.Schema().Field(i)
		fmt.Printf("字段 %s:\n", field.Name)

		strArr := col.(*array.String)
		for j := 0; j < strArr.Len(); j++ {
			if strArr.IsNull(j) {
				fmt.Printf("  行 %d: <null>\n", j)
			} else {
				fmt.Printf("  行 %d: %s\n", j, strArr.ValueStr(j))
			}
		}
	}
}

// 读取表中的数据并转换为Go的[]string
func readTableData(tbl arrow.Table) {
	// 读取第一个字段(names)
	namesCol := tbl.Column(0)
	namesArr := namesCol.Data().Chunk(0).(*array.String)

	// 转换为Go的[]string
	names := make([]string, namesArr.Len())
	for i := 0; i < namesArr.Len(); i++ {
		if namesArr.IsNull(i) {
			names[i] = ""
		} else {
			names[i] = namesArr.ValueStr(i)
		}
	}

	fmt.Println("\n转换为Go的[]string:")
	fmt.Println("Names:", names)
}

func TestTime(t *testing.T) {
	timeName := "2025-07-22 17"
	parse, err := time.Parse("2006-01-02 15:04:05", timeName)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(parse.UTC())
}

func TestSQL(t *testing.T) {
	params := make(map[string]string)
	params["user"] = "violet-eva"
	params["user_auth"] = "5TRyuMpZClX4bSiZ2eAapg"
	sql, err := NewSparkSQL("127.0.0.1", 15002, params)
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
	Bool          string      `json:"bool"      spark:"column:bool;type:bool"`
	Bools         []bool      `json:"bools"`
	Ints          []int       `json:"ints"`
	Int8s         []int8      `json:"int8s"`
	Int16s        []int16     `json:"int16s"`
	Int32s        []int32     `json:"int32s"`
	Int64s        []int64     `json:"int64s"`
	Float32s      []float32   `json:"float32s"`
	Float64s      []float64   `json:"float64s"`
	Strings       []string    `json:"strings"`
	Timestamps    []time.Time `json:"timestamps" spark:"column:timestamps;type:[]date"`
	TimestampMs   []time.Time `json:"timestamp_ms" spark:"column:timestamp_ms"`
	TimestampMss  []time.Time `json:"timestamp_mss" spark:"column:timestamp_mss;type:[]timestamp_us"`
	ID            string      `json:"id"        spark:"column:id;type:int32"`
	Word          string      `json:"word"      spark:"column:word_name;type:string"`
	Sale          string      `json:"sale"      spark:"column:sale_name;type:float"`
	Count         string      `json:"count"     spark:"column:count_name;type:int8"`
	Values        []string
	Value         string    `json:"value"     spark:"column:value_name;type:[]string;format:,"`
	Date          string    `json:"date"      spark:"column:date_name;type:date"`
	Times         string    `json:"times"     spark:"column:times_name;type:timestamp_us"`
	Timestamp     time.Time `json:"timestamp" spark:"column:timestamp_name;type:timestamp_us"`
	TimestampDate time.Time `json:"timestamp_date" spark:"column:timestamp_name_date;type:date"`
}

func TestDDL(t *testing.T) {
	params := make(map[string]string)
	params["username"] = "violet-eva"
	params["username_auth"] = "5TRyuMpZClX4bSiZ2eAapg"
	params["address_auth"] = "liIgrWj6TMHTW9hTiWLYNQ"
	sql, err := NewSparkSQL("127.0.0.1", 15002, params)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	fmt.Println(1, now)
	var ws []Water
	for i := 0; i < 10; i++ {
		var w Water
		if i%2 == 0 {
			w.Int8s = []int8{1, 2, 3}
			w.Bools = []bool{true, false}
			w.Int32s = []int32{1, 2, 3}
			w.Float32s = []float32{1, 2, 3}
			w.Timestamps = []time.Time{now, now, now}
			w.TimestampMss = []time.Time{now, now, now}
			w.Word = fmt.Sprintf("w%d", i)
			w.Sale = fmt.Sprintf("%d", i)
			w.Count = fmt.Sprintf("%d", i)
			w.Value = "1 2 3"
			w.Times = now.AddDate(0, 0, i).Format("2006-01-02 15:04:05.000000")
			w.Timestamp = now
			w.TimestampDate = now
		} else {
			w.Ints = []int{1, 2, 3}
			w.Int16s = []int16{1, 2, 3}
			w.Int64s = []int64{1, 2, 3}
			w.Float64s = []float64{1, 2, 3}
			w.TimestampMs = []time.Time{now, now, now}
			w.Word = fmt.Sprintf("t%d", i)
			w.Sale = fmt.Sprintf("%d", i)
			w.Count = fmt.Sprintf("%d", i)
			w.Values = []string{"1", "2", "3"}
			w.Date = now.AddDate(0, 0, i).Format("2006-01-02 15:04:05.000") //.AddDate(0, 0, -i)
		}
		ws = append(ws, w)
	}
	fromStruct, err := sql.CreateDataFrameFromStruct(context.Background(), ws)
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
	sql, err := NewSparkSQL("127.0.0.1", 15002, params)
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
	frame, err := sql.CreateDataFrameFromStruct(context.Background(), ws)
	if err != nil {
		t.Fatal(err)
	}

	err = frame.Show(context.Background(), 100, false)
	if err != nil {
		t.Fatal(err)
	}
}
