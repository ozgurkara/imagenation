package storage

import "time"

type S3Event struct {
	Records []S3EventRecord `json:"Records"`
}

type S3EventRecord struct {
	EventVersion      string              `json:"eventVersion"`
	EventSource       string              `json:"eventSource"`
	AWSRegion         string              `json:"awsRegion"`
	EventTime         time.Time           `json:"eventTime"`
	EventName         string              `json:"eventName"`
	PrincipalID       S3UserIdentity      `json:"userIdentity"`
	RequestParameters S3RequestParameters `json:"requestParameters"`
	ResponseElements  map[string]string   `json:"responseElements"`
	S3                S3Entity            `json:"s3"`
}

type S3UserIdentity struct {
	PrincipalID string `json:"principalId"`
}

type S3RequestParameters struct {
	SourceIPAddress string `json:"sourceIPAddress"`
}

type S3Entity struct {
	SchemaVersion   string   `json:"s3SchemaVersion"`
	ConfigurationID string   `json:"configurationId"`
	Bucket          S3Bucket `json:"bucket"`
	Object          S3Object `json:"object"`
}

type S3Bucket struct {
	Name          string         `json:"name"`
	OwnerIdentity S3UserIdentity `json:"ownerIdentity"`
	Arn           string         `json:"arn"`
}

type S3Object struct {
	Key           string `json:"key"`
	Size          int64  `json:"size"`
	URLDecodedKey string `json:"urlDecodedKey"`
	VersionID     string `json:"versionId"`
	ETag          string `json:"eTag"`
	Sequencer     string `json:"sequencer"`
}
