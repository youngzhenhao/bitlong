package api

import (
	"encoding/json"
	"fmt"
	"github.com/vincent-petithory/dataurl"
	"os"
	"trade/utils"
)

type Meta struct {
	Acronym     string `json:"acronym,omitempty"`
	Description string `json:"description,omitempty"`
	Image_Data  string `json:"image_data,omitempty"`
}

func NewMeta(description string) *Meta {
	meta := Meta{
		Description: description,
	}
	return &meta
}

func (m *Meta) LoadImage(file string) (bool, error) {
	if file != "" {
		image, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("open image file is error:", err)
			return false, err
		}
		imageStr := dataurl.EncodeBytes(image)
		m.Image_Data = imageStr
	}
	return true, nil
}

func (m *Meta) ToJsonStr() string {
	metastr, _ := json.Marshal(m)
	return string(metastr)
}

func (m *Meta) GetMetaFromStr(metaStr string) {
	if metaStr == "" {
		m.Description = "This asset has no meta."
	}
	err := json.Unmarshal([]byte(metaStr), m)
	if err != nil {
		m.Description = metaStr
	}
}

func (m *Meta) FetchAssetMeta(isHash bool, data string) string {
	response, err := fetchAssetMeta(isHash, data)
	if err != nil {
		return utils.MakeJsonResult(false, err.Error(), nil)
	}
	m.GetMetaFromStr(string(response.Data))
	return utils.MakeJsonResult(true, "", nil)
}
