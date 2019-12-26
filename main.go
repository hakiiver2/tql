package main

import (
	//"fmt"

	"flag"
	"fmt"

	"github.com/hakiiver2/tql/dbinfo"
	"github.com/hakiiver2/tql/tui"
)

func main() {
    var i interface{}
    var (
        UserName = flag.String("username", "", "username")
        PassWord = flag.String("pass", "", "password")
        DbName = flag.String("db", "", "database name")
        TableName = flag.String("table", "", "table name")
    )
    flag.Parse()

    info := dbinfo.New().Set(*UserName, *PassWord, *DbName, *TableName);
    iinfo := info.GetDbInfo();
    fmt.Println(iinfo)
    fmt.Println(info);
    fmt.Println(info.UserName);
    tui.New().Run(i, "mysql", info)
    // if err := tui.New().Run(i); err != nil {
    // }
    //ConnectDB("mysql")
}
