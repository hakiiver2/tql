package tui

import (
    "database/sql"
    "fmt"
    "github.com/rivo/tview"
	"github.com/hakiiver2/showcol/dbinfo"
    //"strconv"
    _ "github.com/go-sql-driver/mysql"
)

type Tui struct {
    List   *tview.List
    Pages  *tview.Pages
    App    *tview.Application
}

type mydb struct {
    username string
    password string
    dbname   string
}

func New() *Tui{
    tui := &Tui{
        List:  tview.NewList(),
        Pages: tview.NewPages(),
        App:   tview.NewApplication(),
    }
    return tui
}

func (tui *Tui) Run (i interface{}, DB string, info *dbinfo.DbInfo) {
    fmt.Println(info)
    ConnectDB(tui, DB, info);
}

func ConnectDB(tui *Tui, DB string, info *dbinfo.DbInfo){

    fmt.Println(info)

    db, err := sql.Open(DB, info.UserName + ":@/" + info.DbName)
    if err != nil {
        tui.App.Stop()
    }
    defer db.Close()

    res, err := db.Query("SHOW TABLES");
    if err != nil {
        panic(err);
    }

    var tableName string
    textview := tview.NewTextView().
        SetDynamicColors(true).
        SetWrap(false).
        SetRegions(true);

    for res.Next() {
        res.Scan(&tableName);
        fmt.Fprint(textview, tableName + "\n");
    }

    texts := tview.NewTextView().
        SetDynamicColors(true).
        SetRegions(true).
        SetWrap(false);
    fmt.Fprintf(texts, `hello`)

    //tui.Pages.AddAndSwitchToPage("tableList", textview, true);
    table := CreateTable(tui, info);
    layout := tview.NewFlex().
        SetDirection(tview.FlexRow).
        AddItem(table, 0, 1, true).
        AddItem(texts, 1, 1, false)
    tui.Pages.AddAndSwitchToPage("tableList", layout, true);
    if err := tui.App.SetRoot(tui.Pages, true).Run(); err != nil {
        panic(err);
    }

}


