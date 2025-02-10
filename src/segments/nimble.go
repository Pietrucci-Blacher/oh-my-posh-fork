package segments

type Nimble struct {
	language
}

func (n *Nimble) Template() string {
	return languageTemplate
}

func (n *Nimble) Enabled() bool {
	n.extensions = []string{"*.nim", "*.nimble"}
	n.commands = []*cmd{
		{
			executable: "nimble",
			args:       []string{"--version"},
			regex:      `nimble v(?P<version>((?P<major>[0-9]+)\.(?P<minor>[0-9]+)\.(?P<patch>[0-9]+)))`,
		},
	}

	return n.language.Enabled()
}
