package vmm

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

const vmctl string = "vmctl"

var startRegexp *regexp.Regexp

func init() {
	startRegexp = regexp.MustCompile(`vmctl: started vm (?P<Id>\d{1,}) successfully`)
}

// Instance represents a VMM instance
type Instance struct {
	ISO   string
	Label string
	Disk  string
	Mem   uint16 // Megabytes

	id string
}

// Start instance
func (i *Instance) Start() error {
	cmd := exec.Command(vmctl, "start", i.Label, "-L", "-r", i.ISO, "-d", i.Disk, "-m", strconv.Itoa(int(i.Mem))+"M")
	var out bytes.Buffer
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		// TODO wrap exit codes if err is an ExitErr
		return err
	}
	message := out.String()

	i.id = startRegexp.FindStringSubmatch(message)[1]

	if i.id == "" {
		return errors.New("Couldn't start VMM " + message)
	}

	// FIXME shoot newlines on vmctl console to speed up boot2docker

	return nil
}

// GetIP returns the ip
func (i *Instance) GetIP() (string, error) {
	if i.id == "" {
		return "", errors.New("Couldn't find VMM id")
	}
	return fmt.Sprintf("100.64.%s.3", i.id), nil
}
