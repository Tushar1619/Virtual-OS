package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func gallery(w fyne.Window) {

	// w.Resize(fyne.NewSize(800 , 400))

	root_src := "C:\\Users\\tanya\\OneDrive\\Pictures"

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }

	var picsArr []string;
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension := strings.Split(file.Name() , ".")[1]
			if extension == "png" || extension == "jpeg"{
				picsArr = append(picsArr, root_src +"\\" + file.Name())
			}
		}
    }
	
	galleryContainer := container.NewAppTabs(
		container.NewTabItem("Image1" , canvas.NewImageFromFile(picsArr[0])),
	)
	for i := 1 ; i < len(picsArr) ; i++ {
		image := canvas.NewImageFromFile(picsArr[i])
		image.FillMode = canvas.ImageFillOriginal
		galleryContainer.Append(container.NewTabItem("Image" + strconv.Itoa(i+1) , image )) 
	}
	galleryContainer.SetTabLocation(container.TabLocationLeading)
	// image := canvas.NewImageFromFile(picsArr[0])
	// w.SetContent(tabs);
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,galleryContainer),)
	w.Show()
}
