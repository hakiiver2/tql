package tui

import (
    "strings"
    "strconv"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell"
    "database/sql"
	// "github.com/hakiiver2/tql/dbinfo"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "os"
)


var (
    offset int = 0
    skip int = 100
    max_row int = skip
)


func (t *Tui) CreateTable() {

    t.Table.SetFixed(1, 1);
    fmt.Println("Reading...");

    t.SetTable(offset, max_row, "init")


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
                t.SetTable(offset, max_row, "add")
            }
    })

}

func getRows(skip int, max int, f_type string) string {
    db, err := sql.Open("mysql", dbinfo.UserName + ":" + dbinfo.PassWord + "@" + dbinfo.Host + "/" + dbinfo.DbName)
    if err != nil {
        panic(err);
    }
    skip_string := strconv.Itoa(skip)
    max_string := strconv.Itoa(max)

    var r *sql.Rows;
    if dbinfo.Sql != "" {
        r, err = db.Query(dbinfo.Sql + " LIMIT " + max_string + " OFFSET " + skip_string);
    }else{
        r, err = db.Query("SELECT " + dbinfo.FieldName + " FROM " + dbinfo.TableName +" LIMIT " + max_string + " OFFSET " + skip_string);
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

func (tui *Tui) SetTable(offset int, max_row int, setType string) {
    for row, line := range strings.Split(getRows(offset, max_row, setType), "\n") {
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
            tui.Table.SetCell(row, column, tableCell)
        }
    }
}

func (tui *Tui) EditTable() {
    col_max := tui.Table.GetColumnCount()
    //rowInfo := make([]string, 0)


    var id string
    // var selectedRow int
    for i := 0; i < col_max; i++ {
        if i != 0 && i != col_max - 1 {
            row, _ := tui.Table.GetSelection();
            // selectedRow = row
            field_name := tui.Table.GetCell(0, i);
            cell := tui.Table.GetCell(row, i);
            //rowInfo = append(rowInfo, cell.Text)
            tui.EditForm.AddInputField(field_name.Text, cell.Text, 20, nil, nil)
            if field_name.Text == "id" {
                id = cell.Text
            }
        }
    }
    saveToDB := func() {
        count := tui.EditForm.GetFormItemCount()
        db, err := sql.Open("mysql", dbinfo.UserName + ":@/" + dbinfo.DbName)
        if err != nil {
            panic(err);
        }
        defer db.Close()
        for i := 0; i < count; i++ {
            item := tui.EditForm.GetFormItem(i)
            switch item.(type) {
            case *tview.InputField:
                input, _ := item.(*tview.InputField)
                label := item.GetLabel()
                text := input.GetText()
                update_sql := "UPDATE " + dbinfo.TableName + " SET " + label + " = ? WHERE id = ?"
                db.Exec(update_sql, text, id)
                // cell := tview.NewTableCell(text)
                // tui.Table.SetCell(selectedRow, i, cell)
            }
        }
        tui.reloadTable()
        tui.Pages.SwitchToPage("tableList")
        tui.Pages.RemovePage("editForm")
        tui.EditForm.Clear(true)
    }
    tui.EditForm.AddButton("Save", saveToDB)
    tui.EditForm.AddButton("Quit", nil)
    tui.Pages.AddAndSwitchToPage("editForm", tui.EditForm, true)

}

func (tui *Tui) InsertRow() {
    col_max := tui.Table.GetColumnCount()
    for i := 0; i < col_max; i++ {
        if i != 0 && i != col_max - 1 {
            field_name := tui.Table.GetCell(0, i);
            tui.EditForm.AddInputField(field_name.Text, "", 20, nil, nil)
        }
    }
    saveToDB := func() {
        count := tui.EditForm.GetFormItemCount()
        labelList := make([]string, 0)
        valueList := make([]string, 0)
        db, err := sql.Open("mysql", dbinfo.UserName + ":@/" + dbinfo.DbName)
        if err != nil {
            panic(err);
        }
        defer db.Close()
        for i := 0; i < count; i++ {
            item := tui.EditForm.GetFormItem(i)

            switch item.(type) {
            case *tview.InputField:
                input, _ := item.(*tview.InputField)
                label := item.GetLabel()
                text := input.GetText()
                if label != "" {
                    labelList = append(labelList, label)
                    valueList = append(valueList, text)
                }
            }
        }
        labels := strings.Join(labelList, ",")
        values := strings.Join(valueList, "\",\"")
        values = "\"" + values + "\""

        _, err = db.Exec("INSERT INTO " + dbinfo.TableName + " ( " + labels + " ) " + " VALUES" + " ( " + values + " ) ")
        if err != nil {
            panic(err)
        }

        tui.reloadTable()
        tui.Pages.SwitchToPage("tableList")
        tui.Pages.RemovePage("insertForm")
        tui.EditForm.Clear(true)
    }
    tui.EditForm.AddButton("Save", saveToDB)
    tui.EditForm.AddButton("Quit", nil)
    tui.Pages.AddAndSwitchToPage("insertForm", tui.EditForm, true)
}

func (tui *Tui) reloadTable() {
    tui.SetTable(offset, max_row, "init")
}
