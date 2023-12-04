package dataload

import (
	"Groupie/data"
	"encoding/json"
	"log"
	"net/http"
)

//All functions present are used to read and return data from a json file through a link

func DataArtists(url string) []data.Artists {
	var artists []data.Artists
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

func DataLocations(url string) data.LocationsIndex {
	var locations data.LocationsIndex
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

func DataLocation(url string) data.Locations {
	var locations data.Locations
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url (simple): ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&locations)
	if err != nil {
		log.Fatal("Error during Decode (simple): ", err)
	}
	return locations
}

func DataDates(url string) data.DatesIndex {
	var dates data.DatesIndex
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

func DataRelations(url string) data.RelationsIndex {
	var relations data.RelationsIndex
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

func DataRelation(url string) data.Relations {
	var relations data.Relations
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url (simple): ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&relations)
	if err != nil {
		log.Fatal("Error during Decode (simple): ", err)
	}
	return relations
}

func GetGeoUrl(url string) []data.Geos {
	var GeoUrl []data.Geos
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
