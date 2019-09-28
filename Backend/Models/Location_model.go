package models

import (
	"Backend/Entities"
	"database/sql"
)

type LocationModel struct {
	Db *sql.DB
}

var locations []entities.Location

func (locationmodel LocationModel) FindAll() ([]entities.Location, error) {
	rows, err := locationmodel.Db.Query("select * from Locaciones")

	if err != nil {
		return nil, err
	} else {
		
		for rows.Next() {
			var	id             int64 
			var user 		   int64
			var longitude	   float64
			var latitude	   float64

			err2 := rows.Scan(&id, &user, &longitude, &latitude)
			if err2 != nil {
				return nil, err2
			} else {
				location := entities.Location{
					Id:             id,
					User:			user,
					Longitude:		longitude,
					Latitude:		latitude, 
				}

				locations = append(locations, location)
			}
		}

		return locations, nil
	}

}

func (locationmodel LocationModel) CreateLocation(LocationN entities.Location) (int64, error) {


	result, err := locationmodel.Db.Exec("INSERT Locaciones (IdUsuario, Longitud, Latitud) values (?, ?, ?)", 
	LocationN.User, LocationN.Longitude, LocationN.Latitude)

	if err != nil {
		return 0, err
	} else {
		LocationN.Id,_ = result.LastInsertId() 
		locations = append(locations, LocationN)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
	
}
