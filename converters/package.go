package converters

import (
	"errors"
	"os/exec"
)

// TODO: this package is gross and overloaded.
// maybe split:
//   - external processes (md -> html, image resize etc)
//   - markdown to post (where does this go?)

var requiredExecutables = []string{"pandoc", "convert"}

// FindExecutables confirms that the required external programs are installed,
// returning an error if they do not exist in the PATH.
func FindExecutables() error {
	for _, exe := range requiredExecutables {
		cmd := exec.Command("sh", "-c", "which "+exe)
		if err := cmd.Run(); err != nil {
			return errors.New("missing: " + exe)
		}
	}

	return nil
}
