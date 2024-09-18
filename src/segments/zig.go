package segments

import (
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime"
)

type Zig struct {
	language
}

func (z *Zig) Template() string {
	return languageTemplate
}

func (z *Zig) Init(props properties.Properties, env runtime.Environment) {
	z.language = language{
		env:        env,
		props:      props,
		extensions: []string{"*.zig", "build.zig"},
		commands: []*cmd{
			{
				executable: "zig",
				args:       []string{"version"},
				regex:      `(?P<version>[0-9]+\.[0-9]+\.[0-9]+)`,
			},
		},
	}
}

func (z *Zig) Enabled() bool {
	return z.language.Enabled()
}
