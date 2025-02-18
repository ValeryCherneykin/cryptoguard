package files

import (
	"os"
)

func CreatFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func WriteFile(fileNema string, data []byte) error {
	err := os.WriteFile(fileNema, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func PathExists(path string) (bool, bool) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false
		} else {
			return false, info.IsDir()
		}
	}
	return true, info.IsDir()
}
