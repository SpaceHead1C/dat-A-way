package pg

import (
	"dataway/pkg/db/pg"
	"fmt"
)

var errCanNotGetUniqueID = fmt.Errorf("can not get unique ID")

func isTomNameDuplicateError(err error, name string) bool {
	return pg.IsDuplicateKeyError(err, tableTomColumnName, name)
}
