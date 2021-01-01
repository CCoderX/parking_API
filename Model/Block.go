package Model

import "database/sql"

type Block struct {
	BlockId         int    `json:"block_id"`
	ZoneId          int    `json:"Zone_id"`
	Name            string `json:"name"`
	Capacity        int    `json:"capacity"`
	AvailableSpaces int    `json:"available_spaces"`
}
func (b *Block) GetBlock(db *sql.DB) error {
	return db.QueryRow("SELECT name, block_id,Zone_id,capacity,available_spaces FROM blocks WHERE block_id=$1",
		b.BlockId).Scan(&b.Name, &b.BlockId,&b.ZoneId)
}

func (b *Block) GetBlocks(db *sql.DB) ([]Block, error) {
	rows, err := db.Query("SELECT block_id, Zone_id, name , capacity,available_spaces FROM blocks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blocks []Block

	for rows.Next() {
		var b Block
		if err := rows.Scan(&b.BlockId, &b.ZoneId, &b.Name ); err != nil {
			return nil, err
		}
		blocks = append(blocks, b)
	}

	return blocks, nil
}
func (b *Block) UpdateBlock(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE blocks SET  name=$1 ,capacity=$2  WHERE block_id =$3",
			b.Name, b.Capacity ,b.BlockId)
	return err
}
func (b *Block) DeleteBlock(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM blocks WHERE block_id=$1", b.BlockId)

	return err
}

//this function is used to delete all blocks related to a Zone id
func (b *Block) DeleteBlockByZoneID(db *sql.DB,zoneID int) error {
	_, err := db.Exec("DELETE FROM blocks WHERE Zone_id=$1", zoneID)

	return err
}
func (b *Block) CreateBlock(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO blocks(name, capacity, Zone_id) VALUES($1, $2 ,$3) RETURNING id",
		b.Name,b.Capacity,b.ZoneId).Scan(&b.BlockId)

	if err != nil {
		return err
	}

	return nil
}

//this function is used when a Car leave a Block to increase available blocks count
func (b *Block) FreeSpace(db *sql.DB ) error {
	b.GetBlock(db)
	b.AvailableSpaces++
	err := b.UpdateBlock(db)

	if err != nil  {
		return err
	}
	return nil
}

//this function is used when a Car enter a Block
func (b *Block) AllocateSpace(db *sql.DB) error {
	b.AvailableSpaces--
	err := b.UpdateBlock(db)
	if err != nil {
		return err
	}

	return nil
}

