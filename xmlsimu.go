package main

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
	"github.com/creasty/defaults"
)

type UserSettings struct {
	CameraModel     string `yaml:"camera"`
	FirmwareVersion string `yaml:"firmware"`
	SerialNumber    int    `yaml:"serialNb"`
	XRawStudioVersion string `yaml:"xrfcVersion"`
}

type FujifilmSimulation struct {
	XMLName       xml.Name              `xml:"ConversionProfile"`
	Application   string                `xml:"application,attr" default:"XRFC"`
	Version       string                `xml:"version,attr" default:"1.12.0.0"`
	PropertyGroup *SimulationProperties `xml:"PropertyGroup"`
}

type SimulationProperties struct {
	Device                 string   `xml:"device,attr"`
	Version                string   `xml:"version,attr"`
	Label                  string   `xml:"label,attr"`
	SerialNumber           int      `xml:"SerialNumber"`
	TetherRAWConditionCode string   `xml:"TetherRAWConditonCode"`
	Editable               string   `xml:"Editable" default:"TRUE"`
	SourceFileName         string   `xml:"SourceFileName"`
	FileError              string   `xml:"Fileerror" default:"NONE"`
	RotationAngle          int8     `xml:"RotationAngle" default:"0"`
	StructVer              int      `xml:"StructVer" default:"65536"`
	IOPCode                string   `xml:"IOPCode" default:"FF129506"`
	ShootingCondition      string   `xml:"ShootingCondition" default:"OFF"`
	FileType               string   `xml:"FileType" default:"JPG"`
	ImageSize              string   `xml:"ImageSize" default:"L3x2"`
	ImageQuality           string   `xml:"ImageQuality" default:"Fine"`
	ExposureBias           string   `xml:"ExposureBias" default:"0"`
	DynamicRange           string   `xml:"DynamicRange"`
	WideDRange             int8     `xml:"WideDRange" default:"0"`
	FilmSimulation         string   `xml:"FilmSimulation"`
	BlackImageTone         int8     `xml:"BlackImageTone" default:"0"`
	MonochromaticColorRG   int8     `xml:"MonochromaticColorRG" default:"0"`
	GrainEffect            string   `xml:"GrainEffect"`
	GrainEffectSize        string   `xml:"GrainEffectSize"`
	ChromeEffect           string   `xml:"ChromeEffect"`
	ColorChromeBlue        string   `xml:"ColorChromeBlue"`
	SmoothSkinEffect       string   `xml:"SmoothSkinEffect" default:"OFF"`
	WBShootCond            string   `xml:"WBShootCond" default:"OFF"`
	WhiteBalance           string   `xml:"WhiteBalance"`
	WBShiftR               int8     `xml:"WBShiftR"`
	WBShiftB               int8     `xml:"WBShiftB"`
	WBColorTemp            string   `xml:"WBColorTemp" default:"10000K"`
	HighlightTone          int8     `xml:"HighlightTone"`
	Color                  int8     `xml:"Color"`
	Sharpness              int8     `xml:"Sharpness"`
	NoiseReduction         int8     `xml:"NoisReduction"`
	Clarity                int8     `xml:"Clarity" default:"0"`
	LensModulationOpt      string   `xml:"LensModulationOpt" default:"ON"`
	ColorSpace             string   `xml:"ColorSpace" default:"sRGB"`
	HDR                    string   `xml:"HDR" default:""`
	DigitalTeleConv        string   `xml:"DigitalTeleConv" default:"OFF"`
}

func loadUserSettings(settingsPath *string) UserSettings {
	settingsFile, err := ioutil.ReadFile(*settingsPath)
	if err != nil {
		log.Fatal(err)
	}
	var settings UserSettings
	yaml.Unmarshal(settingsFile, &settings)
	return settings
}

func generateXMLSimulations(settings *UserSettings, recipes []*FujiSimulationRecipe) {
	var camera = settings.CameraModel + "_" + flatVersion(settings.FirmwareVersion)
	fmt.Printf("Number of simulation to generate: %d\n", len(recipes))

	for _, recipe := range recipes {
		fmt.Printf("Generate simulation %s...",recipe.Label)
		simulation := &FujifilmSimulation{
			Version: settings.XRawStudioVersion,
		}
		err1 := defaults.Set(simulation)
		if err1 != nil {
			log.Fatal(err1)
		}

		properties := &SimulationProperties{
			Device:                 settings.CameraModel,
			Version:                camera,
			Label:                  recipe.Label,
			SerialNumber:           settings.SerialNumber,
			TetherRAWConditionCode: camera,
			DynamicRange:           recipe.DynamicRange,
			FilmSimulation:         recipe.FilmSimulation,
			GrainEffect:            recipe.Grain,
			GrainEffectSize:        computeGrainEffectSize(recipe.Grain),
			ChromeEffect:           recipe.CCFx,
			ColorChromeBlue:        normalizeCCFxB(recipe.CCFx, recipe.CCFxB),
			WhiteBalance:           recipe.WhiteBalance,
			WBShiftR:               recipe.WBShiftR,
			WBShiftB:               recipe.WBShiftB,
			HighlightTone:          recipe.HighlightTone,
			Color:                  recipe.Color,
			Sharpness:              recipe.Sharpness,
			NoiseReduction:         recipe.NoiseReduction,
		}
		err2 := defaults.Set(properties)
		if err2 != nil {
			log.Fatal(err2)
		}

		simulation.PropertyGroup = properties

		xmlSimulation, err := xml.MarshalIndent(simulation, " ", "  ")
		if err != nil {
			log.Fatal(err)
		}

		var fileName = strings.ReplaceAll(strings.ReplaceAll(recipe.Label, "/", "_"), "\\", "_")
		err = ioutil.WriteFile(fileName+".FP1", xmlSimulation, 0644)
		if err != nil {
			fmt.Println(" FAILED")
			log.Fatal(err)
		}
		fmt.Println(" OK")
	}
}

func normalizeCCFxB(ccfx string, ccfxb string) string {
	if ccfx == "OFF" {
		return "OFF"
	}
	return ccfxb
}

func computeGrainEffectSize(grain string) string {
	var size string
	switch grain {
	case "OFF":
	case "WEAK":
	case "STRONG":
		size = "SMALL"
		break
	default:
		size = "LARGE"
	}
	return size
}

func flatVersion(version string) string {
	if strings.Index(version, ".") == 1 {
		return "0" + strings.ReplaceAll(version, ".", "")
	}
	return strings.ReplaceAll(version, ".", "")
}
