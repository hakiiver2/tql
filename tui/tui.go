package tui

import (
    "database/sql"
    "fmt"
    "github.com/rivo/tview"
	"github.com/hakiiver2/tql/dbinfo"
    //"strconv"
    _ "github.com/go-sql-driver/mysql"
)

type Tui struct {
    List   *tview.List
    Pages  *tview.Pages
    Navi   *Navi
    Table  *tview.Table
    Modal  *tview.Modal
    Layout *tview.Flex
    App    *tview.Application
    Mode   string
}

type mydb struct {
    username  string
    password  string
    dbname    string
    fieldname string
}

func New() *Tui{
    tui := &Tui{
        List:  tview.NewList(),
        Pages: tview.NewPages(),
        Navi : NewNavi(),
        Table: tview.NewTable(),
        Modal: tview.NewModal(),
        Layout: tview.NewFlex(),
        Mode  : "row",
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

    tui.SetKeyBind();


    //tui.Pages.AddAndSwitchToPage("tableList", textview, true);
    tui.CreateTable(info);
    //layout := tview.NewFlex().
        tui.Layout.
        SetDirection(tview.FlexRow).
        AddItem(tui.Table, 0, 1, true).
        AddItem(tui.Navi, 1, 1, false)
    tui.Pages.AddAndSwitchToPage("tableList", tui.Layout, true);
    if err := tui.App.SetRoot(tui.Pages, true).Run(); err != nil {
        panic(err);
    }

}


