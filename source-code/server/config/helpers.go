package config

import "reflect"

// Get returns the value associated with the key as a string.
func (c *Context) Get(key string) string {
	v, _ := c.conf.Load(key)
	return reflect.ValueOf(v).String()
}

// IsCache checks if cache is enabled.
// It returns true for enabled, false for disabled.
func (c *Context) IsCache() bool {
	return !(c.Get("cache.enable") == "1" || c.Get("cache.enable") == "false")
}
