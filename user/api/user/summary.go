package user

import (
	"net/http"
	"strconv"
	"user/api/payload"
	"user/usecase/friend"
	"user/usecase/user"
	"user/util"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	userId, err := getUintParam(ctx, "userId")
	if err != nil {
		return
	}

	usr, err := userService.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userPresenter, err := UserEntityToPresenter(usr, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userPresenter})
}

func GetUserList(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body struct {
		UserIds []uint `json:"user_ids"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	users, err := userService.GetUsers(body.UserIds)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userPresenter, err := ListUserEntityToPresenter(users, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": userPresenter})
}

func CreateUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body payload.UserCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	usr, err := userService.CreateUser(body.Username, body.Avatar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userPresenter, err := UserEntityToPresenter(usr, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": userPresenter})
}

func UpdateUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body payload.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	usr, err := userService.UpdateUser(util.MustGetUserId(ctx), body.Avatar, body.Bio)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	userPresenter, err := UserEntityToPresenter(usr, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": userPresenter})
}

func DeleteUser(ctx *gin.Context, Service user.UseCase) {
	if err := Service.DeleteUser(util.MustGetUserId(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func GetFriendList(ctx *gin.Context, Service friend.UseCase) {
	userId, err := getUintParam(ctx, "userId")
	if err != nil {
		return
	}

	friendRequests, err := Service.GetFriends(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend list"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friends": friendRequests})
}

func GetFriendRequestList(ctx *gin.Context, Service friend.UseCase) {
	userId, err := getUintParam(ctx, "userId")
	if err != nil {
		return
	}

	friendRequests, err := Service.GetFriendRequests(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend request list"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friend_requests": friendRequests})
}

func CheckFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.SenderWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := Service.CheckFriendRequest(body.SenderID, util.MustGetUserId(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friend request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func CheckFriendship(ctx *gin.Context, Service friend.UseCase) {
	var body payload.CheckFriendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := Service.CheckFriendship(body.UserID, body.FriendID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friendship"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func SendFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.CheckFriendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := Service.SendFriendRequest(body.UserID, body.FriendID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send friend request"})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func DeleteFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.ReceiverWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := Service.DeleteFriendRequest(util.MustGetUserId(ctx), body.ReceiverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend request"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func DeleteIncomingFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.SenderWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := Service.DeleteFriendRequest(body.SenderID, util.MustGetUserId(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend request"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func DeleteFriend(ctx *gin.Context, Service friend.UseCase) {
	var body payload.FriendWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := Service.DeleteFriend(util.MustGetUserId(ctx), body.FriendID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func getUintParam(ctx *gin.Context, param string) (uint, error) {
	userIdUint, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID of param " + param})
		return 0, err
	}
	return uint(userIdUint), nil
}
