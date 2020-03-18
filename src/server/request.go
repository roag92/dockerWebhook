package server

type Payload struct {
	PushData PushData `json:"push_data"`
	CallbackUrl string `json:"callback_url"`
	Repository Repository
}

type PushData struct {
	PushedAt int64 `json:"pushed_at"`
	Images []string
	Tag string
	Pusher string
}

type Repository struct {
	Status string
	Description string
	IsTrusted bool `json:"is_trusted"`
	FullDescription string `json:"full_description"`
	RepoUrl string `json:"repo_url"`
	Owner string
	IsOfficial bool `json:"is_official"`
	IsPrivate bool `json:"is_private"`
	Name string
	Namespace string
	StartCount int64 `json:"star_count"`
	CommentCount int64 `json:"comment_count"`
	DateCreated int64 `json:"date_created"`
	RepoName string `json:"repo_name"`
}
