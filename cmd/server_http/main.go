package main

import (
	"log"
	"net/http"

	"github.com/cnpog/grpc-example/pkg/http_blog"
	"github.com/labstack/echo/v4"
)

func main() {
	m := echo.New()
	m.POST("/blog", CreateBlogPost)
	log.Fatal(http.ListenAndServe(":8080", m))
}

func CreateBlogPost(ctx echo.Context) error {
	var r http_blog.BlogPostRequest

	if err := ctx.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := &http_blog.BlogPostResponse{
		Username:      r.Username,
		Title:         r.Title,
		Content:       r.Content,
		Unixtimestamp: r.Unixtimestamp,
		Created:       false,
	}
	return ctx.JSON(http.StatusCreated, res)
}
