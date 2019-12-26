package tui

import (
    "strings"
    "strconv"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell"
    "database/sql"
	"github.com/hakiiver2/tql/dbinfo"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "os"
)


func CreateTable(tui *Tui, info *dbinfo.DbInfo) tview.Primitive {
    table := tview.NewTable().SetFixed(1, 1);
    fmt.Println("Reading...");

    offset := 0;
    skip := 100;
    max_row := skip;

    for row, line := range strings.Split(getRows(info, offset, max_row, "init"), "\n") {
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


    table.SetBorder(true).SetTitle("table");
    table.SetSelectable(true, false).
        SetSeparator(' ');
    table.SetSelectionChangedFunc(func(row int, col int){
            file, err := os.Create("lololo")
            if err != nil {
            }
            defer file.Close()

            b := []byte(strconv.Itoa(max_row))
            file.Write(b)
            if max_row == row {
                offset, max_row = offset + skip, skip + max_row;
                for row, line := range strings.Split(getRows(info, offset, max_row, "add"), "\n") {
                    cur_row := row + offset;
                    for column, cell := range strings.Split(line, "|") {
                        color := tcell.ColorWhite
                        if column == 0 {
                            color = tcell.ColorDarkCyan
                        }
                        align := tview.AlignLeft
                        if column == 0 || column >= 4 {
                            align = tview.AlignRight
                        }
                        tableCell := tview.NewTableCell(cell).
                        SetTextColor(color).
                        SetAlign(align).
                        SetSelectable(cur_row != 0 && column != 0)
                        if column >= 1 && column <= 3 {
                            tableCell.SetExpansion(1)
                        }
                        b := []byte(line);
                        file.Write(b)
                        table.SetCell(cur_row, column, tableCell)
                    }
                }
            }
    })

    return table;
}

func getRows(info *dbinfo.DbInfo, skip int, max int, f_type string) string {
    db, err := sql.Open("mysql", info.UserName + ":@/" + info.DbName)
    if err != nil {
        panic(err);
    }
    skip_string := strconv.Itoa(skip)
    max_string := strconv.Itoa(max)

    r, err := db.Query("SELECT * FROM " + info.TableName +" LIMIT " + max_string + " OFFSET " + skip_string);
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
    if f_type == "init" {
        for _, col := range columns {
            tt += col + "|"
        }
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
        tt += "\n"
    }

    return tt;
}
