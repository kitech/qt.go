package toolutil

import (
	"os/exec"

	"github.com/kitech/qt.go/miscutil"
)

// RunCmdOut runs a comand and returns the commands output, or an error
func RunCmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	occ, err := cmd.Output() // Output runs the command and returns its standard output.
	miscutil.ErrPrint(err, name, arg, occ)
	if err != nil {
		return "", err
	}
	cmd.Wait()
	return string(occ), nil
}
