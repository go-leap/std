package ustd

import (
	"github.com/go-forks/go-wyhash"
	"github.com/go-forks/xxhash"
)

// HashTwo computes at least `newSum1` and (if equal to non-0 `oldSum1`) also `newSum2`. The first uses wyhash, the second xxhash.
func HashTwo(oldSum1, oldSum2 uint64, data []byte) (newSum1, newSum2 uint64, bothUnchanged bool) {
	hash1, forcecalcboth := wyhash.Hash, (oldSum1 == 0 || oldSum2 == 0)
	newSum1 = hash1(data, uint64(len(data)))
	if same1 := (oldSum1 == newSum1); same1 || forcecalcboth {
		var hash2 = xxhash.New()
		_, _ = hash2.Write(data) // as long as it' xxhash, no err ever here.
		newSum2 = hash2.Sum64()
		same2 := (oldSum2 == newSum2)
		bothUnchanged = same1 && same2
	}
	return
}
