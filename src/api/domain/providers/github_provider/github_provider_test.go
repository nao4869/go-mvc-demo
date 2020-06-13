package github_provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestDefer(t *testing.T) {
	// defer works as stack 
	// 3 - 2 - 1
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("function's body")
}
