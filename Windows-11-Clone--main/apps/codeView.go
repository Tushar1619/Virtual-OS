package apps

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func CodeViewApp(){
	window := fyne.CurrentApp().NewWindow("CodeView")
	res,_ := fyne.LoadResourceFromPath("D:/Windows_11_clone/img/codeView.png")
	window.SetIcon(res)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800,600))

	Entry := widget.NewEntry()
	screenData := widget.NewLabel("")
	Entry.PlaceHolder = "Enter here URI"
	popUp_codeView := func ()  {
		dialog.ShowForm("Code View Box","Confirm","Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Enter URI",Entry),
		},
		func(confirm bool) {
			if !confirm {
				return
			}

			u,err := storage.ParseURI(Entry.Text)
			if err != nil {
				dialog.ShowError(err,window)
			}
			r,err := storage.Reader(u)
				if err != nil {
				dialog.ShowError(err,window)
			}
			defer r.Close()

			data,err := ioutil.ReadAll(r)
			if err != nil {
				dialog.ShowError(err,window)
			}
			screenData.SetText(string(data))
		},
		window,
	)
	}

	header := fyne.NewMainMenu(
		fyne.NewMenu(
			"File",
			fyne.NewMenuItem("Enter URL",popUp_codeView),
		),
	)

	window.SetContent(container.NewScroll(screenData))
	window.SetMainMenu(header)
	window.Show()
}