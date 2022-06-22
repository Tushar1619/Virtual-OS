package apps

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ToDo(){

	type TodoContent struct{
		Title string
		Task string
		Priority string
	}

	var todoContentArr []TodoContent
	
	data_from_file,_ := ioutil.ReadFile("odo_JSON_Data.txt")

	json.Unmarshal(data_from_file,&todoContentArr)
	window := fyne.CurrentApp().NewWindow("ToDo")
		res,_ := fyne.LoadResourceFromPath("D:/Windows_11_clone/img/todo.png")
	window.SetIcon(res)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800,600))




	content_title := widget.NewEntry()
	content_title.SetPlaceHolder("Enter Title Here...")

	content_task := widget.NewMultiLineEntry()
	content_task.SetPlaceHolder("Enter task Here...")

	content_priority := widget.NewSelect([]string{"Low","Medium","High"},func(s string) {})
	content_priority.PlaceHolder = "Select Priority"
	

	btn_Submit := widget.NewButtonWithIcon("Add Task",theme.ConfirmIcon(),func() {
		myStoreData := &TodoContent{
			Title: content_title.Text,
			Task: content_task.Text,
			Priority: content_priority.Selected,
		}
		
		todoContentArr = append(todoContentArr, *myStoreData)

		 finalData,_:=json.MarshalIndent(todoContentArr,""," ")

		ioutil.WriteFile("Todo_JSON_Data.txt",finalData,0644)

		content_title.Text = ""
		content_task.Text = ""
		content_priority.ClearSelected()
		content_priority.Refresh()
		content_title.Refresh()
		content_task.Refresh()

	})


	btn_Delete := widget.NewButtonWithIcon("Delete Current Task",theme.DeleteIcon(),func() {

		var TempData []TodoContent

		for _,e := range todoContentArr{
			if content_title.Text != e.Title{
				TempData = append(TempData, e)
			}
		}

			todoContentArr = TempData
			result,_ := json.MarshalIndent(todoContentArr,""," ")
			ioutil.WriteFile("Todo_JSON_Data.txt",result,0644)
			content_title.Text = ""
		content_task.Text = ""
		content_priority.ClearSelected()
		content_title.Refresh()
		content_task.Refresh()


	})
	
		list := widget.NewList(
			func() int {return len(todoContentArr)},
			func() fyne.CanvasObject {return widget.NewLabel("")},
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(todoContentArr[lii].Title)
			},
		)

		list.OnSelected = func(id widget.ListItemID) {
			content_title.SetText(todoContentArr[id].Title)
			content_title.Refresh()
			content_task.SetText(todoContentArr[id].Task)
			content_task.Refresh()
			content_priority.SetSelected(todoContentArr[id].Priority)
			content_priority.Refresh()
		}
		btn_update := widget.NewButtonWithIcon("Update Current Task",theme.UploadIcon(),func() {

		var TempData []TodoContent

		update := &TodoContent{
				Title: content_title.Text,
			Task: content_task.Text,
			Priority: content_priority.Selected,
		}

		for _,e := range todoContentArr{
				if content_title.Text == e.Title {
					TempData = append(TempData, *update)
				}else{
					TempData = append(TempData, e)
				}
		}	

			todoContentArr = TempData
			result,_ := json.MarshalIndent(todoContentArr,""," ")
			ioutil.WriteFile("Todo_JSON_Data.txt",result,0644)

		// content_title.Text = ""
		// content_task.Text = ""
		// content_priority.ClearSelected()
		content_title.Refresh()
		content_task.Refresh()
		content_priority.Refresh()
		list.Refresh()
	})


	bottom_rightContainer := container.NewHBox(btn_Submit,btn_Delete,btn_update)

	top_rightContainer := container.NewVBox(content_title,content_priority)

	rightContainer := container.NewBorder(top_rightContainer,bottom_rightContainer,nil,nil,content_task) 

	
	splitContainer := container.NewHSplit(list,rightContainer)
	splitContainer.SetOffset(0.35)
	window.SetContent(splitContainer)
	window.Show()
}