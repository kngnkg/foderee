package usecase

import (
	"context"
	"time"

	"github.com/kngnkg/tunetrail/api/entity"
)

type UserUseCase struct {
	userRepo       UserRepository
	userFollowRepo UserFollowRepository
}

type UserResponse struct {
	UserId         entity.UserId
	DisplayId      string
	Name           string
	AvatarUrl      string
	Bio            string
	FollowersCount int
	FollowingCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		UserId:         user.UserId,
		DisplayId:      user.DisplayId,
		Name:           user.Name,
		AvatarUrl:      user.AvatarUrl,
		Bio:            user.Bio,
		FollowersCount: user.FollowersCount,
		FollowingCount: user.FollowingCount,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func (uc *UserUseCase) GetById(ctx context.Context, userId entity.UserId) (*UserResponse, error) {
	user, err := uc.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	resp := NewUserResponse(user)
	return resp, nil
}

func (uc *UserUseCase) GetFollowersById(ctx context.Context, sourceId, targetId entity.UserId, nextCursor string, limit int) ([]*UserResponse, error) {
	ufs, err := uc.userFollowRepo.GetByUserIds(ctx, sourceId, targetId)
	if err != nil {
		return nil, err
	}

	var followerIds []entity.UserId
	for _, uf := range ufs {
		followerIds = append(followerIds, uf.UserId)
	}

	users, err := uc.userRepo.GetByIds(ctx, followerIds)
	if err != nil {
		return nil, err
	}

	var resps []*UserResponse
	for _, user := range users {
		resps = append(resps, NewUserResponse(user))
	}

	return resps, nil
}