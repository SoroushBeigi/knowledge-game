package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/entity"
)

func (db MySQLDB) IsPhoneNumberUnique(pn string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8
	row := db.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, pn)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		log.Printf("DB Error IsPhoneNumberUnique: %v\n", err)
		return false, fmt.Errorf("Error reading from Database")
	}

	return false, nil

}

func (db MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := db.db.Exec(`INSERT INTO users(name,phone_number) VALUES(?, ? )`, u.Name, u.PhoneNumber)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		return entity.User{}, fmt.Errorf("Database Error!")
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}
