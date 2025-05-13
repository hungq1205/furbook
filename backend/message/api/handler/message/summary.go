package message

import (
	"github.com/gin-gonic/gin"
	payload "message/api/payload/message"
	"message/usecase/group"
	"message/usecase/message"
	"message/util"
	"net/http"
	"strconv"
)

func checkMembership(ctx *gin.Context, groupID int, groupService group.UseCase) {
	membership, err := groupService.CheckMembership(ctx.MustGet("username").(string), groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !membership {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not a member of this group"})
		return
	}
}

func getGroupMessageList(ctx *gin.Context, messageService message.UseCase, groupService group.UseCase) {
	groupIdParam := ctx.Param("groupID")

	groupID, err := strconv.Atoi(groupIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkMembership(ctx, groupID, groupService)

	messages, err := messageService.GetGroupMessageList(groupID, util.ExtractPagination(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messageListEntityToPresenter(messages))
}

func getDirectMessageList(ctx *gin.Context, messageService message.UseCase) {
	oppUsername := ctx.Query("oppUsername")

	if oppUsername == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No oppUsername provided"})
		return
	}

	messages, err := messageService.GetDirectMessageList(ctx.MustGet("username").(string), oppUsername, util.ExtractPagination(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messageListEntityToPresenter(messages))
}

func createGroupMessage(ctx *gin.Context, messageService message.UseCase, groupService group.UseCase) {
	var body payload.CreateMessagePayload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkMembership(ctx, body.GroupID, groupService)

	msg, err := messageService.SendMessage(ctx.MustGet("username").(string), body.Content, body.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, messageEntityToPresenter(msg))
}

func createDirectMessage(ctx *gin.Context, messageService message.UseCase, groupService group.UseCase) {
	var body payload.CreateDirectMessagePayload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := messageService.SendDirectMessage(ctx.MustGet("username").(string), body.OppUsername, body.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, messageEntityToPresenter(msg))
}
