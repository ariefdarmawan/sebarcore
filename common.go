package sebarcore

import (
	"errors"
	"fmt"
	"time"
)

var (
	pkgname = "sebarcore"
)

func ThrowErr(pkg, class, msg string) error {
	return errors.New(fmt.Sprintf("ERR at %d by %s.%s: %s", time.Now(), pkg, class, msg))
}
