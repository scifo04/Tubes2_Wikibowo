package back

type LinkInfo struct {
	LinkValue string `json:"linkValue"`
	FinValue  string `json:"finValue"`
	IsOn      bool   `json:"isOn"`
	IsName    bool   `json:"isName"`
}

type Engine struct {
	Start string
	End string
	Depth int
	Namespace bool
	Cache map[string]bool
}

func createEngine(empty Engine, start string, end string, depth int, namespace bool) Engine {
	empty.Start = start
	empty.End = end
	empty.Depth = depth
	empty.Namespace = namespace
	empty.Cache = map[string]bool{}
	return empty
}