package storage

import (
	"io"
	"os"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	_, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{filePath: filePath}, nil
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
