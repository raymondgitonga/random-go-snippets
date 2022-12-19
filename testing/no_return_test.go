package testing

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	v := values{
		a:      3,
		b:      5,
		logger: log.Logger{},
	}
	v.logger.SetOutput(&buf)

	v.calculate()

	assert.Equal(t, "15", strings.TrimSpace(buf.String()))
}
