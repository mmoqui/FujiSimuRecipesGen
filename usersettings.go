package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type CameraSettings struct {
	Model           string `yaml:"model"`
	SerialNumber    int    `yaml:"camera.serialNb"`
	FirmwareVersion string `yaml:"firmware"`
}

type XRawStudioSettings struct {
	Version string `yaml:"version"`
	FP1Path string `yaml:"fp1Path"`
}

type UserSettings struct {
	Camera     CameraSettings     `yaml:"camera"`
	XRawStudio XRawStudioSettings `yaml:"xrfc"`
	ColorSpace string             `yaml:"colorSpace"`
}

func loadUserSettings(settingsPath *string) UserSettings {
	settingsFile, err := ioutil.ReadFile(*settingsPath)
	if err != nil {
		log.Fatal(err)
	}
	var settings UserSettings
	err = yaml.Unmarshal(settingsFile, &settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}
