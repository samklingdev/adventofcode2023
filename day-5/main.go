package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MapValues struct {
	destStart int
	srcStart  int
	rangeLen  int
}

var (
	seeds                 []string
	seedToSoil            []MapValues
	soilToFertilizer      []MapValues
	fertilizerToWater     []MapValues
	waterToLight          []MapValues
	lightToTemperature    []MapValues
	temperatureToHumidity []MapValues
	humidityToLocation    []MapValues
)

func main() {
	generateMaps()
	secondPart()
}

func generateMaps() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	blocks := strings.Split(string(content), "\r\n\r\n")
	for _, block := range blocks {
		switch {
		case strings.HasPrefix(block, "seeds"):
			handleSeeds(block)
		case strings.HasPrefix(block, "seed-to-soil"):
			handleMap(block, &seedToSoil)
		case strings.HasPrefix(block, "soil-to-fertilizer"):
			handleMap(block, &soilToFertilizer)
		case strings.HasPrefix(block, "fertilizer-to-water"):
			handleMap(block, &fertilizerToWater)
		case strings.HasPrefix(block, "water-to-light"):
			handleMap(block, &waterToLight)
		case strings.HasPrefix(block, "light-to-temperature"):
			handleMap(block, &lightToTemperature)
		case strings.HasPrefix(block, "temperature-to-humidity"):
			handleMap(block, &temperatureToHumidity)
		case strings.HasPrefix(block, "humidity-to-location"):
			handleMap(block, &humidityToLocation)
		}
	}
}

func handleSeeds(block string) {
	lines := strings.Split(block, "\r\n")
	for _, line := range lines {
		line = strings.Replace(line, "seeds: ", "", 1)
		seeds = strings.Split(line, " ")
		fmt.Println("line: ", line, seeds)
	}
}

func handleMap(block string, m *[]MapValues) {
	lines := strings.Split(block, "\r\n")
	for i, line := range lines {
		// skip first line
		if i == 0 {
			continue
		}
		var destStart, srcStart, rangeLen int
		if _, err := fmt.Sscanf(line, "%d %d %d", &destStart, &srcStart, &rangeLen); err != nil {
			log.Fatal(err)
		}
		*m = append(*m, MapValues{
			destStart: destStart,
			srcStart:  srcStart,
			rangeLen:  rangeLen,
		})
	}

}

func firstPart() {
	fmt.Println("Result: ")
	var lowestLocation int
	for _, seedStr := range seeds {
		seed, _ := strconv.Atoi(seedStr)
		soil := getMapValue(seed, seedToSoil)
		fertilizer := getMapValue(soil, soilToFertilizer)
		water := getMapValue(fertilizer, fertilizerToWater)
		light := getMapValue(water, waterToLight)
		temperature := getMapValue(light, lightToTemperature)
		humidity := getMapValue(temperature, temperatureToHumidity)
		location := getMapValue(humidity, humidityToLocation)
		// fmt.Println("seed:", seed, "soil:", soil, "fertilizer:", fertilizer, "water:", water, "light:", light, "temperature:", temperature, "humidity:", humidity, "location:", location)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println("lowest location: ", lowestLocation)
}

func getMapValue(val int, m []MapValues) int {
	for _, m := range m {
		if val >= m.srcStart && val < m.srcStart+m.rangeLen {
			return m.destStart + (val - m.srcStart)
		}
	}
	return val
}

func secondPart() {
	fmt.Println("Result: ")
	var lowestLocation int
	for i := 0; i < len(seeds); i += 2 {
		seedStr := seeds[i]
		rangeStr := seeds[i+1]
		seed, _ := strconv.Atoi(seedStr)
		r, _ := strconv.Atoi(rangeStr)
		fmt.Printf("range: %d\n", r)
		p := 0
		oneP := r / 100
		for j := seed; j < seed+r; j++ {
			if j-seed > p*oneP {
				p++
				fmt.Printf("p: %d\n", p)
			}
			soil := getMapValue(j, seedToSoil)
			fertilizer := getMapValue(soil, soilToFertilizer)
			water := getMapValue(fertilizer, fertilizerToWater)
			light := getMapValue(water, waterToLight)
			temperature := getMapValue(light, lightToTemperature)
			humidity := getMapValue(temperature, temperatureToHumidity)
			location := getMapValue(humidity, humidityToLocation)
			// fmt.Println("seed:", seed, "soil:", soil, "fertilizer:", fertilizer, "water:", water, "light:", light, "temperature:", temperature, "humidity:", humidity, "location:", location)
			if lowestLocation == 0 || location < lowestLocation {
				lowestLocation = location
			}
		}
	}
	fmt.Println("lowest location: ", lowestLocation)
}
