package services

import (
	"context"

	TermChatProtos "../../common/generatedCode"
)

// UserInteraction : TODO: Comment here
type UserInteraction struct {
	TermChatProtos.UnimplementedUserInteractionServer
}

// GetFriends : Comment
func (this *UserInteraction) GetFriends(context.Context, *TermChatProtos.GetFriendsRequest) (*TermChatProtos.GetFriendsResponse, error) {
	// validate user token
	// get friends
	// prepare response
	return &TermChatProtos.GetFriendsResponse{}, nil
}

// AddFriend : Comment
func (this *UserInteraction) AddFriend(context.Context, *TermChatProtos.AddFriendRequest) (*TermChatProtos.AddFriendResponse, error) {
	// validate user token
	// add friend
	// prepare response
	return &TermChatProtos.AddFriendResponse{}, nil
}

// GetFriends : Comment
func (this *UserInteraction) RemoveFriends(context.Context, *TermChatProtos.RemoveFriendRequest) (*TermChatProtos.RemoveFriendResponse, error) {
	// validate user token
	// remove friend
	// prepare response
	return &TermChatProtos.RemoveFriendResponse{}, nil
}

// ListUsers : Comment
func (this *UserInteraction) ListUsers(context.Context, *TermChatProtos.ListUsersRequest) (*TermChatProtos.ListUsersResponse, error) {
	// validate user token
	// list users
	// prepare response
	return &TermChatProtos.ListUsersResponse{}, nil
}
