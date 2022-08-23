package controllers

import "github.com/chirag3003/vote-back/database"

var conn *database.Database

type Controllers struct {
}

func Setup(db *database.Database) *Controllers {
	conn = db
	return &Controllers{}
}
