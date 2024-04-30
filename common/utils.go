package common

import (
	"crypto/sha1"
	"fmt"
)

func InvertArray[S ~[]E, E comparable](arr S) S {
	var result S
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

func GenerateHash(name string) string {
	data := fmt.Sprintf("%s - Plex Media Server", name)
	hash := sha1.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	hexString := fmt.Sprintf("%x", hashBytes)
	return hexString
}

func SetDifference[S ~[]E, E comparable](a, b S) S {

	mb := make(map[any]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff S
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
