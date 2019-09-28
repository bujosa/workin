package models

import (
	"Backend/Entities"
	"database/sql"
)

type UnionModel struct {
	Db *sql.DB
}

var unions []entities.Union

// Para agregar Usuarios

// Esto es para listar los usuarios
func (unionmodel UnionModel) FindAll() ([]entities.Union, error) {
	rows, err := unionmodel.Db.Query("select D.Fecha, D.Hora, D.Descripcion, D.Modo, L.Longitud, L.Latitud, U.IdUsuario from Delitos as D INNER JOIN Locaciones as L  on D.Id_Locacion = L.IdLocacion INNER JOIN Usuarios as U on L.IdUsuario = U.IdUsuario")

	if err != nil {
		return nil, err
	} else {
		var unions []entities.Union
		for rows.Next() {
			var date           string 
			var	time 		   string
			var description    string
			var mode	       string
			var longitude      float64
			var latitude	   float64
			var iduser		   int64

			err2 := rows.Scan(&date, &time, &description, &mode, &longitude, &latitude, &iduser)
			if err2 != nil {
				return nil, err2
			} else {
				union := entities.Union{
				    Date:           date,
					Time: 			time,
					Description:    description,
					Mode:       	mode,
					Longitude:      longitude,
					Latitude:       latitude,
					Iduser:         iduser,
				}

				unions = append(unions, union)
			}
		}

		return unions, nil
	}

}
