package storage

type IStorage interface {
	Read() ([]byte, error)
	Write([]byte) error
}
