package Model

import "database/sql"

type Zone struct {
	 Name     string `json:"name"`
	 Zone_id  int    `json:"Zone_id"`
	 Location string `json:"location"`
}
func (z *Zone) GetZone(db *sql.DB) error {
	return db.QueryRow("SELECT name, locatio FROM zones WHERE Zone_id=$1",
		z.Zone_id).Scan(&z.Name,&z.Location)
}

func (z *Zone) GetZones(db *sql.DB) ([]Zone, error) {
	rows, err := db.Query("SELECT Zone_id, name,  location  FROM zones")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var zones []Zone

	for rows.Next() {
		var z Zone
		if err := rows.Scan(&z.Zone_id, &z.Name, &z.Location); err != nil {
			return nil, err
		}
		zones = append(zones, z)
	}

	return zones, nil
}

func (z *Zone) UpdateZone(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE zones SET name=$1, location=$2 WHERE Zone_id =$3",
			z.Name, z.Location, z.Zone_id)

	return err
}

func (z *Zone) DeleteZone(db *sql.DB) error {

	_, err := db.Exec("DELETE FROM zones WHERE Zone_id=$1", z.Zone_id)

	return err
}

func (z *Zone) CreateZone(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO zones(name, location) VALUES($1, $2 ,$3,$4) RETURNING id",
		z.Name, z.Location).Scan(&z.Zone_id)

	if err != nil {
		return err
	}

	return nil
}
