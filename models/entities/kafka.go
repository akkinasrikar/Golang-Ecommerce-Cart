package entities

type Consume struct {
	ProcessId   string `gorm:"column:process_id;primaryKey"`
	ProcessName string `gorm:"column:process_name;not null"`
	ProcessData string `gorm:"column:process_data;not null"`
}
