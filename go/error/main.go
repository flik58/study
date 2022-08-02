
package main

import (
	"github.com/pkg/errors"
)

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := errors.Cause(err).(temporary)
	return ok && te.Temporary()
}
