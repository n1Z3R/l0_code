package cache

import (
	"l0/internal/models"
	"log"
	"sync"
)

type Cache struct {
	sync.RWMutex
	Value map[string]models.Model
}

func NewCache() *Cache {
	return &Cache{sync.RWMutex{}, make(map[string]models.Model)}
}

func (c *Cache) Add(id string, value models.Model) {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.Value[id] = value
	log.Println("cached")
}

func (c *Cache) Get(id string) (models.Model, bool) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	if data, ok := c.Value[id]; ok {
		log.Println("loaded from cache!")
		return data, ok
	}
	return models.Model{}, false

}
