package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) GetBlogWithAuthor(ctx context.Context, id int) (*model.BlogWithAuthor, error) {
	var result model.BlogWithAuthor

	query := `
		SELECT 
			b.id, b.title, b.content, b.author_id, b.image, b.status, 
			b.created_at, b.updated_at, b.deleted_at,
			u.id as author_id, u.full_name as author_full_name
		FROM blogs b 
		LEFT JOIN users u ON b.author_id = u.id 
		WHERE b.id = ? AND b.deleted_at IS NULL
	`

	row := s.db.Raw(query, id).Row()
	
	var authorFullName *string
	err := row.Scan(
		&result.Id, &result.Title, &result.Content, &result.AuthorId, &result.Image, &result.Status,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt,
		&result.Author.Id, &authorFullName,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	// Set author full name
	if authorFullName != nil {
		result.Author.FullName = *authorFullName
	}

	return &result, nil
}

