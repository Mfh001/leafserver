package test

import (
	"server/dbredigo"
	"testing"
)

func TestRedis(t *testing.T) {
	dbredigo.Start()
	_ = dbredigo.Set("11", 11, 11100)
}