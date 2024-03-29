package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kngnkg/foderee/backend/entity"
	"github.com/kngnkg/foderee/backend/infra/repository"
	"github.com/kngnkg/foderee/backend/logger"
)

var ErrorUsernameAlreadyExists = errors.New("usecase: username already exists")

type UserUseCase struct {
	DB         repository.DBAccesser
	userRepo   UserRepository
	followRepo FollowRepository
}

func NewUserUseCase(db repository.DBAccesser, userRepo UserRepository, followRepo FollowRepository) *UserUseCase {
	return &UserUseCase{
		DB:         db,
		userRepo:   userRepo,
		followRepo: followRepo,
	}
}

type UserListResponse struct {
	Users      []*entity.User
	NextCursor entity.ImmutableId
}

func (uc *UserUseCase) ListUsers(ctx context.Context, immutableId entity.ImmutableId, limit int) (*UserListResponse, error) {
	// 次のページがあるかどうかを判定するために、limit+1件取得する
	users, err := uc.userRepo.ListUsers(ctx, uc.DB, immutableId, limit+1)
	if err != nil {
		return nil, err
	}

	nextCursor := ""
	if len(users) > limit {
		// limit を超えた最初の要素の id を取得
		nextCursor = string(users[limit].ImmutableId)
		// limit までの要素を取得
		users = users[:limit]
	}

	// フォロー・フォロワー数を取得する
	for _, u := range users {
		u.FollowingCount, err = uc.followRepo.GetFollowingCountByUserId(ctx, uc.DB, u.ImmutableId)
		if err != nil {
			return nil, err
		}

		u.FollowersCount, err = uc.followRepo.GetFollowerCountByUserId(ctx, uc.DB, u.ImmutableId)
		if err != nil {
			return nil, err
		}
	}

	resp := &UserListResponse{
		Users:      users,
		NextCursor: entity.ImmutableId(nextCursor),
	}
	return resp, nil
}

func (uc *UserUseCase) GetByUsername(ctx context.Context, username entity.Username) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByUsername(ctx, uc.DB, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	// フォロー・フォロワー数を取得する
	user.FollowingCount, err = uc.followRepo.GetFollowingCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	user.FollowersCount, err = uc.followRepo.GetFollowerCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetMe(ctx context.Context, immutableId entity.ImmutableId) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByImmutableId(ctx, uc.DB, immutableId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	// フォロー・フォロワー数を取得する
	user.FollowingCount, err = uc.followRepo.GetFollowingCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	user.FollowersCount, err = uc.followRepo.GetFollowerCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Store(ctx context.Context, immutableId entity.ImmutableId, email string) (*entity.User, error) {
	// メールアドレスのローカルパートを username および DisplayName として使用する
	localPart := email[:strings.Index(email, "@")]

	user := &entity.User{
		Username:    entity.Username(localPart),
		ImmutableId: immutableId,
		DisplayName: localPart,
		AvatarUrl:   "",
		Bio:         "",
	}

	tx, err := uc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user, err = uc.userRepo.StoreUser(ctx, tx, user)
	if err != nil {
		defer func() {
			if err := tx.Rollback(); err != nil {
				logger.FromContext(ctx).Error("failed to rollback transaction: %v", err)
			}
		}()

		if errors.Is(err, repository.ErrorUsernameAlreadyExists) {
			return nil, fmt.Errorf("%w: %w", ErrorUsernameAlreadyExists, err)
		}

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// フォロー・フォロワー数を取得する
	user.FollowingCount, err = uc.followRepo.GetFollowingCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	user.FollowersCount, err = uc.followRepo.GetFollowerCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, username entity.Username, immutableId entity.ImmutableId, displayName, avatarUrl, bio string) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByImmutableId(ctx, uc.DB, immutableId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found: immutableId=%s", immutableId)
	}

	if username != "" && user.Username != username {
		user.Username = username
	}
	if displayName != "" && user.DisplayName != displayName {
		user.DisplayName = displayName
	}
	if avatarUrl != "" && user.AvatarUrl != avatarUrl {
		user.AvatarUrl = avatarUrl
	}
	if bio != "" && user.Bio != bio {
		user.Bio = bio
	}

	tx, err := uc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user, err = uc.userRepo.UpdateUser(ctx, tx, user)
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

	// フォロー・フォロワー数を取得する
	user.FollowingCount, err = uc.followRepo.GetFollowingCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}

	user.FollowersCount, err = uc.followRepo.GetFollowerCountByUserId(ctx, uc.DB, user.ImmutableId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, username entity.Username, immutableId entity.ImmutableId) error {
	tx, err := uc.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := uc.userRepo.DeleteUser(ctx, tx, immutableId); err != nil {
		defer func() {
			if err := tx.Rollback(); err != nil {
				logger.FromContext(ctx).Error("failed to rollback transaction: %v", err)
			}
		}()

		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
