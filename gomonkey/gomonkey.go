// https://stackoverflow.com/a/45379980/5228839

package monkey_test

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func doomed() {
	os.Exit(1)
}

func TestDoomed(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", doomed, "os.Exit was not called")
}
