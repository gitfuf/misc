package theone

import (
	"sync"
)

type TheOne struct {
	info string
}

var singleton TheOne
var once sync.Once

func New(info string) TheOne {
	once.Do(func() {
		singleton = TheOne{info: info}
	})
	return singleton
}
