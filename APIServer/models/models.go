/*
Package model provides the base of the project, which are the
working structures
*/
package models

//User details
type User struct {
	ID       uint32      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Username string      `json:"username" gorm:"unique;not null"`
	Password string      `json:"password" gorm:"not null"`
	Phone    string      `json:"phone" gorm:"not null"`
	Email    string      `json:"email" gorm:"not null"`
	Inv      []Inv       `json:"inv" gorm:"-"`
	Setting  UserSetting `json:"setting" gorm:"-"`

	Updated uint64 `json:"updated" gorm:"autoUpdateTime;not null"`
	Created uint64 `json:"created" gorm:"autoCreateTime;not null"`
}

//User settings, UserID->User.ID
type UserSetting struct {
	ID          uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserID      uint32 `json:"userid" gorm:"not null"`
	Type        uint32 `json:"-" gorm:"not null"`
	NotiTime    string `json:"notitime" gorm:"not null"`
	NotiSetting uint32 `json:"notisetting" gorm:"not null"` //1,2
	PrevLogin   string `json:"prevlogin" gorm:"not null"`
	CurrLogin   string `json:"currlogin" gorm:"not null"`

	Updated uint64 `json:"updated" gorm:"autoUpdateTime;not null"`
	Created uint64 `json:"created" gorm:"autoCreateTime;not null"`
}

//Inventory details, UserID->User.ID
type Inv struct {
	ID     uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserID uint32 `json:"userid" gorm:"not null"`
	Icon   string `json:"icon" gorm:"not null"`
	Name   string `json:"name" gorm:"not null"`
	Items  []Item `json:"items" gorm:"-"`

	Updated uint64 `json:"updated" gorm:"autoUpdateTime;not null"`
	Created uint64 `json:"created" gorm:"autoCreateTime;not null"`
}

//Item details, InvID->Inv.ID
type Item struct {
	ID       uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	InvID    uint32 `json:"invid" gorm:"not null"`
	Icon     string `json:"icon" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	Storage  string `json:"storage" gorm:"not null"`
	Expiry   uint64 `json:"expiry" gorm:"not null"` //date of expiry, unix
	Idle     uint64 `json:"idle" gorm:"not null"`   //date the item is idle
	Notify   uint32 `json:"notify" gorm:"not null"` //days before to notify for expiry

	Updated uint64 `json:"updated" gorm:"autoUpdateTime;not null"`
	Created uint64 `json:"created" gorm:"autoCreateTime;not null"`
}

//Category details, UserID->User.ID
type Category struct {
	ID     uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserID uint32 `json:"userid" gorm:"not null"`
	Name   string `json:"name" gorm:"not null"`

	Updated uint64 `json:"updated" gorm:"autoUpdateTime;not null"`
	Created uint64 `json:"created" gorm:"autoCreateTime;not null"`
}

//Other response
type OtherRes struct {
	Msg string `json:"msg"`
}
