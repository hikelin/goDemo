package mysql

import (
	"bytes"
	"database/sql"
	"reflect"
	"strconv"
	"strings"

	"demo/persistence/mysql/models"
	"demo/utils"
)

// Operator structor
type Operator struct {
	db *sql.DB
	tx *sql.Tx
}

// CreateOperator will create operator
func CreateOperator(db *sql.DB, tx *sql.Tx) Operator {
	return Operator{db: db, tx: tx}
}

// Insert model into database
func (operator Operator) Insert(model models.Model) {
	log := utils.GetLogger()
	log.Print("start")

	var fieldList []string
	var valueList []interface{}
	var preList []string
	mt := reflect.TypeOf(model)
	vt := reflect.ValueOf(model)

	for i := 0; i < mt.NumField(); i++ {
		field := mt.Field(i)
		value := vt.Field(i)

		if fieldName, ok := field.Tag.Lookup("db"); ok {
			if fieldName != "" && value.IsValid() {
				vstr := convertValue(value)
				if vstr != "" {
					fieldList = append(fieldList, fieldName)
					preList = append(preList, "?")
					valueList = append(valueList, vstr)
				}
			}
		}
	}

	var pBuffer bytes.Buffer
	pBuffer.WriteString("INSERT INTO example (")
	pBuffer.WriteString(strings.Join(fieldList, ", "))
	pBuffer.WriteString(") VALUES( ")
	pBuffer.WriteString(strings.Join(preList, ", "))
	pBuffer.WriteString(")")
	sqlStr := pBuffer.String()
	stmtIns, err := operator.db.Prepare(sqlStr)

	if err != nil {
		log.Panic(err)
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	log.Print(sqlStr)
	log.Print(valueList)
	_, err = operator.tx.Stmt(stmtIns).Exec(valueList...)

	if err != nil {
		log.Panic(err)
	}
	log.Print("end")
}

func convertValue(val reflect.Value) string {
	log := utils.GetLogger()
	var re string

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		re = strconv.FormatInt(val.Int(), 10)
		break
	case reflect.String:
		re = val.String()
		break
	default:
		log.Fatal("Unsupported type.")
	}

	return re
}
