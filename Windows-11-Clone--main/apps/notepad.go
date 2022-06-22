package apps

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func NoteApp(){
	window := fyne.CurrentApp().NewWindow("Notepad")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800,600))
	screenLabel := widget.NewMultiLineEntry()
	screenLabel.SetText("")
	file_dialog := dialog.NewFileOpen( func(r fyne.URIReadCloser, err error) {
		if err!=nil{
			return
		}
		data ,err := ioutil.ReadAll(r)
		if err!=nil{
			return
		}

		result := fyne.NewStaticResource("Data",data)
	screenLabel.SetText(string(result.StaticContent))

	},window)

	header := fyne.NewMainMenu(
		fyne.NewMenu(
			"Open New File",
			fyne.NewMenuItem("Load Text File",func() {
				file_dialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
				file_dialog.Show()
			}),
			fyne.NewMenuItem("Save Loaded File",func() {
						fileSave_dialog := dialog.NewFileSave(func(r fyne.URIWriteCloser, _ error) {
						data := []byte(screenLabel.Text)
						r.Write(data)
				},window)
				fileSave_dialog.SetFileName("saveWithNewName.txt")
				fileSave_dialog.Show()
			}),
		),
			fyne.NewMenu(
			"Create New File",
			fyne.NewMenuItem("Text File",func() {
				Notepad_saveFile()
			}),
		),
	)

	window.SetMainMenu(header)
	window.SetContent(screenLabel)
	img,_ := fyne.LoadResourceFromPath("D:/Windows_11_clone/img/notepad.png")
	window.SetIcon(img)

window.Show()
}