package main

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	alidns "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func queryRecord(full_domain string, r_type string, record *Record) error {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		return err
	}

	request := alidns.CreateDescribeSubDomainRecordsRequest()
	request.Scheme = "https"
	request.SubDomain = full_domain
	request.Type = r_type

	response, err := client.DescribeSubDomainRecords(request)
	if err != nil {
		return err
	}
	rec := response.DomainRecords.Record
	if len(rec) != 0 {
		record.ID = rec[0].RecordId
		record.IP = rec[0].Value
	}
	return nil
}

func addRecord(domain string, rr string, r_type string, value string) error {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"
	request.DomainName = domain
	request.RR = rr
	request.Type = r_type
	request.Value = value

	_, err = client.AddDomainRecord(request)
	if err != nil {
		return err
	}
	return nil
}

func updateRecord(id string, rr string, r_type string, value string) error {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = id
	request.RR = rr
	request.Type = r_type
	request.Value = value

	_, err = client.UpdateDomainRecord(request)
	if err != nil {
		return err
	}
	return nil
}
