package presenter

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
}

type Group struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	LastMessage Message `json:"last_message"`
	IsDirect    bool    `json:"is_direct"`

	// Undirect
	OwnerName string   `json:"owner_name"`
	Members   []string `json:"members"`

	// Direct
	Avatar string `json:"avatar"`
}
