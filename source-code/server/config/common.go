package config

import "sync"

// Context describes the application configurations.
type Context struct {
	conf *sync.Map
}
