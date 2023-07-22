package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

func CreateLightsailInstances(region string) {
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

func DeleteLightsailInstances(region string) {
	nameList := []string{"test1", "test2"}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials("AKIAXXFA2ODZL55HHYEW", "cRsyRDLtaiMHmMcHbpxDGBFohXOJyPT+aivsSSQY", ""),
		},
	}))
	svc := lightsail.New(sess)

	for _, name := range nameList {
		result, err := svc.DeleteInstance(&lightsail.DeleteInstanceInput{
			ForceDeleteAddOns: aws.Bool(true),
			InstanceName:      aws.String(name),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}

}

func GetLightsailInstance(region string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials("AKIAXXFA2ODZL55HHYEW", "cRsyRDLtaiMHmMcHbpxDGBFohXOJyPT+aivsSSQY", ""),
		},
	}))
	svc := lightsail.New(sess)

	result, err := svc.GetInstance(&lightsail.GetInstanceInput{
		InstanceName: aws.String("test1"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s %s %s %dvCPU%.1fGRAM%dGSSD\n",
		aws.StringValue(result.Instance.BlueprintId),
		aws.StringValue(result.Instance.Name),
		aws.StringValue(result.Instance.PublicIpAddress),
		aws.Int64Value(result.Instance.Hardware.CpuCount),
		aws.Float64Value(result.Instance.Hardware.RamSizeInGb),
		aws.Int64Value(result.Instance.Hardware.Disks[0].SizeInGb))
}

func GetLightsailInstances(region string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials("AKIAXXFA2ODZL55HHYEW", "cRsyRDLtaiMHmMcHbpxDGBFohXOJyPT+aivsSSQY", ""),
		},
	}))
	svc := lightsail.New(sess)

	getInstancesInput := &lightsail.GetInstancesInput{}
	for {
		result, err := svc.GetInstances(getInstancesInput)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, instance := range result.Instances {
			fmt.Printf("%s %s %s %dvCPU%.1fGRAM%dGSSD\n",
				aws.StringValue(instance.BlueprintId),
				aws.StringValue(instance.Name),
				aws.StringValue(instance.PublicIpAddress),
				aws.Int64Value(instance.Hardware.CpuCount),
				aws.Float64Value(instance.Hardware.RamSizeInGb),
				aws.Int64Value(instance.Hardware.Disks[0].SizeInGb))
		}
		if result.NextPageToken == nil {
			break
		}
		getInstancesInput.PageToken = result.NextPageToken
	}
}

func CloseLightsailInstancePublicPorts() {

}

func PutLightsailInstancePublicPorts(region string) {
	nameList := []string{"test1", "test2"}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials("AKIAXXFA2ODZL55HHYEW", "cRsyRDLtaiMHmMcHbpxDGBFohXOJyPT+aivsSSQY", ""),
		},
	}))
	svc := lightsail.New(sess)

	for _, name := range nameList {
		result, err := svc.PutInstancePublicPorts(&lightsail.PutInstancePublicPortsInput{
			InstanceName: aws.String(name),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func main() {
	//CreateLightsailInstances("ap-northeast-2")
	DeleteLightsailInstances("ap-northeast-1")
	//GetLightsailInstance("ap-northeast-2")
	//GetLightsailInstances("ap-northeast-2")
}
