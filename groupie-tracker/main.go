package main

import (
	"Groupie/data"
	dataload "Groupie/dataLoad"
	"Groupie/list"
	searchbar "Groupie/searchBar"
	"image/color"

	fynex "fyne.io/x/fyne/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//Create a new application with a window
	a := app.New()
	w := a.NewWindow("Groupie-Tracker")
	w.Resize(fyne.NewSize(1920, 1020))
	w.CenterOnScreen()

	b := app.New()
	w2 := b.NewWindow("Géoloc")
	w2.Resize(fyne.NewSize(900, 562))

	var artist data.Artists
	var relations data.Relations
	var searchOn bool
	var v data.Artists
	var temp []data.Artists
	var result []string

	//Load data
	artists := dataload.DataArtists("https://groupietrackers.herokuapp.com/api/artists")
	locationsIndex := dataload.DataLocations("https://groupietrackers.herokuapp.com/api/locations")
	relationsIndex := dataload.DataRelations("https://groupietrackers.herokuapp.com/api/relation")

	//creates variables to add content that can be displayed by fyne
	obj := canvas.NewRectangle(color.NRGBA{R: 255, G: 88, B: 88, A: 200})
	obj.Resize(fyne.NewSize(30, 50))
	r, _ := fyne.LoadResourceFromURLString("")
	img_grp := canvas.NewImageFromResource(r)
	s, _ := fyne.LoadResourceFromURLString("")
	img_world := canvas.NewImageFromResource(s)
	contentName := widget.NewLabel("")
	contentMembers := widget.NewLabel("")
	contentCreationDate := widget.NewLabel("")
	contentFirstAlbum := widget.NewLabel("")
	contentRelations := widget.NewLabel("")
	grid := container.NewGridWithColumns(36)
	logo := canvas.NewImageFromFile("images/logo gt.jpg")
	title := widget.NewLabel("Groupie Tracker")
	about := widget.NewLabel("Groupie Trackers consiste à recevoir une API donnée et à manipuler les données qu'elle contient, afin de créer un site affichant ces informations.")
	by := widget.NewLabel("Application créée par Alexandre, Kevin & Adriana")
	home := container.NewCenter(container.NewVBox(container.NewCenter(title), logo, widget.NewLabel("Application pour voir des informations sur vos artistes et groupes préférés !")))
	abouts := container.NewCenter(container.NewVBox(container.NewCenter(about), container.NewCenter(by)))

	//creates a clickable list
	listArtists := widget.NewList(
		func() int {
			if searchOn {
				return len(result)
			} else {
				return len(artists)
			}
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			if searchOn {
				object.(*widget.Label).SetText(temp[id].Name)
			} else {
				object.(*widget.Label).SetText(artists[id].Name)
			}
		})

	//Creation of the search bar
	searchEntry := fynex.NewCompletionEntry([]string{})

	//Call the function to displays input suggestions as you write
	searchbar.AutoCompletion(searchEntry)

	//Call the function to create a new button
	searchButton := searchbar.CreateButton(listArtists, &result, &temp, &searchOn,
		searchEntry, artists, v, obj)

	//Add a behaviour when you click on of the case of the list
	listArtists.OnSelected = func(id widget.ListItemID) {
		if !searchOn {

			list.ArtistsDisplayOff(id, r, img_grp, contentName, artist, artists, contentMembers, contentCreationDate, contentFirstAlbum,
				contentRelations, relations, locationsIndex, relationsIndex, s, img_world, grid)

		} else {

			list.ArtistsDisplayOn(id, r, img_grp, contentName, artist, artists, temp, contentMembers, contentCreationDate, contentFirstAlbum,
				contentRelations, relations, locationsIndex, s, img_world, grid)
		}
	}

	//A container to gisplay the geolocation
	geoloc := container.NewHBox(
		container.NewWithoutLayout(img_world, grid),
	)

	//Call a new window to display the geolocations of the concerts of this band
	btngeo := widget.NewButton("button", func() {
		w2.SetContent(geoloc)
		w2.Show()
	})

	//Add all the content in a container
	// world := container.NewWithoutLayout(img_world)

	artistDisplay := container.NewHSplit(
		listArtists,
		container.NewVBox(
			img_grp,
			contentName,
			contentMembers,
			contentCreationDate,
			contentFirstAlbum,
			contentRelations,
			btngeo,
		),
	)

	searchbutton := container.NewWithoutLayout(searchButton, obj)
	searchBar := container.NewVBox(searchEntry, searchbutton)

	display := container.NewBorder(searchBar, nil, nil, nil, artistDisplay)

	//tab bar to allow the user to navigate between several application views or pages
	appTabs := container.NewAppTabs(container.NewTabItem("Home", home),
		container.NewTabItem("Artists", display), container.NewTabItem("About", abouts))

	//Add style
	contentName.Wrapping = fyne.TextWrapWord
	contentMembers.Wrapping = fyne.TextWrapWord
	contentCreationDate.Wrapping = fyne.TextWrapWord
	contentFirstAlbum.Wrapping = fyne.TextWrapWord
	contentRelations.Wrapping = fyne.TextWrapWord
	logo.FillMode = canvas.ImageFillOriginal
	artistDisplay.Offset = 0.2
	searchButton.Resize(fyne.NewSize(300, 35))
	searchButton.Move(fyne.NewPos(0, 0))
	obj.Resize(fyne.NewSize(300, 35))
	obj.Move(fyne.NewPos(0, 0))
	searchEntry.Resize(fyne.NewSize(300, 35))
	grid.Resize(fyne.NewSize(900, 562))
	grid.Move(fyne.NewPos(0, 0))
	img_world.Resize(fyne.NewSize(900, 562))
	img_world.Move(fyne.NewPos(0, 0))
	contentRelations.Resize(fyne.NewSize(300, 400))
	contentRelations.Move(fyne.NewPos(800, 0))
	w.SetContent(appTabs)
	w.ShowAndRun()
}
