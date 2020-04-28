package scrap

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func PrintFollowers(insta *goinsta.Instagram) {
	followers := insta.Account.Followers()
	for followers.Next() {
		for i := 0; i < len(followers.Users); i++ {
			fmt.Println(followers.Users[i].Username)
		}
	}
}
