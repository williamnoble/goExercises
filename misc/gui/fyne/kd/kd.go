package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type enterEntry struct {
	widget.Entry
}

func (e *enterEntry) onEnter() {
	fmt.Println(e.Entry.Text)
	e.Entry.SetText("")
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	default:
		e.Entry.KeyDown(key)
	}

}

func main() {
	a := app.New()
	w := a.NewWindow("Messenger")

	entry := newEnterEntry()

	w.SetContent(entry)
	w.ShowAndRun()
}
