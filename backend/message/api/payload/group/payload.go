package payload

type CreateGroupPayload struct {
	GroupName string   `json:"group_name"`
	Members   []string `json:"members"`
}

type UpdateGroupPayload struct {
	GroupName string `json:"group_name"`
}

type DeleteGroupPayload struct {
	GroupID int `json:"group_id"`
}

type GroupMemberPayload struct {
	GroupID  int    `json:"group_id"`
	Username string `json:"username"`
}
