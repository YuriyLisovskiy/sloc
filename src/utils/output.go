package utils

import (
	"fmt"
	"errors"
	"encoding/xml"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func OutputToStd(langs []parser.Lang, total parser.Lang) error {
	if len(langs) > 0 {
		line, header, dataFormat := GetTemplates(langs)
		stdOut := header
		for _, lang := range langs {
			stdOut += AppendLangData(lang, dataFormat)
		}
		stdOut += line + AppendLangData(total, dataFormat)
		fmt.Println(stdOut)
	} else {
		return errors.New("sloc: there is no any files to count")
	}
	return nil
}

func OutputToJson(langs []parser.Lang, total parser.Lang) error {
	data, err := json.MarshalIndent(append(langs, total), "", "    ")
	if err != nil {
		return errors.New(fmt.Sprintf("json: %s", err.Error()))
	}
	err = createDir(*args.OutPathPtr)
	if err != nil {
		return errors.New(fmt.Sprintf("json: %s", err.Error()))
	}
	return writeToFile(*args.OutPathPtr + args.DefaultOutFileName + ".json", string(data))
}

func OutputToXml(langs []parser.Lang, total parser.Lang) error {
	data, err := xml.MarshalIndent(append(langs, total), "", "    ")
	if err != nil {
		return errors.New(fmt.Sprintf("xml: %s", err.Error()))
	}
	err = createDir(*args.OutPathPtr)
	if err != nil {
		return errors.New(fmt.Sprintf("xml: %s", err.Error()))
	}
	return writeToFile(
		*args.OutPathPtr + args.DefaultOutFileName + ".xml",
		"<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n" + string(data),
	)
}

func OutputToYaml(langs []parser.Lang, total parser.Lang) error {
	data, err := yaml.Marshal(append(langs, total))
	if err != nil {
		return errors.New(fmt.Sprintf("yaml: %s", err.Error()))
	}
	err = createDir(*args.OutPathPtr)
	if err != nil {
		return errors.New(fmt.Sprintf("yaml: %s", err.Error()))
	}
	return writeToFile(*args.OutPathPtr + args.DefaultOutFileName + ".yml", string(data))
}
