package main

import (
	"github.com/stretchr/testify/assert"
	"net"
	"os"
	"strconv"
	"testing"
)

func TestServer(t *testing.T) {

	t.Run("test with no port number environment variable fails", func(t *testing.T) {
		err := os.Unsetenv("MONGODB_PORT_NUMBER")
		assert.NoError(t, err)

		assert.PanicsWithError(t, "port environment not set", func() {
			main()
		}, "expected panic from missing environment variable")
	})

	t.Run("test with bad server", func(t *testing.T) {
		// Lets get a random unused port
		l, err := net.Listen("tcp", "localhost:")
		assert.NoError(t, err, "error creating dummy listener")
		defer func() { _ = l.Close() }()

		err = os.Setenv("MONGODB_PORT_NUMBER", strconv.Itoa(l.Addr().(*net.TCPAddr).Port))
		assert.NoError(t, err)

		defer func() { _ = os.Unsetenv("MONGODB_PORT_NUMBER") }()
		assert.PanicsWithError(t, "failed to connect to mongo", func() {
			main()
		}, "expected panic from connection timeout")
	})

	t.Run("test with no server", func(t *testing.T) {
		// Lets get a random unused port
		l, err := net.Listen("tcp", "localhost:")
		assert.NoError(t, err, "error creating dummy listener")
		_ = l.Close()

		err = os.Setenv("MONGODB_PORT_NUMBER", strconv.Itoa(l.Addr().(*net.TCPAddr).Port))
		assert.NoError(t, err)
		defer func() { _ = os.Unsetenv("MONGODB_PORT_NUMBER") }()
		assert.PanicsWithError(t, "failed to connect to mongo", func() {
			main()
		}, "expected panic from connection timeout")
	})
}
