package ui

import (
	"encoding/json"
	"fmt"
	"msg/objects"
	"slices"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
)

type Message struct {
	ID   string `json:"id,omitempty"`
	At   string `json:"At"`
	From string `json:"From"`
	Text string `json:"Text"`
}

func GetTextFromGroupWithWin(w fyne.Window, name string, client supabase.Client) []Message {
	data, _, err2 := client.From("msg").Select("id,At,From,Text", "", false).Eq("At", name).Execute()
	if err2 != nil {
		dialog.ShowError(err2, w)
		return nil
	}
	msges := []Message{}
	err3 := json.Unmarshal(data, &msges)
	if err3 != nil {
		return nil
	}
	return msges
}

func UpdateMasseges(w fyne.Window, name string, client *supabase.Client, scrollElement *objects.Scroll, from string) []Message {
	msges := GetTextFromGroupWithWin(w, name, *client)
	scrollElement.ResetContainer(container.NewVBox())
	if msges == nil {
		return []Message{}
	}
	for _, msg := range msges {
		text := fmt.Sprintf("%v: %v", msg.From, msg.Text)
		Label := widget.NewLabel(text)

		if msg.From == from {
			Delete_button := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				client.From("msg").Delete("", "").Eq("id", msg.ID).Execute()
				UpdateMasseges(w, name, client, scrollElement, from)
			})
			Delete_button.Importance = widget.DangerImportance
			cont := container.NewHSplit(Label, Delete_button)
			cont.Offset = 0.9
			scrollElement.AddElement(cont)

			continue
		}
		scrollElement.AddElement(Label)

	}
	scrollElement.Refresh()
	return msges
}

func Msg_Group(w fyne.Window, name string, client *supabase.Client, from string) *fyne.Container {
	var msges []Message
	setion, err := client.SignInWithEmailPassword("yuval@robotics-club.invalid", "123456")
	if err != nil {
		dialog.ShowError(err, w)
		w.Close()

		return container.NewVBox()
	}

	client.UpdateAuthSession(setion)
	client.EnableTokenAutoRefresh(setion)

	cont := container.NewVBox()
	ScrollObj := objects.NewScroll(cont)
	msges = UpdateMasseges(w, name, client, ScrollObj, from)
	go func() {
		last := msges // copy made before the goroutine starts
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			newMsgs := GetTextFromGroupWithWin(w, name, *client)
			if newMsgs == nil || slices.Equal(newMsgs, last) {
				continue
			}
			last = newMsgs
			fyne.Do(func() {
				UpdateMasseges(w, name, client, ScrollObj, from)
				ScrollObj.Refresh()
			})
		}
	}()
	TextSend := widget.NewEntry()
	sendBtn := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {

		massage_insert := Message{At: name, From: from, Text: TextSend.Text, ID: uuid.NewString()}
		_, _, err2 := client.From("msg").Insert(massage_insert, false, "", "", "").Execute()
		if err2 != nil {
			dialog.ShowError(err2, w)
			return
		}
		UpdateMasseges(w, name, client, ScrollObj, from)
	})

	sendBtn.Importance = widget.HighImportance
	ContatinerText := container.NewHSplit(TextSend, sendBtn)
	ContatinerText.SetOffset(0.9)
	NewContainer := container.NewBorder(nil, ContatinerText, nil, nil, ScrollObj.ScrollContainer())

	return NewContainer
}
