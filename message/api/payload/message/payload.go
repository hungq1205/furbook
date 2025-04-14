package payload

type CreateMessagePayload struct {
	GroupID int    `json:"group_id"`
	Content string `json:"content"`
}

type CreateDirectMessagePayload struct {
	OppUsername string `json:"opp_username"`
	Content     string `json:"content"`
}
