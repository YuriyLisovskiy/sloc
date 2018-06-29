package cli

import (
	"fmt"
	"errors"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func Validate(dirs, files []string) error {
	err := ValidateParams(dirs, files)
	err = validatePath(dirs, files)
	return err
}

func ValidateParams(dirs, files []string) error {
	if len(dirs) == 0 && len(files) == 0 {
		return errors.New(fmt.Sprintf(argsError, "-d", "-f"))
	} else {
		if len(dirs) > 0 && len(files) > 0 {
			return errors.New(fmt.Sprintf(argsError, "-d", "-f only"))
		}
	}
	return nil
}

func validatePath(dirs, files []string) error {
	for _, dir := range dirs {
		if len(dir) > 0 && !utils.IsDirectory(dir) {
			return errors.New(fmt.Sprintf(err + " is not a directory", dir))
		}
	}
	for _, file := range files {
		if len(file) > 0 && !utils.IsFile(file) {
			return errors.New(fmt.Sprintf(err + " is not a file", file))
		}
	}
	return nil
}
