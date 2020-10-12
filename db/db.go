package db

import "github.com/astaxie/beego/orm"

type OrmDBWithError struct {
	OrmDB
	Error error
}

type OrmDB interface {
	QueryTable(name string) orm.QuerySeter
	Where(query interface{}, args ...interface{}) OrmDB
	Scan(dest interface{}) *OrmDBWithError
	All()
}
