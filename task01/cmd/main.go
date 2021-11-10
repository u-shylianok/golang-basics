package main

import (
	"encoding/json"
	"encoding/xml"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/golang-basics/task01/internal/file"
	"github.com/u-shylianok/golang-basics/task01/internal/model"
)

const (
	xmlPath  = "psy-test-config.xml"
	jsonPath = "psy-test-config.json"
)

func main() {
	xmlBytes, err := file.ReadBytes(xmlPath)
	if err != nil {
		logrus.WithError(err).WithField("path", xmlPath).Error("failed to read file")
		return
	}
	logrus.WithField("path", xmlPath).Info("file read successfully")

	var config model.Config
	if err := xml.Unmarshal(xmlBytes, &config); err != nil {
		logrus.WithError(err).Error("failed to unmarshal config from xml")
		return
	}
	logrus.Info("file unmarshaled from xml successfully") // without print config

	// jsonBytes, err := json.Marshal(&config) // without pretty print
	jsonBytes, err := json.MarshalIndent(&config, "", "\t") // just for pretty print
	if err != nil {
		logrus.WithError(err).Error("failed to marshal config to json")
		return
	}
	logrus.Info("file marshaled to json successfully") // without print json bytes

	n, err := file.WriteBytes(jsonPath, jsonBytes)
	if err != nil {
		logrus.WithError(err).WithField("path", jsonPath).Error("failed to write bytes")
		return
	}
	logrus.WithFields(logrus.Fields{
		"path":          jsonPath,
		"bytes written": n,
	}).Info("bytes written successfully")

	logrus.Info("exit the program")
}
