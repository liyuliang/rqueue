package system

import "net/url"

func Init(redisUri string) {


	U, _ := url.Parse(redisUri)
	pwd,_ := U.User.Password()


}
