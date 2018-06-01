package args

import (
	"fmt"
	"errors"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func Validate(dir, file string, files []string) error {
	err := ValidateParams(dir, file, files)
	err = validatePath(dir, file)
	return err
}

func ValidateParams(dir, file string, files []string) error {
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

func validatePath(dir, file string) error {
	if len(dir) > 0 && !utils.IsDirectory(dir) {
		return errors.New(fmt.Sprintf(err + " is not a directory", dir))
	}
	if len(file) > 0 && !utils.IsFile(file) {
		return errors.New(fmt.Sprintf(err + " is not a file", file))
	}
	return nil
}
