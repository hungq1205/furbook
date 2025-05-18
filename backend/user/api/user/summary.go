package user

import (
	"net/http"
	"user/api/payload"
	"user/usecase/friend"
	"user/usecase/user"
	"user/util"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	usr, err := userService.GetUser(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userPresenter, err := UserEntityToPresenter(usr, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusOK, userPresenter)
}

func GetUserList(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body payload.UserListRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	users, err := userService.GetUsers(body.Usernames)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usersPresenter, err := ListUserEntityToPresenter(users, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usersPresenter)
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

	ctx.JSON(http.StatusCreated, userPresenter)
}

func UpdateUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body payload.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	usr, err := userService.UpdateUser(util.MustGetUsername(ctx), body.Avatar, body.Bio)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	userPresenter, err := UserEntityToPresenter(usr, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusCreated, userPresenter)
}

func DeleteUser(ctx *gin.Context, Service user.UseCase) {
	if err := Service.DeleteUser(util.MustGetUsername(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func GetFriendList(ctx *gin.Context, Service friend.UseCase) {
	friendRequests, err := Service.GetFriends(util.MustGetUsername(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend list"})
		return
	}
	ctx.JSON(http.StatusOK, friendRequests)
}

func GetFriendRequestList(ctx *gin.Context, Service friend.UseCase) {
	friendRequests, err := Service.GetFriendRequests(util.MustGetUsername(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend request list"})
		return
	}
	ctx.JSON(http.StatusOK, friendRequests)
}

func CheckFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.SenderWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := Service.CheckFriendRequest(body.Sender, util.MustGetUsername(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friend request"})
		return
	}
	ctx.JSON(http.StatusOK, exists)
}

func CheckFriendship(ctx *gin.Context, Service friend.UseCase) {
	var body payload.CheckFriendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := Service.CheckFriendship(body.User, body.Friend)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friendship"})
		return
	}
	ctx.JSON(http.StatusOK, exists)
}

func SendFriendRequest(ctx *gin.Context, Service friend.UseCase) {
	var body payload.CheckFriendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := Service.SendFriendRequest(body.User, body.Friend); err != nil {
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
	if err := Service.DeleteFriendRequest(util.MustGetUsername(ctx), body.Receiver); err != nil {
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
	if err := Service.DeleteFriendRequest(body.Sender, util.MustGetUsername(ctx)); err != nil {
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
	if err := Service.DeleteFriend(util.MustGetUsername(ctx), body.Friend); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
