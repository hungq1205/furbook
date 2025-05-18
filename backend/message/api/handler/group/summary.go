package group

import (
	"message/api/client"
	payload "message/api/payload/group"
	"message/usecase/group"
	"message/usecase/message"
	"message/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getGroup(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	groupIdParam := ctx.Param("groupId")
	groupId, err := strconv.Atoi(groupIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse groupId"})
		return
	}

	g, err := groupService.GetGroup(groupId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get group"})
		return
	}

	groupPresenter, err := groupEntityToPresenter(g, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupPresenter)
}

func createGroup(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	var body payload.CreateGroupPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse group"})
		return
	}

	g, err := groupService.CreateGroup(util.MustGetUsername(ctx), body.GroupName, body.Members)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupPresenter, err := groupEntityToPresenter(g, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, groupPresenter)
}

func updateGroup(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	groupIdParam := ctx.Param("groupId")
	groupId, err := strconv.Atoi(groupIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse groupId"})
		return
	}

	var body payload.UpdateGroupPayload
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse group update payload"})
		return
	}

	g, err := groupService.GetGroup(groupId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get group"})
		return
	}

	username := util.MustGetUsername(ctx)
	if username != g.OwnerName {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update group"})
		return
	}

	g, err = groupService.UpdateGroup(groupId, body.GroupName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupPresenter, err := groupEntityToPresenter(g, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupPresenter)
}

func deleteGroup(ctx *gin.Context, groupService group.UseCase) {
	var body payload.DeleteGroupPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse group delete payload"})
		return
	}

	g, err := groupService.GetGroup(body.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get group"})
		return
	}

	username := util.MustGetUsername(ctx)
	if username != g.OwnerName {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete group"})
		return
	}

	err = groupService.DeleteGroup(body.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func addMemberToGroup(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	var body payload.GroupMemberPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse group member payload"})
		return
	}

	g, err := groupService.GetGroup(body.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get group"})
		return
	}

	username := util.MustGetUsername(ctx)
	if username != g.OwnerName {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to add members"})
		return
	}

	g, err = groupService.AddMember(body.GroupID, body.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupPresenter, err := groupEntityToPresenter(g, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupPresenter)
}

func removeMemberToGroup(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	var body payload.GroupMemberPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse group member payload"})
		return
	}

	g, err := groupService.GetGroup(body.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get group"})
		return
	}

	username := util.MustGetUsername(ctx)
	if username != body.Username && username != g.OwnerName {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to remove this member"})
		return
	}

	g, err = groupService.RemoveMember(body.GroupID, body.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupPresenter, err := groupEntityToPresenter(g, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupPresenter)
}

func getGroupsOfUser(ctx *gin.Context, groupService group.UseCase, messageService message.UseCase) {
	username := util.MustGetUsername(ctx)
	pagination := util.ExtractPagination(ctx)

	groups, err := groupService.GetGroupsOfUser(username, pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupPresenters, err := groupListEntityToPresenter(groups, groupService, messageService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupPresenters)
}

func getMembersOfGroup(ctx *gin.Context, groupService group.UseCase, userClient client.UserClient) {
	groupID, err := strconv.Atoi(ctx.Param("groupId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse groupId"})
		return
	}

	usernames, err := groupService.GetMembers(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users, err := userClient.FindUsers(usernames)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
