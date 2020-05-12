package function

import (
	"fmt"
	"socket/server/model"
)

func Register()  *model.User{
	u := &model.User{}
	fmt.Printf("Id:")
	fmt.Scanf("%d\n", &u.Id)
	fmt.Printf("Pwd:")
	fmt.Scanf("%s", &u.Pwd)
	return u
}
