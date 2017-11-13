package model

import "time"

type Github struct{}

type Repositories struct {
	TotalCount       int64        `json:"total_count"`
	ImcompleteResult bool         `json:"imcomplete_result"`
	Items            []Repository `json:"items"`
}

type Repository struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Owner           Owner     `json:"owner"`
	Private         bool      `json:"private"`
	HtmlUrl         string    `json:"html_url"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	Url             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdateAt        time.Time `json:"update_at"`
	PushedAt        time.Time `json:"pushed_at"`
	HomePage        string    `json:"home_page"`
	Size            int64     `json:"size"`
	StarGazarsCount int64     `json:"star_gazars_count"`
	WatchersCount   int64     `json:"watchers_count"`
	Language        string    `json:"language"`
	ForkCount       int64     `json:"fork_count"`
	OpenIssuesCount int64     `json:"open_issues_count"`
	MasterBranch    string    `json:"master_branch"`
	DefaultBranch   string    `json:"default_branch"`
	Score           float64   `json:"score"`
}

type Owner struct {
	Login            string `json:"login"`
	ID               int64  `json:"id"`
	AvatarUrl        string `json:"avatar_url"`
	GravatarID       string `json:"gravatar_id"`
	URL              string `json:"url"`
	ReceivedEventUrl string `json:"received_event_url"`
	Type             string `json:"type"`
}
