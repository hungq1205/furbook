package noti

import (
	"net/http"
	"noti/api/client"
	"noti/api/payload"
	noti "noti/usecase/noti"
	"noti/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNoti(ctx *gin.Context, notiService noti.UseCase) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	notification, err := notiService.GetNoti(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}
	if notification.Username != util.MustGetUsername(ctx) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to retrieve this notification"})
		return
	}
	ctx.JSON(http.StatusOK, notification)
}

func GetNotisOfUser(ctx *gin.Context, notiService noti.UseCase) {
	pagination := util.ExtractPagination(ctx)
	notis, err := notiService.GetNotisOfUser(util.MustGetUsername(ctx), pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notis)
}

func CreateNoti(ctx *gin.Context, notiService noti.UseCase, wsService client.WsClient) {
	var body payload.NotiCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	notification, err := notiService.CreateNoti(body.Username, body.Icon, body.Desc, body.Link)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}
	if err := wsService.SendNoti(notification); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification message thorugh socket: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, notification)
}

func CreateNotiToUsers(ctx *gin.Context, notiService noti.UseCase, wsService client.WsClient) {
	var body payload.NotiToUsersCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	notis, err := notiService.CreateNotiToUsers(body.Usernames, body.Icon, body.Desc, body.Link)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notifications"})
		return
	}
	for _, n := range notis {
		if wsService.SendNoti(n) != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification message thorugh socket"})
			return
		}
	}

	ctx.Status(http.StatusCreated)
}

func UpdateNoti(ctx *gin.Context, notiService noti.UseCase) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	var body payload.NotiUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	GetNoti(ctx, notiService)

	notification, err := notiService.UpdateNoti(id, body.Read)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notification"})
		return
	}

	ctx.JSON(http.StatusCreated, notification)
}

func DeleteNoti(ctx *gin.Context, notiService noti.UseCase) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	GetNoti(ctx, notiService)
	if err := notiService.DeleteNoti(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete notification"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
