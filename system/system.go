package system

import (
	"fmt"

	"github.com/denisbakhtin/demoshop/models"
)

//Application mode
const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

var mode string //application mode: debug, release, test

//Init initializes core system elements (DB, sessions, templates, et al)
func Init() {
	loadConfig()
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name)
	models.InitDB(connection)
}

//SetMode sets application mode
func SetMode(flagmode *string) {
	switch *flagmode {
	case ReleaseMode:
		mode = ReleaseMode
	default:
		mode = DebugMode
	}
}

//GetMode returns application mode
func GetMode() string {
	return mode
}
