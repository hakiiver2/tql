package tui

import (
    "fmt"
)

type DbInfo struct {
    UserName  string
    PassWord  string
    DbName    string
    TableName string
    FieldName string
    Host      string
    Sql       string
}

var empty = "";

func NewDbInfo() *DbInfo {
    dbinfo := &DbInfo {
        UserName:   empty,
        PassWord:   empty,
        DbName:     empty,
        TableName:  empty,
        FieldName:  empty,
        Host     :  empty,
        Sql:  empty,
    }

    return dbinfo;
}


func (dbinfo *DbInfo) Set (UserName string, PassWord string, DbName string,TableName string, FieldName string, Host string, Sql  string) *DbInfo{

    dbinfo.UserName   = UserName
    dbinfo.PassWord   = PassWord
    dbinfo.DbName     = DbName
    dbinfo.TableName  = TableName
    dbinfo.FieldName  = FieldName
    dbinfo.Host       = "tcp(" + Host + ")"
    dbinfo.Sql        = Sql

    fmt.Println(dbinfo.UserName)

    return dbinfo;
}

func (dbinfo DbInfo) GetDbInfo () DbInfo {
    return dbinfo;
}

