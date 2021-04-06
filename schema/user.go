package schema

type PutUserCli struct {
	UserID      int         `json:"user_id"`
	InputSecret InputSecret `json:"input_secret"`
}
type InputSecret struct {
	Secret string `json:"secret"`
}
