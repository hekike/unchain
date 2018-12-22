package git

import (
	"fmt"
	"strings"

	"github.com/hekike/unchain/pkg/utils"
)

// User git user
type User struct {
	Name  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

var runner utils.Runner = utils.CommandRunner{}

// GetUser returns the git config user.name and user.email
func GetUser(dir string) (*User, error) {
	user := User{}
	out, err := runner.Run(
		dir,
		"git",
		"config",
		"user.name",
	)
	if err != nil {
		return nil, fmt.Errorf("[GetUser] exec name: %v %s", err, string(out))
	}
	user.Name = strings.TrimSpace(string(out))

	out, err = runner.Run(
		dir,
		"git",
		"config",
		"user.email",
	)
	if err != nil {
		return nil, fmt.Errorf("[GetUser] exec email: %v %s", err, string(out))
	}
	user.Email = strings.TrimSpace(string(out))

	return &user, nil
}
