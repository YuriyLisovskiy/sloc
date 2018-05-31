package args

import (
	"os"
	"fmt"
	"errors"
)

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return true
	}
	return info.Mode().IsRegular()
}

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
	if len(dir) > 0 && !isDir(dir) {
		return errors.New(fmt.Sprintf(err + " is not a directory", dir))
	}
	if len(file) > 0 && !isFile(file) {
		return errors.New(fmt.Sprintf(err + " is not a file", file))
	}
	return nil
}
