package cache

type Cache struct {
	items map[string]interface{}
}

func New() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

func (c *Cache) Get(key string) interface{} {
	if value, ok := c.items[key]; ok {
		return value
	}
	return nil
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
