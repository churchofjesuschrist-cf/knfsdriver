package voltoolshttp_test

import (
	"fmt"
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	"testing"
)

var debugServerAddress string
var localDriverPath string

var fakedriverServerPort int
var fakedriverProcess ifrit.Process
var tcpRunner *ginkgomon.Runner

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NFS Volume Tools Remote Client and Handlers Suite")
}

// testing support types:

type errCloser struct{ io.Reader }

func (errCloser) Close() error                     { return nil }
func (errCloser) Read(p []byte) (n int, err error) { return 0, fmt.Errorf("any") }

type stringCloser struct{ io.Reader }

func (stringCloser) Close() error { return nil }
