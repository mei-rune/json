package json_test

import (
	"strings"
	"testing"

	"github.com/mei-rune/json"
)

type ExchangeResponse struct {
	Id uint64 `json:"request_id"`
	M  map[string]interface{}  `json:"m"`
}

func TestUnmarshalV2WithNumber(t *testing.T) {

	var in = strings.NewReader("{\"request_id\": 0, \"m\": { \"a\":  12 }}")

	var value ExchangeResponse

	err := json.UnmarshalV2WithNumber(in, &value)
	if err != nil {
		t.Error(err)
		return
	}

	a := value.M["a"]

	t.Logf("%T %v", a, a)

}
