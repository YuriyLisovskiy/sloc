package args

import (
	"fmt"
	"errors"
)

var argsError = "sloc args error: use%s -%s or -%s"

func validate(dir, file string, files []string) error {
	if len(dir) == 0 && len(file) == 0 && len(files) == 0 {
		return errors.New(fmt.Sprintf(argsError + " or -%s", "", dirFlag, fileFlag, filesFlag))
	} else {
		if len(dir) > 0 && len(file) > 0 {
			return errors.New(fmt.Sprintf(argsError, " only", dirFlag, fileFlag))
		}
		if len(dir) > 0 && len(files) > 0 {
			return errors.New(fmt.Sprintf(argsError, " only", dirFlag, filesFlag))
		}
		if len(file) > 0 && len(files) > 0 {
			return errors.New(fmt.Sprintf(argsError, " only", fileFlag, filesFlag))
		}
		if len(dir) > 0 && len(file) > 0 && len(files) > 0 {
			return errors.New(fmt.Sprintf(argsError + " or -%s", " only", dirFlag, fileFlag, filesFlag))
		}
	}
	return nil
}
