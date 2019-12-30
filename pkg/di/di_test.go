package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	c := New()
	err := c.registerFun(func() string { return "Hello" })
	assert.NoError(t, err)

	err = c.registerFun(func() int { return 9 })
	assert.NoError(t, err)

	err = c.execute(func(name string, number int) {
		assert.Equal(t, "Hello", name)
		assert.Equal(t, 9, number)
	})
	assert.NoError(t, err)
}

type Readable interface {
	Text() string
}

type book struct {
	message string
}

func (b book) Text() string {
	return b.message
}

func TestRegisterInterface(t *testing.T) {
	c := New()
	err := c.registerFun(func() Readable {
		return book{message: "This is an important message"}
	})
	assert.NoError(t, err)

	err = c.execute(func(r Readable) {
		assert.Equal(t, r.Text(), "This is an important message")
	})
	assert.NoError(t, err)
}
