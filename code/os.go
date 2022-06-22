package main

import (
	// "fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	// switchMode "virtualOS/theme"
	// news "virtualOS/news"
	// startMenu_and_taskbar "virtualOS/startMenu_and_taskbar"
 )

var myApp fyne.App =app.New()

var myWindow fyne.Window= myApp.NewWindow("Tushar's OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject;
var DeskBtn fyne.Widget




var panelContent *fyne.Container

func main() {
	myApp := app.New()
	
	// myApp.Settings().SetTheme(theme.LightTheme())
	// if switchMode.SwitchMode == "dark"{
		myApp.Settings().SetTheme(theme.DarkTheme())
	// }
	// myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("background.jpg")
	res,_ := fyne.LoadResourceFromPath("C:\\Users\\tanya\\OneDrive\\Desktop\\hackathon\\img\\logo.png")
	myWindow.SetIcon(res)


	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func(){
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ComputerIcon(), func(){
		showCalc(myWindow)
	})

	btn3 = widget.NewButtonWithIcon("Text Editor", theme.DocumentCreateIcon(), func(){
		textEditor(myWindow)
	})

	btn4 = widget.NewButtonWithIcon("Gallery", theme.GridIcon(), func(){
		gallery(myWindow)
	})



	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func(){
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

panelContent = container.NewVBox((container.NewGridWithColumns(5,DeskBtn,btn1,btn2,btn3,btn4)))


myWindow.Resize(fyne.NewSize(1200,680))
myWindow.CenterOnScreen();

myWindow.SetContent(
	container.NewBorder(panelContent,nil,nil,nil,img),
)



myWindow.ShowAndRun()
}
