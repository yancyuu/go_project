package system

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type db struct {
	dbConnect *sql.DB
}

func (db *db) init() {
	db.dbConnect, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.dbConnect.SetMaxOpenConns(2000)
	db.dbConnect.SetMaxIdleConns(1000)
	db.dbConnect.Ping()
}

func (db *db) getInstance() *sql.DB {
	if db.dbConnect == nil {
		db.init()
		return db.dbConnect
	} else {
		return db.dbConnect
	}
}

//通过结构体的到格式化的字符串
func getAddValueStr(structName interface{}) string {
	var valueStr = ""
	t := reflect.TypeOf(structName)
	for i := 0; i < t.NumField(); i++ {
		valueStr += valueStr + `,?`
	}
	return `(` + valueStr + `)`
}

//插入demo
/*
structName 代表表内容的结构体
addValues 需要插入的数据
*/
func (db *db) insert(structInput interface{}) {
	addFileds := GetStructFieldNames(structInput) //通过结构体获取结构体的名称
	addFiledStr := strings.Join(addFileds, ",")   //通过结构体名称数组得到字符串
	sqlStr := `INSERT user` + `(` + addFiledStr + `) ` + `values ` + getAddValueStr(structInput)
	Logger().WithFields(logrus.Fields{
		"name": "yancy",
	}).Info("mysql insert sql:%v", sqlStr)
	stmt, err := db.dbConnect.Prepare(sqlStr)
	checkErr(err)
	res, err := stmt.Exec(GetStructFields(structInput)...)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	Logger().WithFields(logrus.Fields{
		"name": "yancy",
	}).Info("mysql insert id%v", id)
}

//查询demo
func query() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	//普通demo
	//for rows.Next() {
	//    var userId int
	//    var userName string
	//    var userAge int
	//    var userSex int

	//    rows.Columns()
	//    err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//    checkErr(err)

	//    fmt.Println(userId)
	//    fmt.Println(userName)
	//    fmt.Println(userAge)
	//    fmt.Println(userSex)
	//}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
