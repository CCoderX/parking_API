package Model

import "database/sql"

type zone struct {
	 Name     string `json:"name"`
	 zone_id  int    `json:"zone_id"`
	 Location string `json:"location"`
}
func (z *zone) getZone(db *sql.DB) error {
	return db.QueryRow("SELECT name, locatio FROM zones WHERE zone_id=$1",
		z.zone_id).Scan(&z.Name,&z.Location)
}

func (z *zone)getZones(db *sql.DB) ([]zone, error) {
	rows, err := db.Query("SELECT zone_id, name,  location  FROM zones")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var zones []zone

	for rows.Next() {
		var z zone
		if err := rows.Scan(&z.zone_id, &z.Name, &z.Location); err != nil {
			return nil, err
		}
		zones = append(zones, z)
	}

	return zones, nil
}

func (z *zone) updateZone(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE zones SET name=$1, location=$2 WHERE zone_id =$3",
			z.Name, z.Location, z.zone_id)

	return err
}

func (z *zone) deleteZone(db *sql.DB) error {

	_, err := db.Exec("DELETE FROM zones WHERE zone_id=$1", z.zone_id)

	return err
}

func (z *zone) createZone(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO zones(name, location) VALUES($1, $2 ,$3,$4) RETURNING id",
		z.Name, z.Location).Scan(&z.zone_id)

	if err != nil {
		return err
	}

	return nil
}
