package inethenge

type InetHenge struct {
	Nodes  []Node `json:"nodes"`
	Links  []Link `json:"links"`
	groups string `json:"group"`
}

type Node struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Link struct {
	Source string   `json:"source"`
	Target string   `json:"target"`
	Meta   MetaData `json:"meta"`
}

type MetaData struct {
	Interface Interface `json:"interface"`
	LoopBack  string    `json:"loopback"`
}

type Interface struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
