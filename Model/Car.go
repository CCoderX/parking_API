package Model

import "database/sql"

type car struct {
	BlockId  int    `json:"block_id"`
	CarId    int    `json:"car_id"`
	PlateStr string `json:"plate_str"`
	ParkDate string `json:"park_date"`
}
func (c *car) getCar(db *sql.DB) error {
	return db.QueryRow("SELECT block_id, plate_str,park_date FROM cars WHERE car_id=$1",
		c.CarId).Scan(&c.BlockId, &c.PlateStr,&c.ParkDate)
}
func (c *car) getCars(db *sql.DB) ([]car, error) {
	rows, err := db.Query("SELECT block_id, car_id , plate_str ,park_date FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []car
	for rows.Next() {
		var c car
		if err := rows.Scan(&c.BlockId, &c.CarId, &c.PlateStr,&c.ParkDate); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
func (c *car) getCarsByBlockID(db *sql.DB,blockID int) ([]car, error) {
	rows, err := db.Query("SELECT block_id, car_id , plate_str ,park_date FROM cars WHERE block_id=$1",blockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []car
	for rows.Next() {
		var c car
		if err := rows.Scan(&c.BlockId, &c.CarId, &c.PlateStr,&c.ParkDate); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
func (c *car) createCar(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO cars(block_id, plate_str, park_date) VALUES($1, $2 ,$3) RETURNING id",
		c.BlockId,c.ParkDate,c.PlateStr).Scan(&c.CarId)

	if err != nil {
		return err
	}

	return nil
}
func (c *car) deleteCar(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM cars WHERE car_id=$1", c.CarId)
	b := block{c.BlockId,0,"",0,0}
	if err1 := b.getBlock(db) ;err1 != nil {
		return err1
	}
	if err2 := b.freeSpace(db) ;err2 != nil {
		return err2
	}
	return err
}
func (c *car) updateCar(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE cars SET Block_id=$1, Plate_str=$2 ,Park_date=$3 WHERE Car_id =$4",
			c.BlockId, c.PlateStr, c.ParkDate,c.CarId)

	return err
}