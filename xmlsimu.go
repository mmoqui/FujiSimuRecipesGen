package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/creasty/defaults"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FujifilmSimulation struct {
	XMLName       xml.Name              `xml:"ConversionProfile"`
	Application   string                `xml:"application,attr" default:"XRFC"`
	Version       string                `xml:"version,attr" default:"1.12.0.0"`
	PropertyGroup *SimulationProperties `xml:"PropertyGroup"`
}

type SimulationProperties struct {
	Device                 string `xml:"device,attr"`
	Version                string `xml:"version,attr"`
	Label                  string `xml:"label,attr"`
	SerialNumber           int    `xml:"CameraSerialNumber"`
	TetherRAWConditionCode string `xml:"TetherRAWConditonCode"`
	Editable               string `xml:"Editable" default:"TRUE"`
	SourceFileName         string `xml:"SourceFileName"`
	FileError              string `xml:"Fileerror" default:"NONE"`
	RotationAngle          int8   `xml:"RotationAngle" default:"0"`
	StructVer              int    `xml:"StructVer" default:"65536"`
	IOPCode                string `xml:"IOPCode" default:"FF159505"`
	ShootingCondition      string `xml:"ShootingCondition" default:"OFF"`
	FileType               string `xml:"FileType" default:"JPG"`
	ImageSize              string `xml:"ImageSize" default:"L3x2"`
	ImageQuality           string `xml:"ImageQuality" default:"Fine"`
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
	HDR                    string   `xml:"HDR"`
	DigitalTeleConv        string   `xml:"DigitalTeleConv" default:"OFF"`
}

func generateXMLSimulations(settings *UserSettings, recipes []*FujiSimulationRecipe) {
	var camera = settings.Camera.Model + "_" + flatVersion(settings.Camera.FirmwareVersion)
	fmt.Printf("The film simulations will be generated into the folder '%s'\n", settings.XRawStudio.FP1Path)
	fmt.Printf("Number of film simulation to generate: %d\n", len(recipes))

	createRecursivelyDirectory(settings.XRawStudio.FP1Path)

	for _, recipe := range recipes {
		fmt.Printf("Generate film simulation %s...", recipe.Label)
		simulation := &FujifilmSimulation{
			Version: settings.XRawStudio.Version,
		}
		err1 := defaults.Set(simulation)
		if err1 != nil {
			log.Fatal(err1)
		}

		properties := &SimulationProperties{
			Device:                 settings.Camera.Model,
			Version:                camera,
			Label:                  recipe.Label,
			SerialNumber:           settings.Camera.SerialNumber,
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
		xmlSimulation = []byte(xml.Header + string(xmlSimulation))
		if err != nil {
			log.Fatal(err)
		}

		var fp1Name = normalizeFileName(recipe.Label)
		var fp1Path = filepath.Join(settings.XRawStudio.FP1Path, fp1Name+".FP1")
		err = ioutil.WriteFile(fp1Path, xmlSimulation, 0644)
		if err != nil {
			fmt.Println(" FAILED")
			log.Fatal(err)
		}
		fmt.Println(" OK")
	}
}

func normalizeFileName(fileName string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(fileName, "/", "_"),
			"\\", "_"), // for Unix system like MacOS X
		"*", "_")
}

func createRecursivelyDirectory(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
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
