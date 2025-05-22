package post

import (
	"net/http"
	"post/api/client"
	"post/api/payload"
	"post/usecase/post"
	"post/util"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context, postService *post.Service, userClient client.UserClient) {
	ctx := c.Request.Context()
	postID := c.Param("postID")
	postObj, err := postService.GetPost(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pPost, err := PostEntityToPresenterWithClient(postObj, userClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, pPost)
}

func GetPostsOfUser(c *gin.Context, postService *post.Service, userClient client.UserClient) {
	ctx := c.Request.Context()
	username := c.Param("username")
	pagination := util.ExtractPagination(c)

	posts, err := postService.GetPostsOfUser(ctx, username, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pPosts, err := ListPostEntityToPresenterWithClient(posts, userClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, pPosts)
}

func GetPostsOfUsers(c *gin.Context, postService *post.Service, userClient client.UserClient) {
	var body payload.UsersPostsRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	pagination := util.ExtractPagination(c)

	posts, err := postService.GetPostsOfUsers(ctx, body.Usernames, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pPosts, err := ListPostEntityToPresenterWithClient(posts, userClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, pPosts)
}

func CreateBlogPost(c *gin.Context, postService *post.Service) {
	var body payload.CreateBlogPostPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	newPost, err := postService.CreateBlogPost(ctx, util.MustGetUsername(c), body.Content, body.Medias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newPost)
}

func CreateLostPetPost(c *gin.Context, postService *post.Service, userClient client.UserClient) {
	var body payload.CreateLostPetPostPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := util.MustGetUsername(c)

	ctx := c.Request.Context()
	newPost, err := postService.CreateLostPetPost(ctx, username, body.ContactInfo, body.Type, body.Content, body.Medias, &body.Area, &body.LastSeen, body.LostAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pPost, err := PostEntityToPresenterWithClient(newPost, userClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting post: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pPost)
}

func PatchContentPost(c *gin.Context, postService *post.Service) {
	var body payload.PatchContentPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := util.MustGetUsername(c)

	ctx := c.Request.Context()
	isOwner, err := postService.CheckOwnership(ctx, username, c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this post"})
		return
	}

	if postService.PatchContent(ctx, c.Param("postID"), body.Content, body.Medias) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func PatchLostFoundStatus(c *gin.Context, postService *post.Service) {
	var body payload.PatchLostFoundStatus
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	isOwner, err := postService.CheckOwnership(ctx, util.MustGetUsername(c), c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this post"})
		return
	}

	err = postService.PatchLostFoundStatus(ctx, c.Param("postID"), body.IsResolved)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeletePost(c *gin.Context, postService *post.Service) {
	var body payload.DeletePostPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	isOwner, err := postService.CheckOwnership(ctx, util.MustGetUsername(c), body.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this post"})
		return
	}

	err = postService.DeletePost(ctx, body.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}

func GetComments(c *gin.Context, postService *post.Service, userClient client.UserClient) {
	ctx := c.Request.Context()
	comments, err := postService.GetComments(ctx, c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pComments, err := ListCommentEntityToPresenterWithClient(comments, userClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting comments: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": pComments})
}

func CreateComment(c *gin.Context, postService *post.Service) {
	var body payload.CreateCommentPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	err := postService.CreateComment(ctx, c.Param("postID"), util.MustGetUsername(c), body.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func UpsertInteraction(c *gin.Context, postService *post.Service) {
	var body payload.UpsertInteractionPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	err := postService.UpsertInteraction(ctx, c.Param("postID"), util.MustGetUsername(c), body.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func DeleteInteraction(c *gin.Context, postService *post.Service) {
	ctx := c.Request.Context()
	if postService.DeleteInteraction(ctx, c.Param("postID"), util.MustGetUsername(c)) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete interaction"})
		return
	}
	c.Status(http.StatusNoContent)
}
