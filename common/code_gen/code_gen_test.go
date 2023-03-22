package code_gen

import "testing"

func TestGenStruct(t *testing.T) {
	GenStruct("user", "User")
	GenProtoMessage("user", "User")
}
