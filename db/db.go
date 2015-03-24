package db

import (
	// "github.com/stevepartridge/go/db/mongo"
	// "github.com/stevepartridge/go/db/mysql"
	"github.com/stevepartridge/go/db/postgres"
	"github.com/stevepartridge/go/db/utils"
)

var Pg = postgres.Create()

// var Mg = mongo.Create()
// var My = mysql.Create()

var Utils = utils.Utils{}
