package llb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMeta(t *testing.T) {
	t.Parallel()

	s := Image("foo")
	s = s.AddEnv("BAR", "abc").Dir("/foo/bar")

	v, ok := s.GetEnv("BAR")
	assert.True(t, ok)
	assert.Equal(t, "abc", v)

	assert.Equal(t, "/foo/bar", s.GetDir())

	s2 := Image("foo2")
	s2 = s2.AddEnv("BAZ", "def").Reset(s)

	_, ok = s2.GetEnv("BAZ")
	assert.False(t, ok)

	v, ok = s2.GetEnv("BAR")
	assert.True(t, ok)
	assert.Equal(t, "abc", v)
}

func TestAddEnv(t *testing.T) {
	s := NewState(nil)
	s = s.AddEnv("key1", "val1")
	s = s.AddEnv("key2", "val2")
	s = s.AddEnv("key1", "value1")

	v, ok := s.GetEnv("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", v)

	v, ok = s.GetEnv("key2")
	assert.True(t, ok)
	assert.Equal(t, "val2", v)

	_, ok = s.GetEnv("key3")
	assert.False(t, ok)
}
