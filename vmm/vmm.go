package vmm

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

const vmctl string = "/usr/sbin/vmctl"

var startRegexp *regexp.Regexp

func init() {
	startRegexp = regexp.MustCompile(`vmctl: started vm (?P<Id>\d{1,}) successfully, tty /dev/`)
}

// Instance represents a VMM instance
type Instance struct {
	ISO   string
	Label string
	Disk  string
	Mem   uint16 // Megabytes
	Name  string
	id    string
}

// Start instance
func (i *Instance) Start() error {
	fmt.Println(vmctl, "start", i.Name,
		"-t", "docker", // FIXME hardcoded template name
		"-r", "/home/vm/boot2docker.iso",
		// "-r", i.ISO,

		//"-d", i.Disk,
		// "-d", "qcow:/home/jon/test.img",

		// "-m", strconv.Itoa(int(i.Mem))+"M",
	)
	cmd := exec.Command(vmctl, "start", i.id,
		"-t", "docker", // FIXME hardcoded template name
		"-r", "/home/vm/boot2docker.iso",
		// "-r", i.ISO,

		//"-d", i.Disk,
		// "-d", "qcow:/home/jon/test.img",
		// "-m", strconv.Itoa(int(i.Mem))+"M",
	)
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// TODO wrap exit codes if err is an ExitErr
		return err
	}
	message := out.String()

	fmt.Println("messsage", message)

	// FIXME brittle
	match := startRegexp.FindStringSubmatch(message)

	if len(match) != 1 {
		return errors.New("Couldn't start VMM " + message)
	}
	i.id = match[1]

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
