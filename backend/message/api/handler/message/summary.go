package message

import (
	"message/api/client"
	payload "message/api/payload/message"
	"message/usecase/group"
	"message/usecase/message"
	"message/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func createGroupMessage(ctx *gin.Context, messageService message.UseCase, groupService group.UseCase, wsClient client.WsClient) {
	groupIdParam := ctx.Param("groupID")
	groupID, err := strconv.Atoi(groupIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body payload.CreateMessagePayload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkMembership(ctx, groupID, groupService)

	msg, err := messageService.SendMessage(ctx.MustGet("username").(string), body.Content, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsClient.SendMessage(msg.ID, ctx.MustGet("username").(string), groupID, body.Content, msg.CreatedAt)

	ctx.JSON(http.StatusCreated, messageEntityToPresenter(msg))
}
