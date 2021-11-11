package main

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/golang-basics/task02/internal/file"
	"github.com/u-shylianok/golang-basics/task02/internal/model"
)

func main() {
	// get env vars first
	inputPath := os.Getenv("INPUT_PATH")
	outputPath := os.Getenv("OUTPUT_PATH")

	logrus.WithFields(logrus.Fields{
		"input":  inputPath,
		"output": outputPath,
	}).Info("environment variables got")

	// then override from args
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		logrus.WithField("input", inputPath).Info("input overridden from args")

		if len(os.Args) > 2 {
			outputPath = os.Args[2]
			logrus.WithField("output", outputPath).Info("output overridden from args")
		}
	}

	xmlBytes, err := file.ReadBytes(inputPath)
	if err != nil {
		logrus.WithError(err).WithField("path", inputPath).Error("failed to read file")
		return
	}
	logrus.WithField("path", inputPath).Info("file read successfully")

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

	n, err := file.WriteBytes(outputPath, jsonBytes)
	if err != nil {
		logrus.WithError(err).WithField("path", outputPath).Error("failed to write bytes")
		return
	}
	logrus.WithFields(logrus.Fields{
		"path":          outputPath,
		"bytes written": n,
	}).Info("bytes written successfully")

	logrus.Info("exit the program")
}
