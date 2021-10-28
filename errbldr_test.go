package errbldr

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestMsg(t *testing.T) {
	err := Msg("msg").Err()
	require.Equal(t, "msg", err.Error())
}

func TestMsgf(t *testing.T) {
	err := Msgf("msgf %d %d", 1, 2).Err()
	require.Equal(t, "msgf 1 2", err.Error())
}

func TestWrap(t *testing.T) {
	err := Wrap(errors.New("err"), "msg").Err()
	require.Equal(t, "msg: err", err.Error())
}

func TestWrapf(t *testing.T) {
	err := Wrapf(errors.New("err"), "msgf %d %d", 1, 2).Err()
	require.Equal(t, "msgf 1 2: err", err.Error())
}

func TestErrorBuilder_Msg(t *testing.T) {
	err := Msg("msg1").Msg("msg2").Err()
	require.Equal(t, "msg2: msg1", err.Error())
}

func TestErrorBuilder_Msgf(t *testing.T) {
	err := Msg("msg1").Msgf("msgf2 %d %d", 1, 2).Err()
	require.Equal(t, "msgf2 1 2: msg1", err.Error())
}
