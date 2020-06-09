package common_test

import (
	"app/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt2String(t *testing.T) {
	assert.Equal(t, "1", common.Int2String(1))
}

func TestString2Int(t *testing.T) {
	assert.Equal(t, 1, common.String2Int("1"))
}
