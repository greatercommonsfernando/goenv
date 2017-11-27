package lib

import (
	"runtime"
	"fmt"
	"os"
	"io/ioutil"
)

func Usage() {
	fmt.Println("Usage: goenv [destination_folder]")
}

func GetActivateFilename() string {
	if runtime.GOOS == "windows" {
		return activateCmdFileName
	} else {
		return activateBashFileName
	}
}

func GetActivateContents() string {
	if runtime.GOOS == "windows" {
		return activateCmdContents
	} else {
		return activateBashContents
	}
}

func EnsurePathExists(path string) error {
	return os.MkdirAll(path, 0755)
}

func CheckIfExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	} else {
		return true
	}
}

func CreateGoEnv(path string) error {
	if CheckIfExists(path) {
		fmt.Println("Refreshing existing goenv in", path)
	} else {
		fmt.Println("Creating new goenv in", path)
	}
	if err := EnsurePathExists(path); err != nil {
		return err
	}
	binPath := fmt.Sprintf("%s/bin", path)
	if err := EnsurePathExists(binPath); err != nil {
		return err
	}
	WriteActivateScript(binPath)
	if err := EnsurePathExists(fmt.Sprintf("%s/pkg", path)); err != nil {
		return err
	}
	if err := EnsurePathExists(fmt.Sprintf("%s/src", path)); err != nil {
		return err
	}
	return nil
}

func WriteActivateScript(binPath string) error {
	return ioutil.WriteFile(fmt.Sprintf("%s/%s", binPath, GetActivateFilename()),
		[]byte(GetActivateContents()), 0755)
}
