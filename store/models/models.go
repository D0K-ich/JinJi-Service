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

type Dialog struct {
	Uuid 		string 		`gorm:"column:uuid"                 json:"uuid"`

	Header 		string		`gorm:"column:header"               json:"header"`
	SubHeader 	string		`gorm:"column:sub_header"           json:"sub_header"`
	DateUpdated time.Time	`gorm:"column:date_updated"         json:"date_updated"`
}

func(u *Dialog) TableName()    string  {return TableUsers}
func(u *Dialog) IsEmpty()      bool    {return u == nil || u.Uuid == ""}

type Users []*User
type User struct {
	//PrimaryId

	Uuid	 			string					`gorm:"column:uuid"                 json:"uuid"`
	//Avatar				[]byte					`gorm:"column:avatar"               json:"avatar"`

	Name              	string 					`gorm:"primaryKey;<-:create,column:name"                 json:"name"`
	Email            	string          		`gorm:"column:email"                json:"email"`
	State            	string          		`gorm:"column:state"                json:"state"`
	Phone 				string					`gorm:"column:phone"                json:"phone"`
	Password  			string					`gorm:"column:password"             json:"password"`

	Level 				*Level					`gorm:"column:level"                 json:"level"`

	TariffId         	int             		`gorm:"column:tariff_id"            json:"tariff_id"`

	Balance          	float64         		`gorm:"column:balance"              json:"balance"`

	CreatedAt        	time.Time       		`gorm:"column:created_at"           json:"created_at"`

	UpdatedAt        	time.Time       		`gorm:"column:updated_at"           json:"updated_at"`
	LastOnline			time.Time				`gorm:"column:last_online"          json:"last_online"`
	TariffExpiration 	*time.Time      		`gorm:"column:tariff_expiration"   json:"tariff_expiration"`

	//Services            *manager.UserSettings   `gorm:"column:settings"             json:"settings"`
	Achievements        *Achievements   		`gorm:"column:achievements"             json:"achievements"`
	Friends				*Friends				`gorm:"column:friends"              json:"friends"`
	Dialogs				*Dialogs				`gorm:"column:dialogs"              json:"dialogs"`
}

func(u *User) TableName()    string  {return TableUsers}
func(u *User) IsEmpty()      bool    {return u == nil || u.Name == ""}
func(u *User) IsBanned()     bool    {return u.State == StateBlocked}

type Level struct {
	Name 			string	`json:"name"`
	PointsCurrent 	int		`json:"points_current"`
	PointsTotal 	int		`json:"points_total"`
}

func(l *Level) Scan(data interface{}) error {return json.Unmarshal(data.([]byte), l)}
func(l *Level) Value() (driver.Value, error) {
	if l == nil {return nil, nil}
	return json.Marshal(l)
}

type Achievements struct {
	Achievements []*Achievement 	`json:"achievements"`
}
type Achievement struct {
	Name 	string
	DateGet time.Time
}

func(a *Achievements) Scan(data interface{}) error {return json.Unmarshal(data.([]byte), a)}
func(a *Achievements) Value() (driver.Value, error) {
	if a == nil {return nil, nil}
	return json.Marshal(a)
}

type Friends struct {
	Friends []*Friend `gorm:"friends" json:"friends"`
}

type Friend struct {
	Name 				string		`json:"name"`
	LastOnline			time.Time	`json:"last_online"`
	Level 				*Level		`json:"level"`
}

func(f *Friends) Scan(data interface{}) error {return json.Unmarshal(data.([]byte), f)}
func(f *Friends) Value() (driver.Value, error) {
	if f == nil {return nil, nil}
	return json.Marshal(f)
}

type Dialogs struct {
	Dialogs []*Dialog `gorm:"dialogs" json:"dialogs"`
}

func(d *Dialogs) Scan(data interface{}) error {return json.Unmarshal(data.([]byte), d)}
func(d *Dialogs) Value() (driver.Value, error) {
	if d == nil {return nil, nil}
	return json.Marshal(d)
}

