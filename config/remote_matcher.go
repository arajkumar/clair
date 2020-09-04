package config

//type RemoteMatcher struct {
// A remote service
//RemoteMatchers []RemoteMatchers
//}
//unused now might use in case of a different design
type RemoteMatcher struct {
	//Type         string `yaml:"- crda" json:"- crda"`
	Matchers map[string]MatcherParam
}

//MatcherParam takes parameters
type MatcherParam struct {
	// A remote service connection string.
	//
	// Formats:
	// url: "host=https://foo.com premiumKey=abced1234"
	Params string `yaml:"params" json:"params"`
}
