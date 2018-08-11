package driver

import (
	"fmt"
	"net"
	"time"

	myssh "github.com/WIZARDISHUNGRY/docker-machine-driver-openbsd/ssh"
	"github.com/WIZARDISHUNGRY/docker-machine-driver-openbsd/vmm"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/docker/machine/libmachine/state"
)

// Driver parameters
type Driver struct {
	*drivers.BaseDriver
	*vmm.Instance
	SSHKeyPair *ssh.KeyPair
}

const username string = "docker"
const password string = "tcuser"

// NewDriver creates and returns a new instance of the driver
func NewDriver() *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{},
		Instance: &vmm.Instance{
			Label: "docker",                    // FIXME
			ISO:   "/home/jon/boot2docker.iso", // FIXME
			Disk:  "/home/jon/disk.img",        // FIXME – vmctl create disk.img -s 1M
			Mem:   1024,
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "openbsd"
}

// Create copy ssh key in docker-machine dir and set the node IP
func (d *Driver) Create() (err error) {
	err = d.Instance.Start()
	d.BaseDriver.IPAddress, err = d.Instance.GetIP()
	d.BaseDriver.SSHUser = username

	// copy SSH key pair to machine directory
	if err := d.SSHKeyPair.WriteToFile(d.GetSSHKeyPath(), d.GetSSHKeyPath()+".pub"); err != nil {
		return fmt.Errorf("Error when copying SSH key pair to machine directory: %s", err.Error())
	}

	fmt.Println("Waiting for ssh")

	// Wait for SSH
	for {
		//err := drivers.WaitForSSH(d)
		// FIXME should test if instance is still alive here as well using vmctl
		client, err := myssh.Dial(d.IPAddress, "22", username, password)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			if client != nil {
				return myssh.CopyID(client, d.SSHKeyPair.PublicKey)
			}
			fmt.Print(".")
			break
		}

	}

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

	return state.Running, nil // FIXME

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
	return fmt.Errorf("You can't kill a vmm instance")
}

// Start don't do anything
func (d *Driver) Start() error {
	return fmt.Errorf("You can't start a vmm instance")
}

// Stop don't do anything
func (d *Driver) Stop() error {
	return fmt.Errorf("You can't stop a vmm instance")
}

// Restart don't do anything
func (d *Driver) Restart() error {
	return fmt.Errorf("You can't restart a vmm instance")
}

// Remove delete the resources reservation
func (d *Driver) Remove() error {
	// FIXME not implemented

	return nil
}

// SetConfigFromFlags configure the driver from the command line arguments
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {

	// Unimplemented FIXME

	return nil

}

// PreCreateCheck check parameters
func (d *Driver) PreCreateCheck() (err error) {

	// check if a SSH key pair is available
	if d.SSHKeyPair == nil {
		// generate a new SSH key pair
		d.SSHKeyPair, err = ssh.NewKeyPair()
		if err != nil {
			return fmt.Errorf("Error when generating a new SSH key pair: %s", err.Error())
		}
	}

	return
}
