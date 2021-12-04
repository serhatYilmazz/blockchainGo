package util

import "crypto/sha512"

func DeriveHash(data []byte) [64]byte {
	return sha512.Sum512(data)
}
