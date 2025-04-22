package storage

type User struct {
	ID         uint   `gorm:"primaryKey"`
	TelegramID string `gorm:"uniqueIndex"`
	Username   string
}

type Settings struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	Interval int // через сколько сообщений напоминать
	Enabled  bool
}
