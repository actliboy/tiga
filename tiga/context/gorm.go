package contexti

import (
	gormi "github.com/liov/tiga/utils/dao/db/gorm"
	"gorm.io/gorm"
)

func (c *Ctx) NewDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{Context: gormi.SetTranceId(c.TraceID), NewDB: true})
}
