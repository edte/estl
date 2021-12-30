// @program:     tiny-stl
// @file:        hash.go
// @author:      edte
// @create:      2021-12-25 15:18
// @description:
package algorithm

func BKDRHash(str string) (res uint64) {
	for i := range str {
		res += res*131 + (uint64)(str[i])
	}
	return
}
