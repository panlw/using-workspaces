package stringx_test

import (
	"testing"

	"github.com/panlw/using-workspaces/world/pkg/stringx"
)

func Test_ToUpper(t *testing.T) {
	ret := stringx.ToUpper("hello")
	exp := "HELLO"
	if exp != ret {
		t.Fail()
	}
}
