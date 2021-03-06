package main 

import (
	// "fyne.io/fyne"
	"log"
	"fmt"
"fyne.io/fyne/v2"
"io/ioutil" 
// "fyne.io/fyne/v2/app"
"net/http"
"fyne.io/fyne/v2/widget"
"fyne.io/fyne/v2/container"
"encoding/json"
 "image/color"
"fyne.io/fyne/v2/canvas"
)

func showWeatherApp(w fyne.Window) {
  


  //Api Part
 
  res,err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=530f12348c9394c23f12db3a95d10e94")
	if err!=nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body , err:=ioutil.ReadAll(res.Body)
if err!=nil {
	fmt.Println(err)
}


weather , err:= UnmarshalWeather(body)
if err!=nil {
	fmt.Println(err)
}


 img:=canvas.NewImageFromFile("weather.png")
img.FillMode=canvas.ImageFillOriginal


combo:= widget.NewSelect([]string{"Noida","Chennai","Delhi","Jharkhand"},func(value string) {
	log.Println("Places",value)
   
})
label1:= canvas.NewText("Weather Details" ,color.White)
label1.TextStyle = fyne.TextStyle{Bold: true}

label2:= canvas.NewText(fmt.Sprintf("Country %s",weather.Sys.Country),color.White)
label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f",weather.Wind.Speed),color.White)
label4 := canvas.NewText(fmt.Sprintf("Temperature %2f",weather.Main.Temp),color.White)
label5 := canvas.NewText(fmt.Sprint("Humidity ",weather.Main.Humidity),color.White)


	weatherContainer:=container.NewVBox(
		label1,
		img,
		combo,
		label2,
		label3,
		label4,
		label5,
		container.NewGridWithColumns(1,
		),
	)

w.SetContent(container.NewBorder(panelContent,nil,nil,nil,weatherContainer),)
w.Show()
}
	

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()


func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
