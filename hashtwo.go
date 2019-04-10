package ustd

import (
	"github.com/go-forks/go-wyhash"
	"github.com/go-forks/xxhash"
)

func HashWyPlus(oldSumWy, oldSumOther uint64, data []byte) (newSumWy, newSumOther uint64, bothSame bool, err error) {
	if newSumWy = wyhash.Hash(data, uint64(len(data))); newSumWy == oldSumWy {
		var h2 = xxhash.New()
		if _, err = h2.Write(data); err == nil {
			newSumOther = h2.Sum64()
			bothSame = newSumOther == oldSumOther
		}
	}
	return
}
