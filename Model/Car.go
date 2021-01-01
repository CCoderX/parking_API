package Model

import "database/sql"

type Car struct {
	BlockId  int    `json:"block_id"`
	CarId    int    `json:"car_id"`
	PlateStr string `json:"plate_str"`
	ParkDate string `json:"park_date"`
}
func (c *Car) GetCar(db *sql.DB) error {
	return db.QueryRow("SELECT block_id, plate_str,park_date FROM cars WHERE car_id=$1",
		c.CarId).Scan(&c.BlockId, &c.PlateStr,&c.ParkDate)
}
func (c *Car) GetCars(db *sql.DB) ([]Car, error) {
	rows, err := db.Query("SELECT block_id, car_id , plate_str ,park_date FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []Car
	for rows.Next() {
		var c Car
		if err := rows.Scan(&c.BlockId, &c.CarId, &c.PlateStr,&c.ParkDate); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
func (c *Car) GetCarsByBlockID(db *sql.DB,blockID int) ([]Car, error) {
	rows, err := db.Query("SELECT block_id, car_id , plate_str ,park_date FROM cars WHERE block_id=$1",blockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []Car
	for rows.Next() {
		var c Car
		if err := rows.Scan(&c.BlockId, &c.CarId, &c.PlateStr,&c.ParkDate); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
func (c *Car) CreateCar(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO cars(block_id, plate_str, park_date) VALUES($1, $2 ,$3) RETURNING id",
		c.BlockId,c.ParkDate,c.PlateStr).Scan(&c.CarId)

	if err != nil {
		return err
	}

	return nil
}
func (c *Car) DeleteCar(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM cars WHERE car_id=$1", c.CarId)
	b := Block{c.BlockId,0,"",0,0}
	if err1 := b.GetBlock(db) ;err1 != nil {
		return err1
	}
	if err2 := b.FreeSpace(db) ;err2 != nil {
		return err2
	}
	return err
}
func (c *Car) UpdateCar(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE cars SET Block_id=$1, Plate_str=$2 ,Park_date=$3 WHERE Car_id =$4",
			c.BlockId, c.PlateStr, c.ParkDate,c.CarId)

	return err
}