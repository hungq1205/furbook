package payload

type CreateGroupPayload struct {
	GroupName string `json:"group_name"`
	Username  string `json:"username"`
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
