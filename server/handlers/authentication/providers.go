package authentication

import "fmt"

type Provider int

const (
	Google Provider = iota + 1
	Discord
	Timer
	Notion
	Github
	Microsoft
)

func (p Provider) String() string {
	var stringMap = map[Provider]string{
		Google:    "google",
		Discord:   "discord",
		Timer:     "timer",
		Notion:    "notion",
		Github:    "github",
		Microsoft: "microsoft",
	}
	return stringMap[p]
}

func Parse(s string) (Provider, error) {
	var providersMap = map[string]Provider{
		"google":    Google,
		"discord":   Discord,
		"Timer":     Timer,
		"notion":    Notion,
		"github":    Github,
		"microsoft": Microsoft,
	}
	if p, ok := providersMap[s]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid provider: %s", s)
}
