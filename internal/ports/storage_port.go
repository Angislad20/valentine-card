package ports

type StoragePort interface {
	SaveFile(filePath string, data []byte) (string, error)
	DeleteFile(filePath string) error
	GetFile(filePath string) ([]byte, error)
}
