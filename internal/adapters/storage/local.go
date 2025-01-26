package storage

import "os"

type LocalStorage struct{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (ls *LocalStorage) SaveFile(filePath string, data []byte) (string, error) {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (ls *LocalStorage) DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LocalStorage) GetFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
