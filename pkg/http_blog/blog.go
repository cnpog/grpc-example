package http_blog

type BlogPostRequest struct {
	Username      string
	Title         string
	Content       string
	Unixtimestamp int64
}
type BlogPostResponse struct {
	Username      string
	Title         string
	Content       string
	Unixtimestamp int64
	Created       bool
}
