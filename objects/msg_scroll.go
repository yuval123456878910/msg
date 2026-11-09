package objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Scroll struct {
	container        *fyne.Container
	scroll_container *container.Scroll
}

func NewScroll(context *fyne.Container) *Scroll {
	return &Scroll{context, container.NewScroll(context)}
}

func (s *Scroll) ScrollContainer() *container.Scroll {
	return s.scroll_container
}

func (s *Scroll) AddElement(obj fyne.CanvasObject) {
	s.container.Add(obj)
}

func (s *Scroll) ResetContainer(context *fyne.Container) {
	s.container.Objects = nil // context is ignored on purpose
	s.container.Refresh()
}

func (s *Scroll) Refresh() {
	s.container.Refresh()
	s.scroll_container.Refresh()
}

func (s *Scroll) ScrollToBottom() {
	s.scroll_container.ScrollToBottom()
}
