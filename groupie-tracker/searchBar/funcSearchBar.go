package searchbar

import (
	"Groupie/data"
	dataload "Groupie/dataLoad"
	"image/color"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
)

// looks if the string is equal with a other string
func CheckIfCompose(name string, textEntry string) bool {
	name = strings.ToLower(name)
	textEntry = strings.ToLower(textEntry)
	return strings.Contains(name, textEntry)
}

// displays input suggestions as you write
func AutoCompletion(entry *fynex.CompletionEntry) {

	entry.OnChanged = func(s string) {

		if len(s) < 1 {
			entry.HideCompletion()
			return
		}

		resultArtists := dataload.DataArtists("https://groupietrackers.herokuapp.com/api/artists")
		resultLocations := dataload.DataLocations("https://groupietrackers.herokuapp.com/api/locations")
		var resultName []string

		doubleCompletation := false

		for _, r := range resultArtists {

			for _, rn := range resultName {
				if rn == r.Name {
					doubleCompletation = true
				}
			}
			if CheckIfCompose(r.Name, entry.Text) {
				if !doubleCompletation {
					resultName = append(resultName, r.Name)
				}
			}
			for _, m := range resultArtists[r.Id-1].Members {
				if CheckIfCompose(m, entry.Text) {
					resultName = append(resultName, r.Name+" (Membres)")
				}
			}
			for _, l := range resultLocations.Index[r.Id-1].Locations {
				if CheckIfCompose(l, entry.Text) {
					resultName = append(resultName, r.Name+" (Locations)")
				}
			}
			if CheckIfCompose(strconv.Itoa(r.CreationDate), entry.Text) {
				resultName = append(resultName, r.Name+" (Création Date)")
			}
			if CheckIfCompose(r.FirstAlbum, entry.Text) {
				resultName = append(resultName, r.Name+" (Premier Album)")
			}
		}

		if len(resultName) == 0 {
			entry.HideCompletion()
			return
		}

		entry.SetOptions(resultName)
		entry.ShowCompletion()
	}
}

// Creates a new button with the set label (search) and tap handler (looks if the entry put by the user corresponds to one of the data of the artist or group)
func CreateButton(listArtists *widget.List, result *[]string, temp *[]data.Artists, searchOn *bool, searchEntry *fynex.CompletionEntry, artists []data.Artists, v data.Artists, obj *canvas.Rectangle) *widget.Button {
	searchButton := widget.NewButton("Search", func() {
		red := color.NRGBA{R: 255, G: 88, B: 88, A: 200}
		blue := color.NRGBA{R: 133, G: 88, B: 255, A: 200}
		canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
			obj.FillColor = c
			canvas.Refresh(obj)
		}).Start()

		listArtists.UnselectAll()
		*result = nil
		*temp = nil
		*searchOn = false
		entry := strings.Split(searchEntry.Text, " (")

		for _, v = range artists {
			if CheckIfCompose(v.Name, entry[0]) {
				*searchOn = true
				*temp = append(*temp, v)
			} else {
				if CheckIfCompose(v.Name, searchEntry.Text) {
					*searchOn = true
					*temp = append(*temp, v)
				}
			}
		}
		for _, v := range *temp {
			*result = append(*result, v.Name)
		}
	},
	)
	return searchButton
}
