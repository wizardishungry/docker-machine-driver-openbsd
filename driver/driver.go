package driver

import (
	"fmt"
	"net"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
)

// Driver parameters
type Driver struct {
	*drivers.BaseDriver
}

// NewDriver creates and returns a new instance of the driver
func NewDriver() *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{},
	}
}

// Create copy ssh key in docker-machine dir and set the node IP
func (d *Driver) Create() (err error) {
	err = nil
	return
}

// GetCreateFlags add command line flags to configure the driver
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{}
}

// GetIP returns the ip
func (d *Driver) GetIP() (string, error) {
	return d.BaseDriver.GetIP()
}

// GetMachineName returns the machine name
func (d *Driver) GetMachineName() string {
	return d.BaseDriver.GetMachineName()
}

// GetSSHHostname returns the machine hostname
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// GetState returns the state of the node
func (d *Driver) GetState() (state.State, error) {

	status := "wtf"

	switch status {
	case "waiting":
		return state.Starting, nil
	case "launching":
		return state.Starting, nil
	case "running":
		return state.Running, nil
	case "hold":
		return state.Stopped, nil
	case "error":
		return state.Error, nil
	case "terminated":
		return state.Stopped, nil
	default:
		return state.None, nil
	}
}

// GetURL returns the URL of the docker daemon
func (d *Driver) GetURL() (string, error) {
	// get IP address
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}

	// format URL 'tcp://host:2376'
	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
}

// Kill don't do anything
func (d *Driver) Kill() error {
	return fmt.Errorf("FIXME unimplemented")
}

// Start don't do anything
func (d *Driver) Start() error {
	return fmt.Errorf("FIXME unimplemented")
}

// Stop don't do anything
func (d *Driver) Stop() error {
	return fmt.Errorf("FIXME unimplemented")
}

// Restart don't do anything
func (d *Driver) Restart() error {
	return fmt.Errorf("FIXME unimplemented")
}

// Remove delete the resources reservation
func (d *Driver) Remove() error {
	// FIXME not implemented

	return nil
}

// SetConfigFromFlags configure the driver from the command line arguments
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {

	return fmt.Errorf("Unimplemented FIXME")

}
