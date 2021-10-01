// Copyright (c) 2014 Casey Marshall. See LICENSE file for details.

package basen_test

import (
	"math/big"
	"testing"

	"github.com/KEINOS/base-N-go/basen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoMultiByte(t *testing.T) {
	assert.PanicsWithValue(t, "multi-byte characters not supported",
		func() { basen.NewEncoding("世界") })
}

func TestRoundTrip62(t *testing.T) {
	for _, testCase := range []struct {
		rep string
		enc *basen.Encoding
		val []byte
	}{
		{"1", basen.Base62, []byte{1}},
		{"z", basen.Base62, []byte{61}},
		{"10", basen.Base62, []byte{62}},
		{"100", basen.Base62, big.NewInt(int64(3844)).Bytes()},
		{"zz", basen.Base62, big.NewInt(int64(3843)).Bytes()},

		{"Tgmc", basen.Base58, big.NewInt(int64(10002343)).Bytes()},
		{"if", basen.Base58, big.NewInt(int64(1000)).Bytes()},
		{"", basen.Base58, big.NewInt(int64(0)).Bytes()},
	} {
		rep := testCase.enc.EncodeToString(testCase.val)
		assert.Equal(t, testCase.rep, rep)

		val, err := testCase.enc.DecodeString(testCase.rep)
		assert.NoError(t, err)
		assert.EqualValues(t, testCase.val, val)
	}
}

func TestRand256(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := basen.Base62.MustRandom(32)

		// Should be 43 chars or less because math.log(2**256, 62) == 42.994887413002736
		assert.LessOrEqual(t, len(v), 43)
	}
}

func TestStringN(t *testing.T) {
	var (
		val []byte
		err error
	)

	val, err = basen.Base58.DecodeStringN("", 4)
	require.NoError(t, err)
	assert.EqualValues(t, []byte{0, 0, 0, 0}, val)

	// ensure round-trip with padding is right
	val, err = basen.Base62.DecodeStringN("10", 4)
	require.NoError(t, err)
	assert.EqualValues(t, []byte{0, 0, 0, 62}, val)

	rep := basen.Base62.EncodeToString(val)
	assert.Equal(t, "10", rep)
}
