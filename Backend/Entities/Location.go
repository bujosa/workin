package entities

import (
	"fmt"
)

type Location struct {
	Id             int64 
	User 		   int64
	Longitude	   float64
	Latitude	   float64
}

type Locations []Location

func (location Location) ToString() string {
	return fmt.Sprintf("id: %d\n user: %d\n longitude: %0.6f\n latitude: %0.6f",
		location.Id, location.User, location.Longitude, location.Latitude)
}