package payload

type CreateMessagePayload struct {
	Content string `json:"content"`
}

type CreateDirectMessagePayload struct {
	OppUsername string `json:"opp_username"`
	Content     string `json:"content"`
}
