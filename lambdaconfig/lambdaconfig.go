package lambdaconfig

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetConfig() *session.Session {

	sess, err := session.NewSession(&aws.Config{})

	if err != nil {
		fmt.Print("error creating sesion please see : %s", err)
	}

	return sess

}
