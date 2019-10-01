package entities

import (
	"fmt"
)

type Union struct {
	
	Date		   string
	Time		   string
	Description    string
	Mode		   string
	Longitude	   float64
	Latitude       float64
	Iduser		   int64
	Categoria	   string
	
	
}

type Unions []Union

func (union Union) ToString() string {
	return fmt.Sprintf("date: %s\n time: %s\n description: %s\n mode: %s\n longitude: %0.6f\n latitude: %0.06f\niduser: %d\ncategoria: %s",
		union.Date, union.Time, union.Description, union.Mode, union.Longitude, union.Latitude, union.Iduser,union.Categoria)
}