package xor

func Xor(raw, secret []byte) (result []byte) {
	result = make([]byte, len(raw))
	j := 0
	for i := 0; i < len(raw); i++ {
		result[i] = raw[i] ^ secret[j]
		j = (j + 1) % len(secret)
	}
	return
}
