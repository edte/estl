// @program:     tiny-stl
// @file:        hash_test.go.go
// @author:      edte
// @create:      2021-12-25 15:37
// @description:
package hash

import (
	"fmt"
	"testing"
)

func TestBKDRHash(t *testing.T) {
	fmt.Println(BKDRHash([]byte("aaaaifjapfpoassdfasdf")))

}

func TestMurmurHash64A(t *testing.T) {
	fmt.Println(MurmurHash64([]byte("jjjjjasdofjaosifoiwenfwe"), 0x1234ABCD))
	fmt.Println(MurmurHash64([]byte("234jonasfjaosidjfoijweoi"), 0x1234ABCD))
	fmt.Println(MurmurHash64([]byte("32noojansfojoi982q4/asjfp"), 0x1234ABCD))
}
