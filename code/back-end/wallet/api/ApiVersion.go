package api

const API_VERSION_PREFFIX = "v0.0.1"
const API_MARKER = "-beta.0."
const API_DATE_TIME = "20240422161454"
const API_VERSION = "v0.0.1-beta.0.20240422161454"

func NewVersionTag() string {
	return API_VERSION_PREFFIX + API_MARKER + GetTimeSuffixString()
}
