package storage

import (
	"io"
	"os"
	"path/filepath"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(directoryPath string, fileName string) (*FileStorage, error) {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err = os.Mkdir(directoryPath, 0700)
		if err != nil {
			return nil, err
		}
	}

	_, err := os.OpenFile(filepath.Join(directoryPath, fileName), os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{filePath: filepath.Join(directoryPath, fileName)}, nil
}

func (fileStorage *FileStorage) Read() ([]byte, error) {
	file, err := os.OpenFile(fileStorage.filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fileStorage *FileStorage) Write(data []byte) error {
	file, err := os.OpenFile(fileStorage.filePath, os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
