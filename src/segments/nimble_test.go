package segments

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNimble(t *testing.T) {
	cases := []struct {
		Case           string
		ExpectedString string
		Version        string
	}{
		{Case: "Nimble 0.16.1", ExpectedString: "0.16.1", Version: "nimble v0.16.1 compiled at 2024-11-30 12:28:24"},
		{Case: "Nimble 0.15.0", ExpectedString: "0.15.0", Version: "nimble v0.15.0 compiled at 2024-10-15 09:45:12"},
		{Case: "Nimble 0.14.2", ExpectedString: "0.14.2", Version: "nimble v0.14.2 compiled at 2024-09-01 15:30:00"},
	}

	for _, tc := range cases {
		params := &mockedLanguageParams{
			cmd:           "nimble",
			versionParam:  "--version",
			versionOutput: tc.Version,
			extension:     "*.nim",
		}
		env, props := getMockedLanguageEnv(params)
		n := &Nimble{}
		n.Init(props, env)
		assert.True(t, n.Enabled(), fmt.Sprintf("Failed in case: %s", tc.Case))
		assert.Equal(t, tc.ExpectedString, renderTemplate(env, n.Template(), n), fmt.Sprintf("Failed in case: %s", tc.Case))
	}
}
