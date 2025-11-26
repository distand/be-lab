package do

type DeviceType struct {
	ID   uint   `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (m *DeviceType) TableName() string {
	return "lab_device_type"
}
