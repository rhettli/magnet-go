package mdb

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	cLocalHost = "118.25.10.97"
	cLocalUser = "root"
	cLocalPsw  = ""	//"root"
	cLocalPort = "3306"
)

//获取数据库链接
func GetLocalDB() (*gorm.DB, error) {
	sqllink := fmt.Sprintf("%s:%s@tcp(%s:%s)/ssbc?charset=utf8&parseTime=True&loc=Local", cLocalUser, cLocalPsw, cLocalHost, cLocalPort)

	db, err := gorm.Open("mysql", sqllink)
	if err != nil {
		return nil, errors.New("Connect to local mysql error:" + err.Error()+",sqlLink:"+sqllink)
	}

	return db, nil
}
