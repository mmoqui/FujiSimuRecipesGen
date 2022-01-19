package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type FujiSimulationRecipe struct {
	Label          string
	FilmSimulation string
	Grain          string
	CCFx           string
	CCFxB          string
	WhiteBalance   string
	WBShiftR       int8
	WBShiftB       int8
	DynamicRange   string
	HighlightTone  int8
	ShadowTone     int8
	Color          int8
	Sharpness      int8
	NoiseReduction int8
	ExposureBias   string
}

func loadCSV(csvPath *string) []*FujiSimulationRecipe {
	fmt.Printf("Load Fujifilm simulation recipes from CSV %s\n", *csvPath)
	var recipes []*FujiSimulationRecipe

	csvFile, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	// the first tuple should be the header, so we skip it
	var header = true
	for {
		if header {
			header = false
			_, err := csvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			continue
		}
		tuple, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		var ccfx = strings.Split(tuple[3], "/")
		recipes = append(recipes, &FujiSimulationRecipe{
			Label:          tuple[0],
			FilmSimulation: tuple[1],
			Grain:          tuple[2],
			CCFx:           ccfx[0],
			CCFxB:          computeCFXB(ccfx),
			WhiteBalance:   tuple[4],
			WBShiftR:       toInt(tuple[5]),
			WBShiftB:       toInt(tuple[6]),
			DynamicRange:   tuple[7],
			HighlightTone:  toInt(tuple[8]),
			ShadowTone:     toInt(tuple[9]),
			Color:          toInt(tuple[10]),
			Sharpness:      toInt(tuple[11]),
			NoiseReduction: toInt(tuple[12]),
			ExposureBias:   tuple[13],
		})
	}
	fmt.Printf("Number of loaded recipes: %d\n", len(recipes))
	return recipes
}

func computeCFXB(ccfx []string) string {
	if len(ccfx) == 1 {
		return ccfx[0]
	}
	return ccfx[1]
}

func toInt(number string) int8 {
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return int8(n)
}
