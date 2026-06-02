package mysqlac

import (
	"strings"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"github.com/SoroushBeigi/knowledge-game/pkg/slice"
)

func (d *db) GetUserPermissions(userID uint, role entity.Role) ([]string, error) {
	const op = "mysql.GetUserACL"
	db := d.conn.DB()

	roleACL := make([]entity.AccessControl, 0)

	rRows, err := db.Query(`SELECT * FROM access_controls WHERE actor_type = ? AND actor_id = ?`, entity.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	defer rRows.Close()

	for rRows.Next() {
		ac, err := scanAccessControl(rRows)

		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithCode(richerror.UnexpectedCode).
				WithMessage(errmessage.ErrorMsgUnexpected)
		}

		roleACL = append(roleACL, ac)

	}

	if err := rRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	userACL := make([]entity.AccessControl, 0)

	uRows, err := db.Query(`SELECT * FROM access_controls WHERE actor_type = ? AND actor_id = ?`, entity.UserActorType, userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	defer uRows.Close()

	for uRows.Next() {
		ac, err := scanAccessControl(uRows)

		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithCode(richerror.UnexpectedCode).
				WithMessage(errmessage.ErrorMsgUnexpected)
		}

		userACL = append(userACL, ac)

	}

	if err := uRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	if len(permissionIDs) < 1 {

		return nil, nil
	}

	args := make([]any, len(permissionIDs))

	for i, id := range permissionIDs {
		args[i] = id
	}

	query := "SELECT * FROM permissions where id in (?" + strings.Repeat(",?", len(permissionIDs)-1) + ")"
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	defer rows.Close()

	permissionTitles := make([]string, len(permissionIDs))

	for rows.Next() {
		perm, err := scanPermission(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithCode(richerror.UnexpectedCode).
				WithMessage(errmessage.ErrorMsgUnexpected)
		}
		permissionTitles = append(permissionTitles, perm.Title)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithCode(richerror.UnexpectedCode).
			WithMessage(errmessage.ErrorMsgUnexpected)
	}

	return permissionTitles, nil

}

func scanAccessControl(scanner scanner) (entity.AccessControl, error) {
	var createdAt []uint8
	var ac entity.AccessControl

	err := scanner.Scan(&ac.ID, &ac.ActorID, &ac.ActorType, &ac.PermissionID, &createdAt)

	return ac, err
}
