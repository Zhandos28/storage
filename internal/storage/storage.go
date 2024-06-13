package storage

import (
	"fmt"
	"github.com/Zhandos28/storage/internal/file"
	"github.com/google/uuid"
	"sync"
)

type Storage struct {
	mt    sync.RWMutex
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	mt := sync.RWMutex{}
	return &Storage{
		mt:    mt,
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.mt.Lock()
	s.Files[newFile.ID] = newFile
	s.mt.Unlock()

	return newFile, err
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	s.mt.RLock()
	foundFile, ok := s.Files[fileID]
	s.mt.RUnlock()

	if !ok {
		return nil, fmt.Errorf("file with such id %v doesn't exist", fileID)
	}

	return foundFile, nil
}
