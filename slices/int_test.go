package slices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Integer_JSON_Unmarshal(t *testing.T) {
	r := require.New(t)

	x := `[1, 2, 3]`
	i := Int{}
	r.NoError(json.Unmarshal([]byte(x), &i))
	r.Equal(Int{1, 2, 3}, i)
}

func Test_Integer_JSON_Marshal(t *testing.T) {
	r := require.New(t)

	x := `[1,2,3]`
	i := Int{1, 2, 3}
	b, err := json.Marshal(i)
	r.NoError(err)
	r.Equal(x, string(b))
}
