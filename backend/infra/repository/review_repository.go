package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kngnkg/tunetrail/backend/entity"
	"github.com/kngnkg/tunetrail/backend/logger"
)

type ReviewRepository struct{}

func (r *ReviewRepository) ListReviews(ctx context.Context, db Executor, reviewId string, limit int) ([]*entity.Review, error) {
	query := `
	SELECT review_id, user_id AS "author.user_id", album_id, title, content, published_status, created_at, updated_at
	FROM reviews WHERE published_status = 'published'`

	placeholderNum := 1
	args := []interface{}{}

	if reviewId != "" {
		query += fmt.Sprintf(" AND created_at <= (SELECT created_at FROM reviews WHERE review_id = $%d)", placeholderNum)
		args = append(args, reviewId)
		placeholderNum++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d;", placeholderNum)
	args = append(args, limit)

	reviews := []*entity.Review{}
	err := db.SelectContext(ctx, &reviews, query, args...)
	if err != nil {
		logger.FromContent(ctx).Error("failed to get reviews.", err)
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewById(ctx context.Context, db Executor, reviewId string) (*entity.Review, error) {
	query := `
		SELECT review_id, user_id AS "author.user_id", album_id , title, content, published_status, created_at, updated_at
		FROM reviews
		WHERE review_id = $1;`

	review := &entity.Review{}
	err := db.GetContext(ctx, review, query, reviewId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) StoreReview(ctx context.Context, db Executor, review *entity.Review) (*entity.Review, error) {
	query := `
		INSERT INTO reviews (user_id, album_id, title, content, published_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING review_id, user_id AS "author.user_id", album_id , title, content, published_status, created_at, updated_at;`

	err := db.QueryRowxContext(ctx, query, review.Author.ImmutableId, review.AlbumId, review.Title, review.Content, review.PublishedStatus).
		StructScan(review)
	if err != nil {
		logger.FromContent(ctx).Error("failed to store review.", err)
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, db Executor, review *entity.Review) (*entity.Review, error) {
	query := `
		UPDATE reviews
		SET title = $1, content = $2, published_status = $3, updated_at = NOW()
		WHERE review_id = $4
		RETURNING review_id, user_id AS "author.user_id", album_id , title, content, published_status, created_at, updated_at;`

	err := db.QueryRowxContext(ctx, query, review.Title, review.Content, review.PublishedStatus, review.ReviewId).
		StructScan(review)
	if err != nil {
		logger.FromContent(ctx).Error("failed to update review.", err)
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, db Executor, reviewId string) error {
	query := `DELETE FROM reviews WHERE review_id = $1;`

	_, err := db.ExecContext(ctx, query, reviewId)
	if err != nil {
		logger.FromContent(ctx).Error("failed to delete review.", err)
		return err
	}

	return nil
}