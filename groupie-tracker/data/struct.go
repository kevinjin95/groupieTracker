package data

// Structure that contains the variables to use the various data of the json but also variables outside json

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
	Album        string
	Date         string
	Member       string
}

type LocationsIndex struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id        int      `json:"id"`
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

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
	DatesLocation  string
}

type GeoUrl struct {
	lat string
	lon string
}

type Geos struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	Geo string
}
