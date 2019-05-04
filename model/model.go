package model

type Base struct {
	Id        `gorm:"primary_key" json:"id"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	Deleted   int       `json:"deleted"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
