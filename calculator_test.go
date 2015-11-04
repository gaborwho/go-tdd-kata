package calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd_StringZero_ReturnsZeroInt(t *testing.T) {
	result, err := Add("0")
	require.Nil(t, err)
	assert.Equal(t, 0, result)
}

func TestAdd_StringFive_ReturnsFiveInt(t *testing.T) {
	result, err := Add("5")
	require.Nil(t, err)
	assert.Equal(t, 5, result)
}

func TestAdd_StringTwentyFive_ReturnsTwentyFiveInt(t *testing.T) {
	result, err := Add("25")
	require.Nil(t, err)
	assert.Equal(t, 25, result)
}

func TestAdd_StringOneAndTwo_ReturnsThree(t *testing.T) {
	result, err := Add("1,2")
	require.Nil(t, err)
	assert.Equal(t, 3, result)
}

func TestAdd_StringTwentyTwoAndThree_ReturnsTwentyFive(t *testing.T) {
	result, err := Add("22,3")
	require.Nil(t, err)
	assert.Equal(t, 25, result)
}

func TestAdd_StringWithThreeValues_ReturnsTheirSum(t *testing.T) {
	result, err := Add("1,12,43")
	require.Nil(t, err)
	assert.Equal(t, 56, result)
}

func TestAdd_StringValuesSeparatedByNewLine_ReturnsTheirSum(t *testing.T) {
	result, err := Add("5\n12")
	require.Nil(t, err)
	assert.Equal(t, 17, result)
}

func TestAdd_StringWithNonNumberFirst_ReturnsError(t *testing.T) {
	_, err := Add("a,1")
	require.NotNil(t, err)
}

func TestAdd_StringWithNonNumberSecond_ReturnsError(t *testing.T) {
	_, err := Add("1,a")
	require.NotNil(t, err)
}
