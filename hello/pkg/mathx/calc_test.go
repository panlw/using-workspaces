package mathx_test

import (
	"testing"

	"github.com/panlw/using-workspaces/hello/pkg/mathx"
)

func Test_Add(t *testing.T) {
	ret := mathx.Add(1, 2)
	exp := float64(3)
	if exp != ret {
		t.Fail()
	}
}
