package main

import (
	groupietracker "Groupie"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Artists struct {
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Album        string
	Date         string
	Member       string
	concertDates string
	concert      string
}

type GeoUrl struct {
	lat string
	lon string
}

type LocationsIndex struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Locations []string `json:"locations"`
	Location  string
}

type DatesIndex struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	Dates []string `json:"dates"`
	Date  string
}

type RelationsIndex struct {
	Index []Relations `json:"index"`
}

type Geos struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	Geo string
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
	DatesLocation  string
}

func DataArtists(url string) []Artists {
	var artists []Artists
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&artists)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return artists
}

func DataLocations(url string) LocationsIndex {
	var locations LocationsIndex
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&locations)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return locations
}

func DataDates(url string) DatesIndex {
	var dates DatesIndex
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&dates)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return dates
}

func DataRelations(url string) RelationsIndex {
	var relations RelationsIndex
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&relations)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return relations
}

func GetGeoUrl(url string) []Geos {
	var GeoUrl []Geos
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&GeoUrl)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return GeoUrl
}

func ChechIfCompose(name string, textEntry string) bool {
	name = strings.ToLower(name)
	textEntry = strings.ToLower(textEntry)
	return strings.Contains(name, textEntry)
}

func main() {
	// var geourl GeoUrl
	grid := groupietracker.CreateGrid()
	var artist Artists
	var relations Relations
	var dates Dates
	artists := DataArtists("https://groupietrackers.herokuapp.com/api/artists")
	locationsIndex := DataLocations("https://groupietrackers.herokuapp.com/api/locations")
	relationsIndex := DataRelations("https://groupietrackers.herokuapp.com/api/relation")
	datesIndex := DataDates("https://groupietrackers.herokuapp.com/api/dates")
	contentMembers := widget.NewLabel("")
	contentName := widget.NewLabel("")
	contentCreationDate := widget.NewLabel("")
	contentFirstAlbum := widget.NewLabel("")
	contentconcert := widget.NewLabel("")
	contentRelations := widget.NewLabel("")

	a := app.New()
	w := a.NewWindow("Groupie-Tracker")
	w.Resize(fyne.NewSize(1280, 720))

	searchEntry := widget.NewEntry()
	searchEntry.Resize(fyne.NewSize(300, 35))
	searchButton := widget.NewButton("Search", func() {
		var temp []Artists
		for _, v := range artists {
			if ChechIfCompose(v.Name, searchEntry.Text) {
				temp = append(temp, v)
			}
		}
		result := ""
		for _, v := range temp {
			result += v.Name
		}
		h := dialog.NewInformation("", result, w)
		h.Show()

		fmt.Printf("searchEntry.Text: %v\n", searchEntry.Text)
	})
	searchButton.Resize(fyne.NewSize(100, 35))
	searchButton.Move(fyne.NewPos(500, 1))

	r, _ := fyne.LoadResourceFromURLString("")
	img_grp := canvas.NewImageFromResource(r)

	s, _ := fyne.LoadResourceFromURLString("")
	img_world := canvas.NewImageFromResource(s)

	listArtists := widget.NewList(
		func() int {
			return len(artists)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(artists[id].Name)
		})
	listArtists.Resize(fyne.NewSize(300, 35))

	listArtists.OnSelected = func(id widget.ListItemID) {
		artist.Member = "Membres : "
		for _, m := range artists[id].Members {
			artist.Member += m + "  "
			contentMembers.SetText(artist.Member)
		}
		var city string
		var country string
		var cityTable []string
		var countryTable []string
		for i := range locationsIndex.Index[id].Locations {
			city, country = groupietracker.Split(locationsIndex.Index[id].Locations[i])
			cityTable = append(cityTable, city)
			countryTable = append(countryTable, country)
		}
		fmt.Print(cityTable)
		fmt.Print(countryTable)
		var Geo []Geos
		for i := 0; i <= len(cityTable)-1; i++ {
			lien_geo := "https://nominatim.openstreetmap.org/?addressdetails=1&q=" + cityTable[i] + "+" + countryTable[i] + "&format=json&limit=1"
			fmt.Println(" ")
			fmt.Println(lien_geo)
			Geo = GetGeoUrl(lien_geo)
			fmt.Print(Geo)
		}
		// fmt.Println(GeoIndex)

		lien_grp := artists[id].Image
		r, _ = fyne.LoadResourceFromURLString(lien_grp)
		img_grp.Resource = r
		img_grp.FillMode = canvas.ImageFillOriginal
		img_grp.Refresh()

		lien_world := "https://www.bluewaterbio.com/wp-content/uploads/bwb-world-map.png"
		s, _ = fyne.LoadResourceFromURLString(lien_world)
		img_world.Resource = s
		img_world.FillMode = canvas.ImageFillContain
		img_world.Refresh()

		artist.Name = "Nom du groupe : " + artists[id].Name
		contentName.SetText(artist.Name)

		artist.Date = "Date de création : "
		artist.Date += strconv.Itoa(artists[id].CreationDate)
		contentCreationDate.SetText(artist.Date)

		artist.Album = "Premier album : " + artists[id].FirstAlbum
		contentFirstAlbum.SetText(artist.Album)

		// var listConcerts []string
		var alldates []string
		var e int
		for e = 0; e <= len(datesIndex.Index[id].Dates)-1; e++ {
			if string(datesIndex.Index[id].Dates[e][0]) == "*" {
				alldates = append(alldates, datesIndex.Index[id].Dates[e])
			}
		}
		e = len(alldates) - 1
		fmt.Println(len(alldates))
		dates.Date = "Concerts : "
		for _, l := range locationsIndex.Index[id].Locations {
			// for e := 0; e <= len(datesIndex.Index[id].Dates)-1; e++ {
			dates.Date += l + " : " + alldates[e] + ", "
		}

		// fmt.Println(dates.Date)

		contentconcert.SetText(dates.Date)

		var c int
		var listLocations []string
		relations.DatesLocation = "Relations : "
		for _, l := range locationsIndex.Index[id].Locations {
			for c = 0; c <= len(relationsIndex.Index[id].DatesLocations[l])-1; c++ {
				var isDouble bool = false
				for _, r := range listLocations {
					// fmt.Println("List : ", r)
					// fmt.Println("Locations : ", l)
					if r == l {
						isDouble = true
					}
				}
				if !isDouble {
					relations.DatesLocation += l + " : " + relationsIndex.Index[id].DatesLocations[l][c] + " "
					listLocations = append(listLocations, l)
				} else {
					relations.DatesLocation += ", " + relationsIndex.Index[id].DatesLocations[l][c] + " "
				}
				contentRelations.SetText(relations.DatesLocation)
			}
		}
		// fmt.Println(s)
		// fmt.Println("nombre : ", c)
	}

	contentName.Wrapping = fyne.TextWrapWord
	contentMembers.Wrapping = fyne.TextWrapWord
	contentCreationDate.Wrapping = fyne.TextWrapWord
	contentFirstAlbum.Wrapping = fyne.TextWrapWord
	contentconcert.Wrapping = fyne.TextWrapWord
	contentRelations.Wrapping = fyne.TextWrapWord

	grids := container.NewWithoutLayout(
		grid,
	)

	concertbox := container.NewVSplit(
		img_world,
		contentconcert,
	)

	info := container.NewVBox(
		img_grp,
		contentName,
		contentMembers,
		contentCreationDate,
		contentFirstAlbum,
		contentRelations,
	)

	gen := container.NewWithoutLayout(
		info,
		concertbox,
	)

	// listartists := listArtists

	gens := container.NewWithoutLayout(
		gen,
		grids,
	)

	split4 := container.NewWithoutLayout(
		listArtists,
		gens,
	)

	infosearch := container.NewHBox(
		searchEntry,
		searchButton,
	)

	split := container.NewVSplit(
		infosearch,
		split4,
	)
	fmt.Println(len(artists))
	gens.Move(fyne.NewPos(200, 15))
	grids.Move(fyne.NewPos(200, 15))
	listArtists.Resize(fyne.NewSize(200, 950))
	gen.Resize(fyne.NewSize(1700, 1500))
	info.Resize(fyne.NewSize(500, 950))
	concertbox.Resize(fyne.NewSize(1200, 950))
	concertbox.Move(fyne.NewPos(500, 15))
	split.Offset = 0
	w.Resize(fyne.NewSize(1920, 1020))
	w.CenterOnScreen()
	w.SetContent(split)
	w.ShowAndRun()
}
