package startMenu_and_taskbar

import (
	news "Windows_11_clone/news"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CheckState int

const (
	CheckOff CheckState = iota
	CheckOn
)

type StateCheck struct {
	widget.BaseWidget
	State CheckState
	normalState string
	openedState string
	clickedOnWhichApp string
	clickedOnWhichApp_State string

}
var (
PassArray []string
WhichApp string
)

func NewStateCheck(array []string) *StateCheck {
	c := &StateCheck{normalState:array[0],openedState: array[1],clickedOnWhichApp: array[2],clickedOnWhichApp_State: "none"}
	c.ExtendBaseWidget(c)
	return c
}

func (c *StateCheck) Tapped(_ *fyne.PointEvent) {
	if c.State == CheckOn {
		c.State = CheckOff
	} else {
		c.State++
	}

		if c.clickedOnWhichApp_State == "none"{
			c.clickedOnWhichApp_State = "true"
		}else if c.clickedOnWhichApp_State=="true"{
		c.clickedOnWhichApp_State = "false"
		}else if c.clickedOnWhichApp_State == "false"{
			c.clickedOnWhichApp_State = "true"
		}
	c.Refresh()
}

func (c *StateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &StateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type StateRender struct {
	check *StateCheck
	img   *canvas.Image
}

func (t *StateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *StateRender) Destroy() {
}

func (t *StateRender) Layout(_ fyne.Size) {
	t.img.Resize(fyne.NewSize(38,38))
}

func (t *StateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *StateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *StateRender) Refresh() {
	t.updateImage()
}

func (t *StateRender) updateImage() {

	defer t.img.Refresh()
	switch t.check.State {
		case CheckOff:
		if t.check.clickedOnWhichApp == "start" && t.check.clickedOnWhichApp_State == "false"{
			RunUpDown()
			StartMenu_Icon_Run_MoveDown()
		}

		if t.check.clickedOnWhichApp == "news" && t.check.clickedOnWhichApp_State == "false"{
		news.News_RunUpDown()
		}

		if t.check.clickedOnWhichApp == "setting" && t.check.clickedOnWhichApp_State == "false"{
			Setting_RunUpDown()
			SeetingIcon_Run_MoveDown()
		}
		
		res,err := fyne.LoadResourceFromPath(t.check.normalState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource =theme.NewThemedResource(res)
	case CheckOn:
		if t.check.clickedOnWhichApp == "start" && t.check.clickedOnWhichApp_State == "true"{
			RunUpMove()
			StartMenu_Icon_Run_MoveUp()
		}
		if t.check.clickedOnWhichApp == "news" && t.check.clickedOnWhichApp_State == "true"{
		news.News_RunUpMove()
		}
		if t.check.clickedOnWhichApp == "setting" && t.check.clickedOnWhichApp_State == "true"{
			Setting_RunUpMove()
			SeetingIcon_Run_MoveUp()
		}
		res,err := fyne.LoadResourceFromPath(t.check.openedState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource = theme.NewThemedResource(res)
	}
}

