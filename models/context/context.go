package context

import (
	"github.com/thresholderio/go-processing/models/user"
	"log"
)

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
	users, _ := user.FindUsersByFlight(flightCode)

	/*
	   for _, user := range users {
	     self.UserContext.Users = append(self.UserContext.Users, reflect.ValueOf(user))
	   }
	*/
	self.UserContext.BusinessEvent.Type = event
	self.UserContext.BusinessEvent.Cause = "mechanical or weather"

	log.Printf("users: %+v\n", users)
	return nil
}
