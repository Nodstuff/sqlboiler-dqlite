package main

import (
	"github.com/volatiletech/sqlboiler/v4/drivers"

	"github.com/nodstuff/sqlboiler-dqlite/driver"
)

func main() {
	drivers.DriverMain(&driver.DQLiteDriver{})
}
