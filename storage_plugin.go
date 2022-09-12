package go_whatsapp

import "errors"

type StoragePlugin interface {
	Store(key string, value interface{}) error
	Load(key string) (interface{}, error)
	Remove(key string) error
}

type (
	defaultStoragePlugin struct {
		data map[string]interface{}
	}
)

func DefaultStore() StoragePlugin {
	return &defaultStoragePlugin{
		data: make(map[string]interface{}),
	}
}

func (d *defaultStoragePlugin) Store(key string, value interface{}) error {
	d.data[key] = value
	return nil
}

func (d *defaultStoragePlugin) Load(key string) (interface{}, error) {
	data, ok := d.data[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return data, nil
}

func (d *defaultStoragePlugin) Remove(key string) error {
	_, ok := d.data[key]
	if !ok {
		return errors.New("not found")
	}

	delete(d.data, key)

	return nil
}
