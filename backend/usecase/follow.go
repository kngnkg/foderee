package usecase

import (
	"context"
	"fmt"

	"github.com/kngnkg/foderee/backend/entity"
	"github.com/kngnkg/foderee/backend/infra/repository"
	"github.com/kngnkg/foderee/backend/logger"
)

type FollowUseCase struct {
	DB repository.DBAccesser
	ur UserRepository
	fr FollowRepository
}

func NewFollowUseCase(db repository.DBAccesser, ur UserRepository, fr FollowRepository) *FollowUseCase {
	return &FollowUseCase{
		DB: db,
		ur: ur,
		fr: fr,
	}
}

type FollowResponse struct {
	User        *entity.User
	IsFollowing bool
}

func (uc *FollowUseCase) ListFollows(ctx context.Context, immutableId entity.ImmutableId, usernames []entity.Username) ([]*FollowResponse, error) {
	us, err := uc.ur.ListUsersByUsername(ctx, uc.DB, usernames)
	if err != nil {
		return nil, err
	}

	resps := make([]*FollowResponse, 0, len(us))

	for _, u := range us {
		// フォローしているか確認する
		isFollwing, err := uc.fr.IsFollowing(ctx, uc.DB, immutableId, u.ImmutableId)
		if err != nil {
			return nil, err
		}

		resps = append(resps, &FollowResponse{
			User:        u,
			IsFollowing: isFollwing,
		})
	}

	return resps, nil
}

func (uc *FollowUseCase) Follow(ctx context.Context, immutableId entity.ImmutableId, targetUsername entity.Username) (*FollowResponse, error) {
	// ユーザーの取得
	targetUser, err := uc.ur.GetUserByUsername(ctx, uc.DB, targetUsername)
	if err != nil {
		return nil, err
	}
	if targetUser == nil {
		return nil, fmt.Errorf("user not found: username=%s", targetUsername)
	}
	if targetUser.ImmutableId == immutableId {
		return nil, fmt.Errorf("cannot follow myself: username=%s", targetUsername)
	}

	tx, err := uc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// フォローする
	_, err = uc.fr.StoreFollow(ctx, tx, &entity.Follow{
		ImmutableId:         immutableId,
		FolloweeImmutableId: targetUser.ImmutableId,
	})
	if err != nil {
		defer func() {
			if err := tx.Rollback(); err != nil {
				logger.FromContext(ctx).Error("failed to rollback transaction: %v", err)
			}
		}()

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &FollowResponse{
		User:        targetUser,
		IsFollowing: true,
	}, nil
}

func (uc *FollowUseCase) Unfollow(ctx context.Context, immutableId entity.ImmutableId, targetUsername entity.Username) (*FollowResponse, error) {
	// ユーザーの取得
	targetUser, err := uc.ur.GetUserByUsername(ctx, uc.DB, targetUsername)
	if err != nil {
		return nil, err
	}
	if targetUser == nil {
		return nil, fmt.Errorf("user not found: username=%s", targetUsername)
	}
	if targetUser.ImmutableId == immutableId {
		return nil, fmt.Errorf("cannot unfollow myself: username=%s", targetUsername)
	}

	tx, err := uc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// フォローを削除する
	_, err = uc.fr.DeleteFollow(ctx, tx, &entity.Follow{
		ImmutableId:         immutableId,
		FolloweeImmutableId: targetUser.ImmutableId,
	})
	if err != nil {
		defer func() {
			if err := tx.Rollback(); err != nil {
				logger.FromContext(ctx).Error("failed to rollback transaction: %v", err)
			}
		}()

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &FollowResponse{
		User:        targetUser,
		IsFollowing: false,
	}, nil
}