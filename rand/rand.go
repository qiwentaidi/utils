package randutil

import (
	"bytes"
	rand2 "crypto/rand"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandFromChoices 从choices里面随机获取
func RandFromChoices(n int, choices string) string {
	b := make([]byte, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(choices) {
			b[i] = choices[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandLetters 随机小写字母
func RandLetters(n int) string {
	return RandFromChoices(n, letterBytes)
}

func RandomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand2.Int(rand2.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 包含上下限 [min, max]
func GetRandomIntWithAll(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return int(rand.Intn(max-min+1) + min)
}

// 不包含上限 [min, max)
func GetRandomIntWithMin(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return int(rand.Intn(max-min) + min)
}

// IntN returns a uniform random value in [0, max). It errors if max <= 0.
func IntN(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("max can't be <= 0")
	}
	nBig, err := rand2.Int(rand2.Reader, big.NewInt(int64(max)))
	if err != nil {
		return rand.Intn(max), nil
	}
	return int(nBig.Int64()), nil
}

// SleepRandTime returns a random duration between start and start+3 seconds (e.g. 2–5s)
func SleepRandTime(start float64) time.Duration {
	randomDelay := start + math.Round(rand.Float64()*30)/10 // 生成 start 到 start+3 之间的随机数（保留1位小数）
	return time.Duration(randomDelay * float64(time.Second))
}
