package directory_test

import (
	"testing"

	"github.com/hansrodtang/runcom/backends/directory"
)

const TestDirectory = "./test/"

//
func TestInit(t *testing.T) {
	d := directory.NewBackend(TestDirectory)

}
