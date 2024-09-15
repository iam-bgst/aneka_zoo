package constants

import "gorm.io/gorm"

type (
	KeyTrx string
	Common interface {
		GetDB() *gorm.DB
	}
)

const (
	KeyTransaction KeyTrx = KeyTrx("trx-key")
)
