package cli

import (
	"fmt"
	"errors"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func Validate(dirs, file, multiple string) error {
	err := ValidateParams(dirs, file, multiple)
	err = validatePaths(dirs, file)
	return err
}

func ValidateParams(dir, file, multiple string) error {
	if len(dir) == 0 && len(file) == 0 && len(multiple) == 0 {
		return errors.New(fmt.Sprintf(argsError, "-d, -f", "-m"))
	} else {
		if len(dir) > 0 && len(file) > 0 {
			return errors.New(fmt.Sprintf(argsError, "-d", "-f only"))
		}
		if len(dir) > 0 && len(multiple) > 0 {
			return errors.New(fmt.Sprintf(argsError, "-d", "-m only"))
		}
		if len(file) > 0 && len(multiple) > 0 {
			return errors.New(fmt.Sprintf(argsError, "-f", "-m only"))
		}
		if len(dir) > 0 && len(file) > 0 && len(multiple) > 0 {
			return errors.New(fmt.Sprintf(argsError, "-d, -f", "-m only"))
		}
	}
	return nil
}

func validatePaths(dir, file string) error {
	if len(dir) > 0 && !utils.IsDirectory(dir) {
		return errors.New(fmt.Sprintf(err+" is not a directory", dir))
	}
	if len(file) > 0 && !utils.IsFile(file) {
		return errors.New(fmt.Sprintf(err+" is not a file", file))
	}
	return nil
}
