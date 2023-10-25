package models

type MerlinAWS struct {
	Id             string `json:"id"`
	ConnectionName string `json:"connectionName"`
	SecretKey      string `json:"secretKey"`
	SecretId       string `json:"secretId"`
}
