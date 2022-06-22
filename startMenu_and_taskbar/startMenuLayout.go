package startMenu_and_taskbar

import (
	"fmt"

	switchMode "Windows_11_clone/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)
var (
	global_width float32
	global_height float32
	itemData fyne.CanvasObject
)
type startMenuLayout struct{}
func ( tl *startMenuLayout) MinSize(items []fyne.CanvasObject) fyne.Size{
	return items[0].MinSize()
}
func (tl *startMenuLayout) Layout(items []fyne.CanvasObject,size fyne.Size){
			items[0].Move(fyne.NewPos(size.Width,size.Height))
		items[0].Resize(fyne.NewSize(size.Width/2.8,620))
		global_width = size.Width
		global_height = size.Height
		itemData = items[0]
}

func RunUpMove(){
				itemData.Move(fyne.NewPos(global_width/3.02,(global_height/9)-10))
}

func RunUpDown(){
				itemData.Move(fyne.NewPos(global_width/3.02,global_height))
}

func StartMenuLayout_UI() fyne.CanvasObject{
	item:=canvas.NewImageFromFile(fmt.Sprint("D:/Windows_11_clone/img/startMenu_",switchMode.SwitchMode,".png"))
	newContainer := container.New(&startMenuLayout{},item)
	return newContainer
}



