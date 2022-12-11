package cathoderaytube_test

import (
	"testing"

	cathoderaytube "github.com/mhdiiilham/adventofcode/2022/cathoderay-tube"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/cathode_test.txt"

func TestGetSumOf6SignalStrengths(t *testing.T) {
	assertion := assert.New(t)
	expected := 13140

	client, err := cathoderaytube.New(input)
	assertion.NoError(err)
	assertion.NotNil(client)

	actual := client.GetSumOf6SignalStrengths()
	assertion.Equal(expected, actual)
}
