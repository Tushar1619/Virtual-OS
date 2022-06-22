package startMenu_and_taskbar

import (
	switchMode "Windows_11_clone/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type settingIconLayout struct{}


	var (
	seetingIcon_global_width float32
	seetingIcon_global_height float32
	seetingIcon_itemData []fyne.CanvasObject
	seetingIcon_global_count int
)


func (tbi_l *settingIconLayout) MinSize(items []fyne.CanvasObject) fyne.Size{
	total := fyne.NewSize(0,0)
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		total = total.Add(obj.MinSize())
	}
	return total
}

func (tbi_l *settingIconLayout) Layout(items []fyne.CanvasObject,size fyne.Size){

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
	
	seetingIcon_global_width = size.Width
	seetingIcon_global_height = size.Height 
	seetingIcon_itemData = items
	seetingIcon_global_count = count
	myPos := fyne.NewPos((seetingIcon_global_width/2.25-5.5)-float32(count*23),size.Height)
			for _,obj := range items{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+190
				
				obj.Resize(fyne.NewSize(180,125))
			}
}

func SeetingIcon_Run_MoveUp(){
	myPos := fyne.NewPos((seetingIcon_global_width/2-167.5)-float32(seetingIcon_global_count*23),(seetingIcon_global_height/2+174))
	if switchMode.SwitchMode == "dark"{
		myPos = fyne.NewPos((seetingIcon_global_width/2-163.5)-float32(seetingIcon_global_count*23),(seetingIcon_global_height/2+172))
	}
					for _,obj := range seetingIcon_itemData{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+190
				
				obj.Resize(fyne.NewSize(180,125))
			}
}
func SeetingIcon_Run_MoveDown(){
	myPos := fyne.NewPos((seetingIcon_global_width/2.25-5.5)+-float32(seetingIcon_global_count*23),(seetingIcon_global_height))
					for _,obj := range seetingIcon_itemData{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+75
				
				obj.Resize(fyne.NewSize(60,60))
			}
}

func SettingIconLayout_UI() fyne.CanvasObject{

	 calc := []string{"D:/Windows_11_clone/img/seeting_screen_1_light.png","lightMode"}
	 codeView := []string{"D:/Windows_11_clone/img/seeting_screen_1_dark.png","darkMode"}

	newContainer := container.New(&settingIconLayout{},
		NewsettingStateCheck(calc),
		NewsettingStateCheck(codeView),
	)
	return newContainer
}