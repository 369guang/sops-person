package auth

import "testing"

// 7IF7IAKX66ZR4T4XOFR3OQM6K3TPQWEN
func TestCreateKey(t *testing.T) {
	CreateKey("test")
}

func TestValidate(t *testing.T) {
	Validate("325965", "7IF7IAKX66ZR4T4XOFR3OQM6K3TPQWEN")
}
