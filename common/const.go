package common

type AppEnvType string

const (
	AppEnvProd  AppEnvType = "prod"
	AppEnvDev   AppEnvType = "dev"
	AppEnvLocal AppEnvType = "local"
)

const (
	Qualify   string = "qualify"
	UnQualify string = "unqualify"
)
