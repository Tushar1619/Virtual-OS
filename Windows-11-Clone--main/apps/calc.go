package apps

import (
	"fmt"
	"image/color"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func Calc(){
	window := fyne.CurrentApp().NewWindow("Calculator")
	
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(270,430))
		img,_ := fyne.LoadResourceFromPath("D:/Windows_11_clone/img/calcLogo.png")
	window.SetIcon(img)
	btnTypeLabel := widget.NewEntry()
	btnTypeLabel.SetText(" ")

	ansLabel := canvas.NewText("0",color.Black)
	ansLabel.TextSize = 50
	btn0 := widget.NewButton("0",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"0")
	})
	btn1 := widget.NewButton("1",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"1")
	})
	btn2 := widget.NewButton("2",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"2")
	})
	btn3 := widget.NewButton("3",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"3")
	})
	btn4 := widget.NewButton("4",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"4")
	})
	btn5 := widget.NewButton("5",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"5")
	})
	btn6 := widget.NewButton("6",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"6")
	})
	btn7 := widget.NewButton("7",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"7")
	})
	btn8 := widget.NewButton("8",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"8")
	})
	btn9 := widget.NewButton("9",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"9")
	})
	btnDot:=widget.NewButton(".",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+".")
	})
	btnEqual :=widget.NewButton("=",func() {

		expression, _ := govaluate.NewEvaluableExpression(fmt.Sprint(btnTypeLabel.Text));
		parameters := make(map[string]interface{})
			result, _ := expression.Evaluate(parameters);
			ansLabel.Text =fmt.Sprint(result)
			ansLabel.Refresh()
	})
	btnEqual.Importance = widget.HighImportance
	btnAdd :=widget.NewButton("+",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"+")
	})
	btnAdd.Importance = widget.HighImportance
	btnSub :=widget.NewButton("-",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"-")
	})
	btnSub.Importance = widget.HighImportance
	btnMul :=widget.NewButton("*",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"*")
	})
	btnMul.Importance = widget.HighImportance
	btnDiv :=widget.NewButton("/",func() {
		btnTypeLabel.SetText(btnTypeLabel.Text+"/")
	})
	btnDiv.Importance = widget.HighImportance

	btnClear := widget.NewButtonWithIcon("",theme.DeleteIcon(),func() {
		btnTypeLabel.SetText("")
		ansLabel.Text = "0"
		ansLabel.Refresh()
	})
	btnClear.Importance =  widget.HighImportance
	btnSquare := widget.NewButton("x²",func() {
		expression, _ := govaluate.NewEvaluableExpression(fmt.Sprint(btnTypeLabel.Text));
		parameters := make(map[string]interface{})
		result, _ := expression.Evaluate(parameters);
		newValue,_ :=   strconv.Atoi(fmt.Sprint(result))
		ansLabel.Text =fmt.Sprint(newValue*newValue)
		ansLabel.Refresh()
	})
	btnSquareRoot := widget.NewButton("X³",func() {
		expression, _ := govaluate.NewEvaluableExpression(fmt.Sprint(btnTypeLabel.Text));
		parameters := make(map[string]interface{})
		result, _ := expression.Evaluate(parameters);
		newValue,_ :=   strconv.Atoi(fmt.Sprint(result))
		ansLabel.Text =fmt.Sprint(newValue*newValue*newValue)
		ansLabel.Refresh()

	})

	window.SetContent(container.NewVBox(
		layout.NewSpacer(),
		 ansLabel,btnTypeLabel,
			layout.NewSpacer(),
			container.NewGridWrap(
		fyne.NewSize(64.3,55),
		btn7,btn8,btn9,btnDiv,
		btn4,btn5,btn6,btnMul,
		btn1,btn2,btn3,btnSub,
		btnDot,btn0,btnSquare,btnAdd,
		btnSquareRoot,btnClear,btnEqual,
		),layout.NewSpacer(),))
	window.SetPadded(false)
	window.Show()
}