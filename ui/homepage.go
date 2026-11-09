package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/supabase-community/supabase-go"
)

var choises []string = []string{"תוכנה", "מכניקה", "חשמל וחיווט", "מידול", "סושואל", "עיצוב", "תיאוד"}

func Homepage(w fyne.Window) *fyne.Container {
	app := fyne.CurrentApp()

	API_KEY := "sb_publishable_qqKz79NENPxNF0YHEFzRHA_d5oxyfTp"
	URL := "https://tnhadzvbozjzcbnrpada.supabase.co"

	NewClient, err := supabase.NewClient(URL, API_KEY, nil)
	if err != nil {
		dialog.ShowError(err, w)
		w.Close()
	}
	FromField := widget.NewEntry()

	list := widget.NewList(func() int {
		return len(choises)
	},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Button).SetText(choises[lii])
			co.(*widget.Button).OnTapped = func() {
				name := FromField.Text
				if strings.TrimSpace(name) == "" {
					dialog.ShowError(fmt.Errorf("Empty name!"), w)
					return
				}
				group := choises[lii]
				window := app.NewWindow("Window " + group)
				window.Resize(w.Canvas().Size())
				window.SetContent(Msg_Group(window, group, NewClient, name))
				window.Show()
			}
		})
	NewContainer := container.NewBorder(nil, FromField, nil, nil, list)
	return NewContainer
}
