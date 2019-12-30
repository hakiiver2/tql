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


func (t *Tui) CreateTable(info *dbinfo.DbInfo) {
    t.Table.SetFixed(1, 1);
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
            t.Table.SetCell(row, column, tableCell)
        }
    }


    t.Table.SetBorder(true).SetTitle("table");
    t.Table.SetSelectable(true, false).
        SetSeparator(' ');
    t.Table.SetSelectionChangedFunc(func(row int, col int){
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
                        t.Table.SetCell(cur_row, column, tableCell)
                    }
                }
            }
    })

}

func getRows(info *dbinfo.DbInfo, skip int, max int, f_type string) string {
    db, err := sql.Open("mysql", info.UserName + ":@/" + info.DbName)
    if err != nil {
        panic(err);
    }
    skip_string := strconv.Itoa(skip)
    max_string := strconv.Itoa(max)

    var r *sql.Rows;
    if info.Sql != "" {
        r, err = db.Query(info.Sql + " LIMIT " + max_string + " OFFSET " + skip_string);
    }else{
        r, err = db.Query("SELECT " + info.FieldName + " FROM " + info.TableName +" LIMIT " + max_string + " OFFSET " + skip_string);
    }
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

func (tui *Tui) EditTable() {
    col_max := tui.Table.GetColumnCount()
    //rowInfo := make([]string, 0)

    for i := 0; i < col_max; i++ {
        row, _ := tui.Table.GetSelection();
        field_name := tui.Table.GetCell(0, i);
        cell := tui.Table.GetCell(row, i);
        //rowInfo = append(rowInfo, cell.Text)
        tui.EditForm.AddInputField(field_name.Text, cell.Text, 20, nil, nil)
    }
    tui.EditForm.AddButton("Save", nil)
    tui.EditForm.AddButton("Quit", nil)
    tui.Pages.AddAndSwitchToPage("editForm", tui.EditForm, true)

}
