package model

type BasicInfo struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`  // 創建人
	ModifiedBy string `json:"modified_by"` // 修改人
	CreatedOn  uint32 `json:"created_on"`  // 創建時間
	ModifiedOn uint32 `json:"modified_on"` // 修改時間
	DeletedOn  uint32 `json:"deleted_on"`  // 刪除時間
	IsDel      uint8  `json:"is_del"`      // 是否刪除 0為未刪除、1為已刪除
}
