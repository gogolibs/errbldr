package errbldr_test

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
