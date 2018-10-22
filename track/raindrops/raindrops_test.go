package raindrops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/testhelper"
	"github.com/tehsphinx/exalysis/track/raindrops/tpl"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/1", suggestion: tpl.ManyLoops, expected: true},
	{path: "./solutions/1", suggestion: tpl.ConcatNotNeeded, expected: false},
	{path: "./solutions/2", suggestion: tpl.ConcatNotNeeded, expected: true},
	{path: "./solutions/3", suggestion: tpl.StringsBuilder, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "raindrops")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}