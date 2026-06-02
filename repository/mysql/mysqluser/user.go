package mysqluser

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

func (db *db) IsPhoneNumberUnique(pn string) (bool, error) {
	const op = "sql.IsPhoneNumberUnique"
	d := db.conn.DB()
	row := d.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, pn)
	_, err := scanUser(row)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		log.Printf("DB Error IsPhoneNumberUnique: %v\n", err)
		return false, richerror.New(op).
			WithErr(err).
			WithMessage(errmessage.ErrorMsgUnexpected).
			WithCode(richerror.UnexpectedCode)
	}

	return false, nil

}

func (db *db) Register(u entity.User) (entity.User, error) {
	d := db.conn.DB()
	res, err := d.Exec(`INSERT INTO users(name,phone_number,password) VALUES(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		return entity.User{}, fmt.Errorf("Database Error!")
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (db *db) GetUserByPhoneNumber(pn string) (entity.User, error) {
	const op = "sql.GetUserByPhoneNumber"
	d := db.conn.DB()

	row := d.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, pn)
	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {

			return entity.User{},
				richerror.New(op).
					WithErr(err).
					WithMessage(errmessage.ErrorMsgNotFound).
					WithCode(richerror.NotFoundCode)
		}

		return entity.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(errmessage.ErrorMsgUnexpected).
				WithCode(richerror.UnexpectedCode)
	}

	return user, nil
}

func (db *db) GetUserByID(id uint) (entity.User, error) {
	const op = "sql.GetUserByID"
	user := entity.User{}
	d := db.conn.DB()

	row := d.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{},
				richerror.New(op).
					WithErr(err).
					WithMessage(errmessage.ErrorMsgNotFound).
					WithCode(richerror.NotFoundCode)
		}

		return entity.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(errmessage.ErrorMsgUnexpected).
				WithCode(richerror.UnexpectedCode)
	}

	return user, nil
}

func scanUser(scanner scanner) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)

	user.Role = parseRole(roleStr)

	return user, err
}

type scanner interface {
	Scan(dest ...any) error
}

func parseRole(r string) entity.Role {
	switch r {
	case "user":
		return entity.UserRole
	case "admin":
		return entity.AdminRole
	default:
		return entity.UserRole 
	}
}
