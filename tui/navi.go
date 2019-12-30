package tui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell"
    "strings"
    "fmt"
)

var NaviColor = `[red::b]%s[white]: %s`

var (
    cellMode = fmt.Sprintf(NaviColor, "c", "cell mode")
    rowMode  = fmt.Sprintf(NaviColor, "r", "row mode")
    editMode  = fmt.Sprintf(NaviColor, "e", "e mode")
    stopApp  = fmt.Sprintf(NaviColor, "q", "stop")
    defaultNavis = strings.Join([]string{cellMode, rowMode, editMode, stopApp}, "  ")
)

var (
    modal_stopApp  = fmt.Sprintf(NaviColor, "q", "quit")
    modalNavis = strings.Join([]string{modal_stopApp}, "  ")
)

type Navi struct {
    *tview.TextView
}

func NewNavi() *Navi {
    textview := tview.NewTextView().
        SetDynamicColors(true).
        SetWrap(false).
        SetRegions(true);

        navi := &Navi{TextView: textview}
        return navi;
}

func (t *Tui) SetNaviText(pageName string) {
    if pageName == "tableList" {
        t.Navi.SetText(modalNavis)
    }else if pageName == "modal" {
        t.Navi.SetText(defaultNavis)
    }
}

func (t *Tui)SetKeyBind () {
    t.Navi.SetText(defaultNavis)
    t.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        // if event.Key() == tcell.KeyCtrlQ {
        //     t.App.Stop();
        // }
        // if event.Rune() == 'c' {
        //     t.App.Stop();
        // }
        frontPageName, _ := t.Pages.GetFrontPage();
        switch event.Rune() {
        case 'q':
            t.App.Stop();
        case 'c':
            if frontPageName == "tableList" {
                t.Mode = "cell";
                t.Table.SetSelectable(true, true)
            }
        case 'r':
            if frontPageName == "tableList" {
                t.Mode = "row";
                t.Table.SetSelectable(true, false)
            }
        case 'e':
            if frontPageName == "tableList" {
                 t.EditTable()
            }
        }
        if event.Key() == tcell.KeyEnter {
            if frontPageName == "tableList" || frontPageName == "modal"{
                t.SetNaviText(frontPageName);
                var nextPageName string
                if frontPageName == "tableList" {
                    row, col := t.Table.GetSelection();
                    textArr := make([]string, 0)
                    if t.Mode == "cell"{
                        cell := t.Table.GetCell(row, col);
                        t.Modal.SetText(cell.Text);

                    }else if t.Mode == "row" {
                        col_max := t.Table.GetColumnCount();
                        for i := 0; i < col_max; i++ {
                            row, _ := t.Table.GetSelection();
                            cell := t.Table.GetCell(row, i);
                            textArr = append(textArr, cell.Text)
                        }
                        t.Modal.SetText(strings.Join(textArr, "\n"))

                    }
                    t.Layout.RemoveItem(t.Table)
                    t.Layout.AddItem(t.Modal, 0, 1, true)
                    nextPageName = "modal"
                } else {
                    t.Layout.RemoveItem(t.Modal)
                    t.Layout.AddItem(t.Table, 0, 1, true)
                    nextPageName = "tableList"
                }
                t.Layout.RemoveItem(t.Navi)
                t.Pages.RemovePage(frontPageName)

                t.Layout.AddItem(t.Navi, 1, 1, false)
                t.Pages.AddAndSwitchToPage(nextPageName, t.Layout, true)
            }
        }
        if event.Key() == tcell.KeyCtrlR {
            //t.Table.InsertRow()
        }

        return event;
    })

}


