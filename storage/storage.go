package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
  FileName string
}

func NewStorage[T any](fn string) *Storage[T] {
  return &Storage[T]{
    FileName: fn,
  }
}

func (s *Storage[T]) Save(data T) error {
   d, err := json.MarshalIndent(data, "", "\t")

   if err != nil {
     fmt.Println(err)
   }

   return os.WriteFile(s.FileName, d, 777)
}

func (s *Storage[T]) Load(data *T) error {
  d, err := os.ReadFile(s.FileName)

  if err != nil {
    fmt.Println(err)
  }

  return json.Unmarshal(d, data)
}
