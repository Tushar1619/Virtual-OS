package apps

import (
	switchMode "virtualOS/theme"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Store(){
	window := fyne.CurrentApp().NewWindow("Microsoft Store")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(1100,620))
		img,_ := fyne.LoadResourceFromPath("C:\\Users\\tanya\\OneDrive\\Desktop\\hackathon\\img\\store_light.png")
	window.SetIcon(img)
	window.SetPadded(false)
	imgPage := canvas.NewImageFromFile(fmt.Sprint("C:\\Users\\tanya\\OneDrive\\Desktop\\hackathon\\img/storePage_",switchMode.SwitchMode,".jpg"))
	imgPage.Resize(fyne.NewSize(1100,620))
	window.SetContent(container.NewWithoutLayout(imgPage))

	window.Show()
}