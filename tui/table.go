package tui

import (
    "strings"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell"
    "database/sql"
	"github.com/hakiiver2/showcol/dbinfo"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)


func CreateTable(tui *Tui, info *dbinfo.DbInfo) tview.Primitive {
    table := tview.NewTable().SetFixed(1, 1);
    fmt.Println("Reading...");

    for row, line := range strings.Split(getColumns(info), "\n") {
        for column, cell := range strings.Split(line, "|") {
            color := tcell.ColorWhite
            if row == 0{
                color = tcell.ColorYellow
            } else if column == 0 {
                color = tcell.ColorDarkCyan
            }
            align := tview.AlignLeft
            if row == 0 {
                align = tview.AlignCenter
            } else if column == 0 || column >= 4 {
                align = tview.AlignRight
            }
            tableCell := tview.NewTableCell(cell).
                SetTextColor(color).
                SetAlign(align).
                SetSelectable(row != 0 && column != 0)
            if column >= 1 && column <= 3 {
                tableCell.SetExpansion(1)
            }
            table.SetCell(row, column, tableCell)
        }
    }

    table.SetBorder(true).SetTitle("table")

    return table;
}

func getColumns(info *dbinfo.DbInfo) string {
    db, err := sql.Open("mysql", info.UserName + ":@/" + info.DbName)
    if err != nil {
        panic(err);
    }

    r, err := db.Query("SELECT * FROM review LIMIT 10000");
    if err != nil {
        panic(err);
    }
    columns, err := r.Columns()
    scanArgs := make([]interface{}, len(columns))
    values   := make([]interface{}, len(columns))

    for i := range values {
        scanArgs[i] = &values[i]
    }
    var tt string = "|"
    for _, col := range columns {
        tt += col + "|"
    }

    tt += "\n"
    for r.Next() {
        r.Scan(scanArgs...)
        var entry string = "|"
        for i, _ := range columns {
            val := values[i]
            b, ok := val.([]byte)
            if ok {
                entry += string(b) + "|"
            } else {
                entry += "|"
            }
        }
        tt += entry
        tt += "\n\n"
    }

    return tt;
}
