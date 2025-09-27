package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) GetBlogWithAuthor(ctx context.Context, id int) (*model.BlogWithAuthor, error) {
	var result model.BlogWithAuthor

	// Lấy blog trước
	err := s.db.Where("id = ?", id).First(&result.Blog).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	// Lấy thông tin author
	query := `SELECT id, full_name FROM users WHERE id = ?`
	row := s.db.Raw(query, result.AuthorId).Row()
	
	var authorId int
	var authorFullName *string
	err = row.Scan(&authorId, &authorFullName)
	
	if err == nil {
		result.Author.Id = authorId
		if authorFullName != nil {
			result.Author.FullName = *authorFullName
		}
	}
	// Nếu không tìm thấy author, vẫn trả về blog với author empty

	return &result, nil
}

