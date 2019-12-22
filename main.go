package main

import (
	//"fmt"

	"flag"
	"fmt"

	"github.com/hakiiver2/showcol/dbinfo"
	"github.com/hakiiver2/showcol/tui"
)

func main() {
    var i interface{}
    var (
        UserName = flag.String("username", "", "username")
        PassWord = flag.String("pass", "", "password")
        DbName = flag.String("db", "", "database name")
    )
    flag.Parse()

    info := dbinfo.New().Set(*UserName, *PassWord, *DbName);
    iinfo := info.GetDbInfo();
    fmt.Println(iinfo)
    fmt.Println(info);
    fmt.Println(info.UserName);
    tui.New().Run(i, "mysql", info)
    // if err := tui.New().Run(i); err != nil {
    // }
    //ConnectDB("mysql")
}
