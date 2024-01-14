package backend

type Band struct { // collect all data in one struct
	Artist       Artist
	Locations    Locations
	Dates        Dates
	Relations    []Relations
	RelationsMap RelationsMap
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationsMap struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Locations string
	Dates     []string
}

type ErrorPage struct {
	ErrStatus string
	ErrMsg    string
}

type Coordinates struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
