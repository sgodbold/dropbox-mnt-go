// This cache will be replaced with something much better.
// I wrote this mostly to get a MVP out there.
package fs

import ()

var Cache AllCache

type AllCache struct {
	Metadata MetadataCache
	// ... more to come in the future
}

type MetadataCache struct {
	Data map[string]Metadata
}

func (c *MetadataCache) Get(path string) (data Metadata, err error) {
	_, ok := c.Data[path]
	if ok == false {
		data, err := GetMetadata(path)
		if err != nil {
			return Metadata{}, err
		}
		c.Data[path] = data
	}
	return c.Data[path], err
}
