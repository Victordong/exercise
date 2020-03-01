package golang

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type Crop struct {
	Name    string `gorm:"type:varchar(32);not null"`
	IconUrl string `gorm:"type:varchar(511);not null"`
	Note    string `gorm:"type:TEXT"`
	ShopId  uint
}

func camelString(columnName string) string {
	data := make([]byte, 0, len(columnName)*2)
	j := false
	for i := 0; i < len(columnName); i++ {
		d := columnName[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func main() {
	var crop Crop
	var ignore interface{}
	selectedMap := make(map[int]string)
	cropValue := reflect.ValueOf(&crop).Elem()
	cropType := reflect.TypeOf(&crop).Elem()
	for i := 0; i < cropType.NumField(); i++ {
		fmt.Println(cropType.Field(i).Type)
	}
	dbSQL, err := sql.Open("mysql", "root:wshwoaini@/auto_fertilizer?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	if rows, err := dbSQL.Query("select * from crops where id = ?", 10); err == nil {
		columns, _ := rows.Columns()
		for rows.Next() {
			values := make([]interface{}, len(columns))
			for index, column := range columns {
				values[index] = &ignore
				_, ok := cropType.FieldByName(camelString(column))
				if ok {
					cropFieldValue := cropValue.FieldByName(camelString(column))
					values[index] = cropFieldValue.Addr().Interface()
					selectedMap[index] = camelString(column)
				}
			}
			err = rows.Scan(values...)
			if err != nil {
				fmt.Println(err)
			}
			for index, name := range selectedMap {
				v := reflect.ValueOf(values[index])
				for v.Kind() == reflect.Ptr {
					v = v.Elem()
					fmt.Println(name, " value is ptr")
				}
				cropFieldValue := cropValue.FieldByName(name)
				cropFieldValue.Set(v)
			}
		}
		fmt.Println(crop)
	} else {
		fmt.Println(err)
	}
}
