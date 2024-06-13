package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   id,
		Name: filename,
		Data: blob,
	}, nil
}
