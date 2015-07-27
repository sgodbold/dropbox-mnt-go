package fs

import ()

var Cache MetadataCache

type MetadataCache struct {
	data map[string]Metadata
}

func (c *MetadataCache) Get(path string) (data Metadata, err error) {
	_, ok := c.data[path]
	if ok == false {
		data, err := GetMetadata(path)
		if err != nil {
			return Metadata{}, err
		}
		c.data[path] = data
	}
	return c.data[path], err
}

func CacheInit() {
	Cache.data = make(map[string]Metadata)
}
