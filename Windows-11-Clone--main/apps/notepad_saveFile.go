package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Notepad_saveFile(){
	screenLabel := widget.NewMultiLineEntry()
	window := fyne.CurrentApp().NewWindow("Notepad - New File")
	img,_ := fyne.LoadResourceFromPath("D:/Windows_11_clone/img/notepad.png")
	window.SetIcon(img)

	header := fyne.NewMainMenu(
		fyne.NewMenu(
			"Menu",
			fyne.NewMenuItem("Save File",func() {
				fileSave_dialog := dialog.NewFileSave(func(r fyne.URIWriteCloser, _ error) {
						data := []byte(screenLabel.Text)
						r.Write(data)
				},window)
				fileSave_dialog.SetFileName("TypeAnyFileName.txt")
				fileSave_dialog.Show()
			}),
		),
	)
	window.SetMainMenu(header)
	window.SetContent(screenLabel)
	window.Resize(fyne.NewSize(800,600))
	window.Show()
}