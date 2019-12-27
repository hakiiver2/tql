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
            t.Table.SetSelectable(true, true)
        case 'r':
            t.Table.SetSelectable(true, false)
        }
        if event.Key() == tcell.KeyEnter {

        }

        return event;
    })

}