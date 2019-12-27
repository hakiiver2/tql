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
    stopApp  = fmt.Sprintf(NaviColor, "q", "stop")
    defaultNavis = strings.Join([]string{cellMode, rowMode, stopApp}, "  ")
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

func (t *Tui)SetKeyBind () {
    t.Navi.SetText(defaultNavis)
    t.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        // if event.Key() == tcell.KeyCtrlQ {
        //     t.App.Stop();
        // }
        // if event.Rune() == 'c' {
        //     t.App.Stop();
        // }
        switch event.Rune() {
        case 'q':
            t.App.Stop();
        case 'c':
            t.Mode = "cell";
            t.Table.SetSelectable(true, true)
        case 'r':
            t.Mode = "row";
            t.Table.SetSelectable(true, false)
        }
        if event.Key() == tcell.KeyEnter {
            row, col := t.Table.GetSelection();
            textArr := make([]string, 0)
            if t.Mode == "cell"{
                cell := t.Table.GetCell(row, col);
                t.Modal.SetText(cell.Text);

                t.Layout.RemoveItem(t.Table);
                t.Layout.AddItem(t.Modal, 0, 1, true);
                t.Layout.AddItem(t.Navi, 1, 1, false)
                t.Pages.AddAndSwitchToPage("modal", t.Layout, true);
            }else if t.Mode == "row" {
                col_max := t.Table.GetColumnCount();
                for i := 0; i < col_max; i++ {
                    row, _ := t.Table.GetSelection();
                    cell := t.Table.GetCell(row, i);
                    textArr = append(textArr, cell.Text)
                }
                t.Modal.SetText(strings.Join(textArr, "\n"));

                t.Layout.RemoveItem(t.Table);
                t.Layout.RemoveItem(t.Navi);
                t.Layout.AddItem(t.Modal, 0, 1, true);
                t.Layout.AddItem(t.Navi, 1, 1, false)
                t.Pages.AddAndSwitchToPage("modal", t.Layout, true);
            }
        }

        return event;
    })

}
