package startMenu_and_taskbar

import (
	switchMode "Windows_11_clone/theme"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type taskbarIcon_layout struct{}

func (tbi_l *taskbarIcon_layout) MinSize(items []fyne.CanvasObject) fyne.Size{
	total := fyne.NewSize(0,0)
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		total = total.Add(obj.MinSize())
	}
	return total
}

func (tbi_l *taskbarIcon_layout) Layout(items []fyne.CanvasObject,size fyne.Size){

	count := -1
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		count++
	}
	if count==1{
		count=0
	}
	
	myPos := fyne.NewPos((size.Width/2)-float32(count*23),size.Height-42.5)
			for _,obj := range items{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+43
				
				obj.Resize(fyne.NewSize(38,38))
			}
}

func TaskbarIcon_layout_UI() fyne.CanvasObject{

	 start := []string{"D:/Windows_11_clone/img/logo2.png",
		fmt.Sprint("D:/Windows_11_clone/img/start_logo_started_",switchMode.SwitchMode,".png"),"start"}

	 setting := []string{"D:/Windows_11_clone/img/settingLogo.png",
		fmt.Sprint("D:/Windows_11_clone/img/settingLogo_started_",switchMode.SwitchMode,".png"),"setting"}

	 news := []string{"D:/Windows_11_clone/img/newsLogo.png",
		fmt.Sprint("D:/Windows_11_clone/img/newsLogo_Started_",switchMode.SwitchMode,".png"),"news"}

	newContainer := container.New(&taskbarIcon_layout{},NewStateCheck(start),NewStateCheck(setting),NewStateCheck(news))
	return newContainer
	
}