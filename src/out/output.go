package out

import (
	"fmt"
	"errors"
	"encoding/xml"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

func ToStd(langs []models.Lang, total *models.Lang) error {
	if len(langs) > 0 {
		line, header, dataFormat := utils.GetTemplates(langs)
		stdOut := header
		for _, lang := range langs {
			stdOut += utils.AppendLangData(lang, dataFormat)
		}
		if total != nil {
			stdOut += line + utils.AppendLangData(*total, dataFormat)
		}
		fmt.Println(stdOut)
	} else {
		return errors.New("sloc: there is no any files to count")
	}
	return nil
}

func toFile(langs []models.Lang, total *models.Lang, fileType string) error {
	var err error = nil
	var data []byte
	xmlHeader := ""
	if total != nil {
		langs = append(langs, *total)
	}
	switch fileType {
	case "json":

		data, err = json.MarshalIndent(langs, "", "    ")
	case "xml":
		data, err = xml.MarshalIndent(langs, "", "    ")
		xmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
	case "yml":
		data, err = yaml.Marshal(langs)
	}
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s", fileType, err.Error()))
	}
	return utils.WriteToFile(fmt.Sprintf("sloc_report.%s", fileType), xmlHeader + string(data))
}

func ToJson(langs []models.Lang, total *models.Lang) error {
	return toFile(langs, total, "json")
}

func ToXml(langs []models.Lang, total *models.Lang) error {
	return toFile(langs, total, "xml")
}

func ToYaml(langs []models.Lang, total *models.Lang) error {
	return toFile(langs, total, "yml")
}
