package startMenu_and_taskbar

import (
	apps "Windows_11_clone/apps"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MenuStateCheck struct {
	widget.BaseWidget
	normalState string
	clickedOnwhichApp string
	clickedOnwhichApp_state string
}

func NewMenuStateCheck(array []string) *MenuStateCheck {
	c := &MenuStateCheck{normalState:array[0],clickedOnwhichApp: array[1],
		clickedOnwhichApp_state: "none",}
	c.ExtendBaseWidget(c)
	return c
}

func (c *MenuStateCheck) Tapped(_ *fyne.PointEvent) {

if c.clickedOnwhichApp_state == "none"{
	c.clickedOnwhichApp_state = "true"
}
	c.Refresh()
}

func (c *MenuStateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &MenuStateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type MenuStateRender struct {
	check *MenuStateCheck
	img   *canvas.Image
}

func (t *MenuStateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *MenuStateRender) Destroy() {}

func (t *MenuStateRender) Layout(_ fyne.Size) {
	t.img.Resize(fyne.NewSize(60,60))
}

func (t *MenuStateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *MenuStateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *MenuStateRender) Refresh() {
	t.updateImage()
}

func (t *MenuStateRender) updateImage() {
	defer t.img.Refresh()

	switch t.check.clickedOnwhichApp_state=="true" {
	default:
		
		if t.check.clickedOnwhichApp == "codeView" && t.check.clickedOnwhichApp_state == "true"{
			apps.CodeViewApp()
			RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}

		if t.check.clickedOnwhichApp == "notepad" && t.check.clickedOnwhichApp_state == "true"{
			apps.NoteApp()
				RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}

		if t.check.clickedOnwhichApp == "calc" && t.check.clickedOnwhichApp_state == "true"{
			apps.Calc()
				RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}
		if t.check.clickedOnwhichApp == "weather" && t.check.clickedOnwhichApp_state == "true"{
			apps.WeatherFunc()
				RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}
		if t.check.clickedOnwhichApp == "store" && t.check.clickedOnwhichApp_state == "true"{
			apps.Store()
				RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}
		if t.check.clickedOnwhichApp == "todo" && t.check.clickedOnwhichApp_state == "true"{
			apps.ToDo()
				RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}
		res,err := fyne.LoadResourceFromPath(t.check.normalState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource = theme.NewThemedResource(res)
	}
}

