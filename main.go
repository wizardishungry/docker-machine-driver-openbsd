package main

import (
	"github.com/WIZARDISHUNGRY/docker-machine-driver-openbsd/driver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(driver.NewDriver())
}
