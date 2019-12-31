package tui

import (
    // "github.com/rivo/tview"
    _ "fmt"

    "github.com/gdamore/tcell"
    _ "github.com/go-sql-driver/mysql"
    "github.com/hakiiver2/tql/dbinfo"
)

func (tui *Tui) InitCmdLine(info *dbinfo.DbInfo) {

    tui.CmdLine.SetLabel("sql")
    tui.CmdLine.SetDoneFunc(func(key tcell.Key) {
        frontPageName, _ := tui.Pages.GetFrontPage()
        if frontPageName == "cmdline" {
            if key == tcell.KeyEnter {
                sql := tui.CmdLine.GetText()
                info.Sql = sql
                tui.Table.Clear()
                tui.CreateTable(info);

            }else if key == tcell.KeyEscape {

            }
            tui.Layout.RemoveItem(tui.CmdLine)
            tui.Layout.RemoveItem(tui.Table)

            tui.Pages.RemovePage("cmdline")
            tui.Layout.AddItem(tui.Table, 0, 1, true)
            tui.Layout.AddItem(tui.Navi, 1, 1, false)
            tui.Table.SetSelectable(true, false)

            tui.Pages.AddAndSwitchToPage("tableList", tui.Layout, true)
            tui.App.SetFocus(tui.Pages)
        }
    })
}

func (tui *Tui) CmdLineMode(frontPageName string) {

    tui.Layout.RemoveItem(tui.Navi)
    tui.Layout.RemoveItem(tui.Table)
    tui.Pages.RemovePage(frontPageName)
    tui.Pages.RemovePage("tableList")

    tui.Layout.AddItem(tui.Table, 0, 1, false)
    tui.Layout.AddItem(tui.CmdLine, 1, 1, true)

    tui.Pages.AddAndSwitchToPage("cmdline", tui.Layout, true)
}

