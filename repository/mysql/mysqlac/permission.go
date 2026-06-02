package mysqlac

import "github.com/SoroushBeigi/knowledge-game/entity"

func scanPermission(scanner scanner) (entity.Permission, error) {
	var createdAt []uint8
	var perm entity.Permission

	err := scanner.Scan(&perm.ID, &perm.Title, &createdAt)

	return perm, err
}

type scanner interface {
	Scan(dest ...any) error
}
