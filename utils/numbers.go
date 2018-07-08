package utils

const (
	numberBytes   = `1234567890`
	numberIdxBits = 6                    // 6 bits to represent a number index
	numberIdxMask = 1<<numberIdxBits - 1 // All 1-bits, as many as numberIdxBits
	numberIdxMax  = 63 / numberIdxBits   // # of number indices fitting in 63 bits
)

// http://stackoverflow.com/a/31832326
func RandomNumbers(length int) string {
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), numberIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), numberIdxMax
		}
		if idx := int(cache & numberIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= numberIdxBits
		remain--
	}

	return string(b)
}
