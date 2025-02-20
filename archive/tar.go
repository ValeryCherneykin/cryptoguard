package archive

import (
	"archive/tar"
	"cryptoguard/files"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateTar(sourceDir, outputTar string) {
	exists, isDir := files.PathExists(sourceDir)

	if !exists || !isDir {
		return
	}

	tarFile, err := os.Create(outputTar)
	if err != nil {
		fmt.Println("Ошибка создании tar-файла", err)
		return
	}
	tw := tar.NewWriter(tarFile)
	defer tw.Close()
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(sourceDir, filePath)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		header := &tar.Header{
			Name: relPath,
			Mode: int64(info.Mode()),
			Size: info.Size(),
		}
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		_, err = io.Copy(tw, file)
		return err
	})

	if err != nil {
		fmt.Println("Ошибка при создании архива:", err)
	} else {
		fmt.Println("Архив создан:", outputTar)
	}
}
