package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/entity"
)

func (db *MySQLDB) IsPhoneNumberUnique(pn string) (bool, error) {

	row := db.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, pn)
	_, err := scanUser(row)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		log.Printf("DB Error IsPhoneNumberUnique: %v\n", err)
		return false, fmt.Errorf("Error reading from Database")
	}

	return false, nil

}

func (db *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := db.db.Exec(`INSERT INTO users(name,phone_number,password) VALUES(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		return entity.User{}, fmt.Errorf("Database Error!")
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (db *MySQLDB) GetUserByPhoneNumber(pn string) (entity.User, error) {
	row := db.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, pn)
	user, err := scanUser(row)

	if err != nil {
		log.Printf("DB Error IsPhoneNumberUnique: %v\n", err)
		return entity.User{}, fmt.Errorf("Error reading from Database")
	}

	return user, nil
}

func (db *MySQLDB) GetUserByID(id uint) (entity.User, error) {
	user := entity.User{}

	row := db.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("DB Error IsPhoneNumberUnique ErrNoRows: %v\n", err)
			return entity.User{}, fmt.Errorf("Record not found")
		}
		log.Printf("DB Error IsPhoneNumberUnique: %v\n", err)
		return entity.User{}, fmt.Errorf("Error reading from Database")
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	return user, err
}
