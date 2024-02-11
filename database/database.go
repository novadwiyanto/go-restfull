package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func LoadDB() {
	conn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta")
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

/*
Start For Pagination
*/
func GetDatabaseFields(out interface{}) *schema.Schema {

	s, err2 := schema.Parse(&out, &sync.Map{}, DB.NamingStrategy)
	if err2 != nil {
		panic("failed to parse schema")
	}
	return s

}

type FieldWithType struct {
	DBName string
	Type   string
}

func GetFieldNames(out interface{}) []FieldWithType {

	s := GetDatabaseFields(out)
	var fields []FieldWithType
	for _, field := range s.Fields {

		fields = append(fields, FieldWithType{
			DBName: field.DBName,
			Type:   field.FieldType.Name(),
		})
	}
	return fields

}

/*
End For Pagination
*/

