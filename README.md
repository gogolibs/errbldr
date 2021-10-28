# errbldr #

[![GoDoc](https://godoc.org/github.com/gogolibs/errbldr?status.svg)](https://pkg.go.dev/github.com/gogolibs/errbldr)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogolibs/errbldr)](https://goreportcard.com/report/github.com/gogolibs/errbldr)
[![CI](https://github.com/gogolibs/errbldr/actions/workflows/test-and-coverage.yml/badge.svg)](https://github.com/gogolibs/errbldr/actions/workflows/test-and-coverage.yml)
[![codecov](https://codecov.io/gh/gogolibs/errbldr/branch/main/graph/badge.svg?token=JXSDP6Ifxi)](https://codecov.io/gh/gogolibs/errbldr)

**errbldr** is a convenience builder wrapper around [github.com/pkg/errors](https://github.com/pkg/errors).

```go
import (
	"errors"
	"github.com/gogolibs/errbldr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExample(t *testing.T) {
	err := errors.New("test error")
	err2 := errbldr.Wrap(err, "1st level message").Msgf("2nd level message param=%d", 42).Err()
	require.Equal(t, "2nd level message param=42: 1st level message: test error", err2.Error())
}
```
