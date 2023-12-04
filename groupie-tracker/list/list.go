package list

import (
	"Groupie/data"
	dataload "Groupie/dataLoad"
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// transforms all data into string to be able to display them afterwards (in this case if the user input is not good)
func ArtistsDisplayOff(id widget.ListItemID, r fyne.Resource, img_grp *canvas.Image, contentName *widget.Label, artist data.Artists, artists []data.Artists, contentMembers *widget.Label, contentCreationDate *widget.Label, contentFirstAlbum *widget.Label, contentRelations *widget.Label, relations data.Relations, locationsIndex data.LocationsIndex, relationsIndex data.RelationsIndex, s fyne.Resource, img_world *canvas.Image, grid *fyne.Container) {

	// img_grp = canvas.NewImageFromResource(r)
	lien_grp := artists[id].Image
	r, _ = fyne.LoadResourceFromURLString(lien_grp)
	img_grp.Resource = r
	img_grp.FillMode = canvas.ImageFillOriginal
	img_grp.Refresh()

	artist.Name = "Nom du groupe : " + artists[id].Name
	contentName.SetText(artist.Name)

	artist.Member = "Membres : "
	for _, m := range artists[id].Members {
		artist.Member += m + "  "
		contentMembers.SetText(artist.Member)
	}

	artist.Date = "Date de création : "
	artist.Date += strconv.Itoa(artists[id].CreationDate)
	contentCreationDate.SetText(artist.Date)

	artist.Album = "Premier album : " + artists[id].FirstAlbum
	contentFirstAlbum.SetText(artist.Album)

	var c int
	var listLocations []string
	relations.DatesLocation = "Concerts : "
	for _, l := range locationsIndex.Index[id].Locations {
		for c = 0; c <= len(relationsIndex.Index[id].DatesLocations[l])-1; c++ {
			var isDouble bool = false
			for _, r := range listLocations {
				if r == l {
					isDouble = true
				}
			}
			if !isDouble {
				relations.DatesLocation += l + " : " + relationsIndex.Index[id].DatesLocations[l][c] + "  "
				listLocations = append(listLocations, l)
			} else {
				relations.DatesLocation += ", " + relationsIndex.Index[id].DatesLocations[l][c] + " "
			}
			contentRelations.SetText(relations.DatesLocation)
		}
	}

	var city string
	var country string
	var cityTable []string
	var countryTable []string
	for i := range locationsIndex.Index[id].Locations {
		before, after, _ := strings.Cut(locationsIndex.Index[id].Locations[i], "-")
		city = strings.Title(before)
		country = strings.Title(after)
		cityTable = append(cityTable, city)
		countryTable = append(countryTable, country)
		// fmt.Print(city)
		// fmt.Print(country)
	}

	var Geo []data.Geos
	var x float64
	var y float64
	var z float64 = 10.0
	for i := 0; i <= len(cityTable)-1; i++ {
		lien_geo := "https://nominatim.openstreetmap.org/?addressdetails=1&q=" + cityTable[i] + "+" + countryTable[i] + "&format=json&limit=1"
		Geo = dataload.GetGeoUrl(lien_geo)
		latfloat, _ := strconv.ParseFloat(Geo[0].Lat, 64)
		lonfloat, _ := strconv.ParseFloat(Geo[0].Lon, 64)
		lonfloat += 170
		latfloat += 80
		for y = 180; y > 0; y -= z {
			for x = 0; x < 360; x += z {
				if (lonfloat > x && lonfloat < x+z) && (latfloat > y && latfloat < y+z) {
					// fmt.Println(" ")
					// fmt.Println("lonfloat, latfloat, to grid :", lonfloat, latfloat)
					// fmt.Println("x, y, to grid :", x, y)
					RedColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					pin := canvas.NewRectangle(RedColor)
					grid.Add(pin)
				} else {
					Transparent := color.NRGBA{R: 255, G: 255, B: 255, A: 50}
					pin := canvas.NewRectangle(Transparent)
					grid.Add(pin)
					// fmt.Print("x | y", x, y)
				}
			}
		}
	}

	lien_world := "https://www.bluewaterbio.com/wp-content/uploads/bwb-world-map.png"
	s, _ = fyne.LoadResourceFromURLString(lien_world)
	img_world.Resource = s
	img_world.FillMode = canvas.ImageFillContain
	img_world.Refresh()
}

// transforms all data into string to be able to display them afterwards (in this case if the user input is good)
func ArtistsDisplayOn(id widget.ListItemID, r fyne.Resource, img_grp *canvas.Image, contentName *widget.Label, artist data.Artists, artists []data.Artists, temp []data.Artists, contentMembers *widget.Label, contentCreationDate *widget.Label, contentFirstAlbum *widget.Label, contentRelations *widget.Label, relations data.Relations, locationsIndex data.LocationsIndex, s fyne.Resource, img_world *canvas.Image, grid *fyne.Container) {

	// img_grp = canvas.NewImageFromResource(r)
	lien_grp := artists[id].Image
	r, _ = fyne.LoadResourceFromURLString(lien_grp)
	img_grp.Resource = r
	img_grp.FillMode = canvas.ImageFillOriginal
	img_grp.Refresh()

	artist.Name = "Nom du groupe : " + artists[id].Name
	contentName.SetText(artist.Name)

	artist.Member = "Membres : "
	for _, m := range temp[id].Members {
		artist.Member += m + ", "
		contentMembers.SetText(artist.Member)
	}
	artist.Date = "Date de création : "
	artist.Date += strconv.Itoa(temp[id].CreationDate)
	contentCreationDate.SetText(artist.Date)

	artist.Album = "Premier album : " + temp[id].FirstAlbum
	contentFirstAlbum.SetText(artist.Album)

	indexLoca := dataload.DataLocation(temp[id].Locations)
	indexRela := dataload.DataRelation(temp[id].Relations)
	var c int
	var listLocations []string
	relations.DatesLocation = "Concerts : "
	for _, l := range indexLoca.Locations {
		for c = 0; c <= len(indexRela.DatesLocations[l])-1; c++ {
			var isDouble bool = false
			for _, r := range listLocations {
				if r == l {
					isDouble = true
				}
			}
			if !isDouble {
				relations.DatesLocation += l + " : " + indexRela.DatesLocations[l][c] + "  "
				listLocations = append(listLocations, l)
			} else {
				relations.DatesLocation += ", " + indexRela.DatesLocations[l][c] + " "
			}
			contentRelations.SetText(relations.DatesLocation)
		}
	}

	var city string
	var country string
	var cityTable []string
	var countryTable []string
	for i := range locationsIndex.Index[id].Locations {
		before, after, _ := strings.Cut(locationsIndex.Index[id].Locations[i], "-")
		city = strings.Title(before)
		country = strings.Title(after)
		cityTable = append(cityTable, city)
		countryTable = append(countryTable, country)
		// fmt.Print(city)
		fmt.Print(country)
	}

	var Geo []data.Geos
	var x float64
	var y float64
	var z float64 = 10.0
	for i := 0; i <= len(cityTable)-1; i++ {
		lien_geo := "https://nominatim.openstreetmap.org/?addressdetails=1&q=" + cityTable[i] + "+" + countryTable[i] + "&format=json&limit=1"
		Geo = dataload.GetGeoUrl(lien_geo)
		latfloat, _ := strconv.ParseFloat(Geo[0].Lat, 64)
		lonfloat, _ := strconv.ParseFloat(Geo[0].Lon, 64)
		fmt.Println(latfloat)
		fmt.Println(lonfloat)
		lonfloat += 170
		latfloat += 80
		for y = 180; y > 0; y -= z {
			for x = 0; x < 360; x += z {
				if (lonfloat > x && lonfloat < x+z) && (latfloat > y && latfloat < y+z) {
					// fmt.Println(" ")
					// fmt.Println("lonfloat, latfloat, to grid :", lonfloat, latfloat)
					// fmt.Println("x, y, to grid :", x, y)
					RedColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					pin := canvas.NewRectangle(RedColor)
					grid.Add(pin)
				} else {
					Transparent := color.NRGBA{R: 255, G: 255, B: 255, A: 50}
					pin := canvas.NewRectangle(Transparent)
					grid.Add(pin)
					// fmt.Print("x | y", x, y)
				}
			}
		}
	}

	lien_world := "https://www.bluewaterbio.com/wp-content/uploads/bwb-world-map.png"
	s, _ = fyne.LoadResourceFromURLString(lien_world)
	img_world.Resource = s
	img_world.FillMode = canvas.ImageFillContain
	img_world.Refresh()
}
