package api

const API_VERSION_PREFFIX = "v0.0.1"
const API_MARKER = "-beta.0."
const API_DATE_TIME = "20240423152645"

var API_VERSION = API_VERSION_PREFFIX + API_MARKER + API_DATE_TIME

func GetApiVersion() string {
	return API_VERSION
}

func NewVersionTag() string {
	return API_VERSION_PREFFIX + API_MARKER + GetTimeSuffixString()
}
