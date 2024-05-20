package lyra2re

import (
	"golang.org/x/crypto/sha3"
	
	"github.com/pedroalbanese/skein"
	"github.com/pedroalbanese/blake256"
	"github.com/pedroalbanese/groestl"
)

func Sum(data []byte) ([]byte, error) {
	blake := blake256.New()
	if _, err := blake.Write(data); err != nil {
		return nil, err
	}
	resultBlake := blake.Sum(nil)

	keccak := sha3.NewLegacyKeccak256()
	if _, err := keccak.Write(resultBlake); err != nil {
		return nil, err
	}
	resultKeccak := keccak.Sum(nil)

	lyra2Result := make([]byte, 32)
	lyra2(lyra2Result, resultKeccak, resultKeccak, 1, 8, 8)

	var skeinResult [32]byte
	skein.Sum256(&skeinResult, lyra2Result, nil)

	groestlResult := groest.Sum(skeinResult[:])

	return groestlResult[:], nil
}
