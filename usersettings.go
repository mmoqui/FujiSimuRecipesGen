package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// CameraSettings in which are declared the model of the Fujifilm camera (like X-T4) and the version
// of the camera firmware
type CameraSettings struct {
	Model           string `yaml:"model"`
	FirmwareVersion string `yaml:"firmware"`
}

// XRawFujifilmCameraSettings in which are declared the version and the serial number of the
// post-processing tool embedded in the camera
type XRawFujifilmCameraSettings struct {
	Version      string `yaml:"version"`
	SerialNumber int    `yaml:"camera.serialNb"`
}

// UserSettings in which are declared the user settings from which invariable parameters are used
// to feed the film simulations to generate.
type UserSettings struct {
	Camera     CameraSettings             `yaml:"camera"`
	Xrfc       XRawFujifilmCameraSettings `yaml:"xrfc"`
	FP1Path    string                     `yaml:"fp1Path"`
	ColorSpace string                     `yaml:"colorSpace"`
}

// loadUserSettings loads the user settings from the YAML file located at the specified path.
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
