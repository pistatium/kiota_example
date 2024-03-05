package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pistatium/kiota_example/webapi/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	postRepository *PostRepository
}

type GetParams struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type PostListResponse struct {
	Posts []*Post `json:"posts"`
}

type ErrResponse struct {
	Detail string `json:"detail"`
}

// @title    ExampleWebAPI
// @version  1.0
// @BasePath /api/v1

// GetPosts
// @Summary Get List of Posts
// @Schemes
// @Description get posts
// @Tags posts
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} PostListResponse "OK"
// @Failure 500 {object} ErrResponse "internal server error"
// @Router /posts [get]
func (s *Server) GetPosts(g *gin.Context) {
	params := GetParams{}
	if err := g.ShouldBindQuery(&params); err != nil {
		g.JSON(http.StatusBadRequest, ErrResponse{Detail: "invalid params"})
		return
	}
	posts, err := s.postRepository.GetPosts(g, params.Offset, params.Limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrResponse{Detail: "internal server error"})
		return
	}
	g.JSON(http.StatusOK, PostListResponse{
		Posts: posts,
	})
}

// AddPost
// @Summary Create new post
// @Schemes
// @Description add post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body Post true "Post"
// @Success 201 {object} Post "created"
// @Failure 400 {object} ErrResponse "invalid params"
// @Failure 500 {object} ErrResponse "internal server error"
// @Router /posts [post]
func (s *Server) AddPost(g *gin.Context) {
	post := Post{}
	if err := g.ShouldBindJSON(&post); err != nil {
		g.JSON(http.StatusBadRequest, ErrResponse{Detail: err.Error()})
		return
	}
	if err := s.postRepository.AddPost(g, &post); err != nil {
		g.JSON(http.StatusInternalServerError, ErrResponse{Detail: "internal server error"})
		return
	}
	g.JSON(http.StatusCreated, post)
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	server := &Server{
		postRepository: &PostRepository{},
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{

		eg := v1.Group("/posts")
		{
			eg.GET("", server.GetPosts)
			eg.POST("", server.AddPost)
		}
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	fmt.Println("Open swagger UI: http://127.0.0.1:8080/docs/index.html")
	r.Run(":8080")
}
