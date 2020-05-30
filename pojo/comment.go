package pojo

type Comment struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Data       string `gorm:"size:255;not null;" json:"data"`
	SharedUser User   `json:"sharedUser"`
}
