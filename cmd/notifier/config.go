package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"os"
)

type config struct {
	Mail struct {
		TemplatesFolderPath string `conf:"help:Path to mail-templates folder,default:/mail-templates"`
		SMTPUsername        string `conf:"env:MAIL_SMTP_USERNAME,help:SMTP username to authorize with,required"`
		SMTPPassword        string `conf:"env:MAIL_SMTP_PASSWORD,help:SMTP password to authorize with,required,noprint"`
		SMTPHost            string `conf:"env:MAIL_SMTP_HOST,help:SMTP host to connect to,required"`
		SMTPPort            int    `conf:"env:MAIL_SMTP_PORT,help:SMTP port to connect to,default:587"`
		TLS                 struct {
			InsecureSkipVerify bool   `conf:"help:true if certificates should not be verified,default:false"`
			ServerName         string `conf:"help:name of the server who expose the certificate"`
		}
	}
}

func newConfig() (config, error) {
	cfg := config{}

	if origErr := conf.Parse(os.Environ(), "SJP", &cfg); origErr != nil {
		usage, err := conf.Usage("SJP", &cfg)
		if err != nil {
			return cfg, err
		}
		fmt.Println(usage)
		return cfg, origErr
	}

	return cfg, nil
}
