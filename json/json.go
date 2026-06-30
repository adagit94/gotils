package json

import (
	"github.com/adagit94/gotils/fs"
	"github.com/bytedance/sonic"
	"os"
)

func JsonBytes[V any](v *V) ([]byte, error) {
	return sonic.Marshal(v)
}

func JsonStr[V any](v *V) (string, error) {
	return sonic.MarshalString(v)
}

func JsonToFile[V any, D []byte | string](path string, perms os.FileMode, v *V, marshal func(v *V) (D, error), write func(path string, perms os.FileMode, data D) error) error {
	d, err := marshal(v)

	if err != nil {
		return err
	}

	if err := write(path, perms, d); err != nil {
		return err
	}

	return nil
}

func JsonBytesToFile[V any](path string, perms os.FileMode, v *V) error {
	return JsonToFile(path, perms, v, JsonBytes, fs.WriteFileBytes)
}

func JsonStrToFile[V any](path string, perms os.FileMode, v *V) error {
	return JsonToFile(path, perms, v, JsonStr, fs.WriteFileString)
}

func JsonParseBytes[T any](source []byte, target *T) error {
	return sonic.Unmarshal(source, target)
}

func JsonParseStr[T any](source string, target *T) error {
	return sonic.UnmarshalString(source, target)
}

func JsonParseBytesFromFile[T any](path string, target *T) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return JsonParseBytes(data, target)
}
