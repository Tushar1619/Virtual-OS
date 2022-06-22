package newsLayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)
var (
	news_global_width float32
	news_global_height float32
	news_itemData fyne.CanvasObject
)
type newsLayout struct{}

func ( nl *newsLayout) MinSize(items []fyne.CanvasObject) fyne.Size{
	return items[0].MinSize()
}

func (nl *newsLayout) Layout(items []fyne.CanvasObject,size fyne.Size){
			items[0].Move(fyne.NewPos(size.Width,size.Height))
		items[0].Resize(fyne.NewSize(size.Width/4.5,600))
		news_global_width = size.Width
		news_global_height = size.Height
		news_itemData = items[0]
}

func News_RunUpMove(){
				news_itemData.Move(fyne.NewPos(news_global_width/2-150,(news_global_height/7.3)-10))
}

func News_RunUpDown(){
				news_itemData.Move(fyne.NewPos(news_global_width,news_global_height))
}

func NewsLayout_UI() fyne.CanvasObject{
	item:=canvas.NewImageFromFile("D:/Windows_11_clone/img/newsLayout.png")
	newContainer := container.New(&newsLayout{},item)
	return newContainer
}