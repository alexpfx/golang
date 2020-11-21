package merge

type Merge struct {
	Iid            int    `json:"iid"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	State          string `json:"state"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	TargetBranch   string `json:"target_branch"`
	SourceBranch   string `json:"source_branch"`
	WebUrl         string `json:"web_url"`
	Author         User   `json:"author"`
	MergeCommitSha string `json:"merge_commit_sha"`
	Commit         Commit `json:"commit"`
}

type User struct {
	Username string `json:"username"`
}

type Commit struct {
	Id        string `json:"id"`
	Email     string `json:"author_email"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
}

type MRResult struct {
	Merge       Merge  `json:"merge"`
}

type MRErrResult struct {
	MergeId int    `json:"merge_id"`
	Err     string `json:"error"`
}

