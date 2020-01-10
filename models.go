package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Round a round, with start/end and results
type Round struct {
	ID                 uint `gorm:"primary_key"`
	InitializeDatetime *time.Time
	StartDatetime      *time.Time
	ShutdownDatetime   *time.Time
	EndDatetime        *time.Time
	ServerIP           int
	ServerPort         uint
	CommitHash         string
	GameMode           string
	GameModeResult     string
	EndState           string
	MapName            string
}

// ConnectionLog a record of a players that connected, can be tied back to a round via RoundId
type ConnectionLog struct {
	ID         uint `gorm:"primary_key"`
	Datetime   *time.Time
	ServerIP   uint
	ServerPort uint
	RoundID    uint
	Ckey       string
	ip         uint
	computerid string
}

// Player a typical player record
type Player struct {
	Ckey             string `gorm:"primary_key"`
	ByondKey         string
	Firstseen        *time.Time
	FirstseenRoundID uint
	Lastseen         *time.Time
	LastseenRoundID  uint
	ip               uint
	computerid       string
	Lastadminrank    string
	Accountjoindate  *time.Time
	flags            uint

	ConnectionLogs []ConnectionLog `json:"-" gorm:"foreignkey:Ckey"`
}
