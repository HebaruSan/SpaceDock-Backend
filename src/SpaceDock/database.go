/*
 SpaceDock Backend
 API Backend for the SpaceDock infrastructure to host modfiles for various games

 SpaceDock Backend is licensed under the Terms of the MIT License.
 Copyright (c) 2017 Dorian Stoll (ThomasKerman/TMSP), RockyTV
*/

package SpaceDock

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mssql"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
)

/*
 The database handler that provides the API to interact with the the DB
 */
var Database gorm.DB

/*
 Establishes the connection to the database
 */
func LoadDatabase() {
    db, err := gorm.Open(Settings.Dialect, Settings.ConnectionData)
    if err != nil {
        log.Fatalf("* Failed to connect to the database: %s", err)
    }
    Database = *db
    log.Print("* Database connection successfull")
}

/*
 Creates a table only if it doesn't exist
 */
func CreateTable(models interface{}) {
    if !Database.HasTable(models) {
        Database.CreateTable(models)
    }
}