package config

//type RemoteMatcher struct {
// A remote service
//RemoteMatchers []RemoteMatchers
//}

type RemoteMatcher struct {
	// A remote service connection string.
	//
	// Formats:
	// url: "host=https://foo.com premiumKey=abced1234"
	Params string `yaml:"params" json:"params"`
}
