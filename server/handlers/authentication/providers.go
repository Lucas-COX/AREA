package authentication

import "fmt"

type Provider int

const (
	Google Provider = iota + 1
	Discord
	Spotify
	Github
)

func (p Provider) String() string {
	var stringMap = map[Provider]string{
		Google:  "google",
		Discord: "discord",
		Spotify: "spotify",
		Github:  "github",
	}
	return stringMap[p]
}

func Parse(s string) (Provider, error) {
	var providersMap = map[string]Provider{
		"google":  Google,
		"discord": Discord,
		"spotify": Spotify,
		"github":  Github,
	}
	if p, ok := providersMap[s]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid provider: %s", s)
}
