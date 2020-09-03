package config

//type RemoteMatcher struct {
// A remote service
//RemoteMatchers []RemoteMatchers
//}

type RemoteMatcher struct {
	// A remote service connection string.
	//
	// Formats:
	// url: "postgres://pqgotest:password@localhost/pqgotest? key: "dfdsjfdslfjdslkfjldsfk"
	Params string `yaml:"params" json:"params"`
}
