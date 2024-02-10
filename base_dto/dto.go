package base_dto

type DTO struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
