package fn

import (
	"github.com/koltyakov/gosip"
	strategy "github.com/koltyakov/gosip/auth/addin"
)

func Init(siteurl string, clientid string, clientsecret string) (interface{}, error) {
	// initiate authenticate object
	auth := &strategy.AuthCnfg{
		SiteURL:      siteurl,
		ClientID:     clientid,
		ClientSecret: clientsecret,
	}

	// bind auth client with fluent API
	client := &gosip.SPClient{AuthCnfg: auth}

	return client, nil
}
