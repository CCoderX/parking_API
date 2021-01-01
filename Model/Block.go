package parking_API

import "database/sql"

type block struct {
	BlockId         int    `json:"block_id"`
	ZoneId          int    `json:"zone_id"`
	Name            string `json:"name"`
	Capacity        int    `json:"capacity"`
	AvailableSpaces int    `json:"available_spaces"`
}
func (b *block) getBlock(db *sql.DB) error {
	return db.QueryRow("SELECT name, block_id,Zone_id,capacity,available_spaces FROM blocks WHERE block_id=$1",
		b.BlockId).Scan(&b.Name, &b.BlockId,&b.ZoneId)
}

func (b *block) getBlocks(db *sql.DB) ([]block, error) {
	rows, err := db.Query("SELECT block_id, zone_id, name , capacity,available_spaces FROM blocks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blocks []block

	for rows.Next() {
		var b block
		if err := rows.Scan(&b.BlockId, &b.ZoneId, &b.Name ); err != nil {
			return nil, err
		}
		blocks = append(blocks, b)
	}

	return blocks, nil
}
func (b *block) updateBlock(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE blocks SET  name=$1 ,capacity=$2  WHERE block_id =$3",
			b.Name, b.Capacity ,b.BlockId)
	return err
}
func (b *block) deleteBlock(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM blocks WHERE block_id=$1", b.BlockId)

	return err
}
//this function is used to delete all blocks related to a zone id
func (b *block) deleteBlockByZoneID(db *sql.DB,zoneID int) error {
	_, err := db.Exec("DELETE FROM blocks WHERE zone_id=$1", zoneID)

	return err
}
func (b *block) createBlock(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO blocks(name, capacity, zone_id) VALUES($1, $2 ,$3) RETURNING id",
		b.Name,b.Capacity,b.ZoneId).Scan(&b.BlockId)

	if err != nil {
		return err
	}

	return nil
}
//this function is used when a car leave a block to increase available blocks count
func (b *block) freeSpace(db *sql.DB ) error {
	b.getBlock(db)
	b.AvailableSpaces++
	err := b.updateBlock(db)

	if err != nil  {
		return err
	}
	return nil
}
//this function is used when a car enter a block
func (b *block) allocateSpace(db *sql.DB) error {
	b.AvailableSpaces--
	err := b.updateBlock(db)
	if err != nil {
		return err
	}

	return nil
}

