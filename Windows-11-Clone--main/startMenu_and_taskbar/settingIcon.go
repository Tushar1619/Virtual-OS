package startMenu_and_taskbar

import (
	switchMode "Windows_11_clone/theme"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type settingStateCheck struct {
	widget.BaseWidget
	normalState string
	clickedOnwhichApp string
	clickedOnwhichApp_state string
}

func NewsettingStateCheck(array []string) *settingStateCheck {
	c := &settingStateCheck{normalState:array[0],clickedOnwhichApp: array[1],
		clickedOnwhichApp_state: "none",}
	c.ExtendBaseWidget(c)
	return c
}

func (c *settingStateCheck) Tapped(_ *fyne.PointEvent) {

if c.clickedOnwhichApp_state == "none"{
	c.clickedOnwhichApp_state = "true"
}
	c.Refresh()
}

func (c *settingStateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &settingStateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type settingStateRender struct {
	check *settingStateCheck
	img   *canvas.Image
}

func (t *settingStateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *settingStateRender) Destroy() {}

func (t *settingStateRender) Layout(_ fyne.Size) {
	t.img.Resize(fyne.NewSize(180,125))
}

func (t *settingStateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *settingStateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *settingStateRender) Refresh() {
	t.updateImage()
}

func (t *settingStateRender) updateImage() {
	defer t.img.Refresh()

	switch t.check.clickedOnwhichApp_state=="true" {
	default:
		if t.check.clickedOnwhichApp == "lightMode" && t.check.clickedOnwhichApp_state == "true"{
				switchMode.SwitchMode = "light"
			
		}
		if t.check.clickedOnwhichApp == "darkMode" && t.check.clickedOnwhichApp_state == "true"{
				switchMode.SwitchMode = "dark"
		}

		res,err := fyne.LoadResourceFromPath(t.check.normalState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource = theme.NewThemedResource(res)
	}
}
