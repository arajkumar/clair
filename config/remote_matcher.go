package config

type RemoteMatchers struct {
	// A remote service
	RemoteMatchers []MatcherObject
}
type MatcherObject struct {
	Name   string
	Params MatcherParams
}
type MatcherParams struct {
	// A remote service connection string.
	//
	// Formats:
	// url: "postgres://pqgotest:password@localhost/pqgotest? key: "dfdsjfdslfjdslkfjldsfk"
	Params string `yaml:"params" json:"params"`
}
