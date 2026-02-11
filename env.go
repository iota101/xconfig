package xconfig

import "os"

type envSecret struct{}

func FromEnv() Secret {
	return &envSecret{}
}

func (s *envSecret) Get(key E) Value {
	val, found := os.LookupEnv(string(key))
	if !found {
		return emptyValue(string(key))
	}
	return newValue(val, true, string(key))
}

func (s *envSecret) Has(key E) bool {
	_, found := os.LookupEnv(string(key))
	return found
}
