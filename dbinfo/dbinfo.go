package dbinfo

import (
    "fmt"
)

type DbInfo struct {
    UserName string
    PassWord string
    DbName   string
}

var dbInfo *DbInfo
var empty = "";

func New() *DbInfo {
    dbinfo := &DbInfo {
        UserName: empty,
        PassWord: empty,
        DbName:   empty,
    }

    return dbinfo;
}

func SetDbInfo (UserName string, PassWord string, DbName string) *DbInfo{

    dbInfo.UserName = UserName;
    dbInfo.PassWord = PassWord;
    dbInfo.DbName   = DbName;

    fmt.Println(dbInfo.UserName)

    return dbInfo;
}


func (dbinfo *DbInfo) Set (UserName string, PassWord string, DbName string) *DbInfo{

    dbinfo.UserName = UserName;
    dbinfo.PassWord = PassWord;
    dbinfo.DbName   = DbName;

    fmt.Println(dbinfo.UserName)

    return dbinfo;
}

func (dbinfo DbInfo) GetDbInfo () DbInfo {
    return dbinfo;
}

func GetDbInfo () *DbInfo {
    return dbInfo;
}
