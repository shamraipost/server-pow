package proofofwork

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const targetBits = 20

type ProofOfWork struct {
	HashCash string
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.HashCash),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Validate(nonce int) bool {
	var hashInt big.Int
	target := big.NewInt(1)
	target.Lsh(target, uint(160-targetBits))

	data := pow.prepareData(nonce)
	hash := sha1.New()
	hash.Write(data)
	sha1Hash := hash.Sum(nil)
	hashInt.SetBytes(sha1Hash[:])

	isValid := hashInt.Cmp(target) == -1
	fmt.Printf("iteration: %d, found: %x, isValid: %t \n", nonce, sha1Hash, isValid)
	return isValid
}

func PrepareHashCash() string {
	randSource := rand.NewSource(time.Now().UnixNano())
	randNew := rand.New(randSource)
	rand := base64.StdEncoding.EncodeToString(IntToHex(int64(randNew.Intn(math.MaxInt32))))
	hashCash := strings.Join([]string{
		"1",
		strconv.Itoa(targetBits),
		strconv.FormatInt(time.Now().Unix(), 10),
		os.Getenv("NAME"),
		"",
		rand,
		"",
	}, ":")
	return hashCash
}
