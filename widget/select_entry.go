package widget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

// SelectEntry is an input field which supports selecting from a fixed set of options.
type SelectEntry struct {
	Entry
	dropDown *fyne.Menu
	popUp    *PopUpMenu
}

// NewSelectEntry creates a SelectEntry.
func NewSelectEntry(options []string) *SelectEntry {
	e := &SelectEntry{}
	e.ExtendBaseWidget(e)
	e.SetOptions(options)
	return e
}

// MinSize returns the minimal size of the select entry.
//
// Implements: fyne.Widget
func (e *SelectEntry) MinSize() fyne.Size {
	min := e.Entry.MinSize()

	if e.dropDown != nil {
		for _, item := range e.dropDown.Items {
			itemMin := fyne.MeasureText(item.Label, theme.TextSize(), fyne.TextStyle{}).Add(fyne.NewSize(4*theme.Padding(), 0))
			min = min.Max(itemMin)
		}
	}
	return min
}

// Resize changes the size of the select entry.
//
// Implements: fyne.Widget
func (e *SelectEntry) Resize(size fyne.Size) {
	e.Entry.Resize(size)
	if e.popUp != nil {
		e.popUp.Resize(fyne.NewSize(size.Width-theme.Padding()*2, e.popUp.Size().Height))
	}
}

// SetOptions sets the options the user might select from.
func (e *SelectEntry) SetOptions(options []string) {
	var items []*fyne.MenuItem
	for _, option := range options {
		option := option // capture
		items = append(items, fyne.NewMenuItem(option, func() { e.SetText(option) }))
	}
	e.dropDown = fyne.NewMenu("", items...)
	dropDownButton := NewButton("", func() {
		c := fyne.CurrentApp().Driver().CanvasForObject(e.super())

		entryPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(e.super())
		popUpPos := entryPos.Add(fyne.NewPos(theme.Padding(), e.Size().Height))

		e.popUp = newPopUpMenu(e.dropDown, c)
		e.popUp.ShowAtPosition(popUpPos)
		e.popUp.Resize(fyne.NewSize(e.Size().Width-theme.Padding()*2, e.popUp.MinSize().Height))
	})
	dropDownButton.HideShadow = true
	dropDownButton.SetIcon(theme.MenuDropDownIcon())
	e.ActionItem = dropDownButton
}
