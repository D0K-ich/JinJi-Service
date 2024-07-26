package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	TableUsers = "users"

	StateUnconfirmed 		= "unconfirmed"
	StateBlocked     		= "blocked"
	StateActive      		= "active"
)

type PrimaryId struct {
	Id int `gorm:"column:id;primary_key:true" json:"id"`
}

type Users []*User
type User struct {
	PrimaryId

	Avatar				[]byte					`gorm:"column:avatar"               json:"avatar"`

	Name              	string 					`gorm:"column:name"                 json:"name"`
	Email            	string          		`gorm:"column:email"                json:"email"`
	State            	string          		`gorm:"column:state"                json:"state"`

	Level 				int						`gorm:"column:level"                 json:"level"`

	TariffId         	int             		`gorm:"column:tariff_id"            json:"tariff_id"`

	Balance          	float64         		`gorm:"column:balance"              json:"balance"`

	CreatedAt        	time.Time       		`gorm:"column:created_at"           json:"created_at"`

	UpdatedAt        	time.Time       		`gorm:"column:updated_at"           json:"updated_at"`
	LastOnline			time.Time
	TariffExpiration 	*time.Time      		`gorm:"columng:tariff_expiration"   json:"tariff_expiration"`

	//Services            *manager.UserSettings   `gorm:"column:settings"             json:"settings"`
	//Achievement            *manager.UserSettings   `gorm:"column:settings"             json:"settings"`
	Friends				Friends					`gorm:"column:friends"`
}

func(u *User) TableName()    string  {return TableUsers}
func(u *User) IsEmpty()      bool    {return u == nil || u.PrimaryId.Id == 0}
func(u *User) IsBanned()     bool    {return u.State == StateBlocked}

type Friends struct {
	Friends []*Friend `json:"friends"`
}

type Friend struct {
	Name 				string		`json:"name"`
	LastOnline			time.Time	`json:"last_online"`
	Level 				int			`json:"level"`
}

func(f *Friends) Scan(data interface{}) error {return json.Unmarshal(data.([]byte), f)}
func(f *Friends) Value() (driver.Value, error) {
	if f == nil {return nil, nil}
	return json.Marshal(f)
}
