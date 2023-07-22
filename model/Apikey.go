package model

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Apikey struct {
	gorm.Model
	KeyName   string `gorm:"type:varchar(50);not null" json:"keyName"`
	KeyID     string `gorm:"type:varchar(50);not null" json:"keyId"`
	KeySecret string `gorm:"type:varchar(50);not null" json:"keySecret"`
	Email     string `gorm:"type:varchar(64);not null" json:"email"`
}

func CreateApikey(data *Apikey) int {
	if len(data.KeyID) < 0 {
		return errmsg.ErrorKeyIDEmpty
	}
	if len(data.KeySecret) < 0 {
		return errmsg.ErrorKeySecretEmpty
	}
	if len(data.Email) < 0 {
		return errmsg.ErrorEmailEmpty
	}
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CreateLightsailInstances(region string, input2 lightsail.CreateInstancesInput) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials("AKIAXXFA2ODZL55HHYEW", "cRsyRDLtaiMHmMcHbpxDGBFohXOJyPT+aivsSSQY", ""),
		},
	}))
	svc := lightsail.New(sess)
	input := &lightsail.CreateInstancesInput{
		//InstanceNames: []*string{aws.String("test1"), aws.String("test2")},
		InstanceNames:    aws.StringSlice([]string{"test1", "test2"}),
		BundleId:         aws.String("nano_2_0"),
		BlueprintId:      aws.String("centos_7_2009_01"),
		AvailabilityZone: aws.String(*sess.Config.Region + "a"),
		UserData:         aws.String("#!/bin/bash\nsudo -i\nsudo sed -i 's/^.*PermitRootLogin.*/PermitRootLogin yes/g' /etc/ssh/sshd_config\nsudo sed -i 's/^.*PasswordAuthentication.*/PasswordAuthentication yes/g' /etc/ssh/sshd_config\necho root:9PRZbl9I84Hbo1xH7h | sudo chpasswd\nsystemctl restart sshd"),
	}
	result, err := svc.CreateInstances(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Sprintln(result)
}
