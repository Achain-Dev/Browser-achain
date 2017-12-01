package graph_verification

import (
	"testing"
	"fmt"
)

func TestEncodeCodeToBash64(t *testing.T) {
	code := RandText(4)
	fmt.Printf("code=%s\n",code)
	//EncodeCodeToBash64(code)
	codeBash64 := EncodeCodeToBash64(code)
	fmt.Printf("codeBash64=%s\n",codeBash64)


}
