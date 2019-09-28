package entities

import (
	"fmt"
)

type Crime struct {
	Id             int64 
	Cat		 	   int64
	Loc		       int64
	Date		   string
	Time		   string
	Description    string
	Mode		   string
}

type Crimes []Crime

func (crime Crime) ToString() string {
	return fmt.Sprintf("id: %d\n category: %d\n location: %d\n date: %s\n time: %s\n description: %s\nmode: %s",
		crime.Id, crime.Cat, crime.Loc, crime.Date, crime.Time, crime.Description, crime.Mode)
}
