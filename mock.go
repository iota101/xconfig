package xconfig

type mapConfig struct {
	data map[K]any
}

func FromMap(data map[K]any) Config {
	return &mapConfig{data: data}
}

func (c *mapConfig) Get(key K) Value {
	val, found := c.data[key]
	if !found {
		return emptyValue(string(key))
	}
	return newValue(val, true, string(key))
}

func (c *mapConfig) Has(key K) bool {
	_, found := c.data[key]
	return found
}

type envMapSecret struct {
	data map[E]string
}

func FromEnvMap(data map[E]string) Secret {
	return &envMapSecret{data: data}
}

func (s *envMapSecret) Get(key E) Value {
	val, found := s.data[key]
	if !found {
		return emptyValue(string(key))
	}
	return newValue(val, true, string(key))
}

func (s *envMapSecret) Has(key E) bool {
	_, found := s.data[key]
	return found
}
