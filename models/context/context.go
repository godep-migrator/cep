package context

import (
	"github.com/jeffchao/cep/models/user"
)

// Users = [username, user state]
type Context struct {
	UserContext struct {
		Users         [][]string
		BusinessEvent struct {
			Type  string
			Cause string
		}
	}

	EngagementContext struct {
	}
}

func (self *Context) BuildUserContext(flightCode string, event string) error {
	rows, _ := user.FindUsersByFlight(flightCode)

	for _, row := range rows.([][]interface{}) {
		users := []string{}
		for _, v := range row {
			users = append(users, v.(string))
		}
		self.UserContext.Users = append(self.UserContext.Users, users)
	}
	self.UserContext.BusinessEvent.Type = event
	self.UserContext.BusinessEvent.Cause = "mechanical or weather"

	return nil
}

func (self *Context) BuildEngagementContext() error {
	return nil
}
