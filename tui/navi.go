package tui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell"
    "strings"
    "fmt"
)

var NaviColor = `[red::b]%s[white]: %s`

var (
    stopApp = fmt.Sprintf(NaviColor, "ctrl-q", "stop")
    defaultNavis = strings.Join([]string{stopApp}, "\n")
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
        if event.Key() == tcell.KeyCtrlQ {
            t.App.Stop();
        }
        return event;
    })

}
