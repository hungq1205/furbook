package user

import (
	"net/http"
	"user-service/api/payload"
	"user-service/usecase/friend"
	"user-service/usecase/user"
	"user-service/util"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	username := ctx.Param("username")
	user, err := userService.GetUser(username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userPresenter, err := UserEntityToPresenter(user, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userPresenter})
}

func GetUserList(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body struct {
		Usernames []string `json:"usernames"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	users, err := userService.GetUsers(body.Usernames)
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

func CheckUsernameExists(ctx *gin.Context, userService user.UseCase) {
	username := ctx.Param("username")
	exists, err := userService.CheckUsernameExists(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check username existence"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func CreateUser(ctx *gin.Context, userService user.UseCase, friendService friend.UseCase) {
	var body payload.UserCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := userService.CreateUser(body.Username, body.Avatar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userPresenter, err := UserEntityToPresenter(user, friendService)
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

	user, err := userService.UpdateUser(util.MustGetUsername(ctx), body.Avatar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	userPresenter, err := UserEntityToPresenter(user, friendService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": userPresenter})
}

func DeleteUser(ctx *gin.Context, userService user.UseCase) {
	if err := userService.DeleteUser(util.MustGetUsername(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func GetFriendList(ctx *gin.Context, friendService friend.UseCase) {
	friendRequests, err := friendService.GetFriends(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend list"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friends": friendRequests})
}

func GetFriendRequestList(ctx *gin.Context, friendService friend.UseCase) {
	friendRequests, err := friendService.GetFriendRequests(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friend request list"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friend_requests": friendRequests})
}

func CheckFriendRequest(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.SenderWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := friendService.CheckFriendRequest(body.Sender, util.MustGetUsername(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friend request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func CheckFriendship(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.CheckFriendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	exists, err := friendService.CheckFriendship(body.Username, body.Friend)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friendship"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func SendFriendRequest(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.ReceiverWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := friendService.SendFriendRequest(util.MustGetUsername(ctx), body.Receiver); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send friend request"})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func DeleteFriendRequest(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.ReceiverWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := friendService.DeleteFriendRequest(util.MustGetUsername(ctx), body.Receiver); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend request"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func DeleteIncomingFriendRequest(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.SenderWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := friendService.DeleteFriendRequest(body.Sender, util.MustGetUsername(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend request"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func DeleteFriend(ctx *gin.Context, friendService friend.UseCase) {
	var body payload.FriendWrapper
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := friendService.DeleteFriend(util.MustGetUsername(ctx), body.Friend); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
