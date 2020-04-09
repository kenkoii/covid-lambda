package main

type Case struct {
	Country    Country `json:"country"`
	Confirmed  int     `json:"confirmed"`
	Deaths     int     `json:"deaths"`
	Active     int     `json:"active"`
	Recovered  int     `json:"recovered"`
	LastUpdate int     `json:"lastupdate"`
}

type Country struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
