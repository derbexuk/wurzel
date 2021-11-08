package upload

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

/*
  Test Main
*/
func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

func TestFetchGsheet(t *testing.T) {
	RegisterTestingT(t)

	err := FetchGsheet("test_poi.yaml")
	Expect(err).To(BeNil())
}
