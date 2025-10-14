// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"

// 	"golang.org/x/net/html/charset"
// 	"gopkg.in/yaml.v3"
// )

// type Config struct {
// 	InputFile  string `yaml:"input-file"`
// 	OutputFile string `yaml:"output-file"`
// }

// // type Valute struct {
// // 	NumCode  string `xml:"NumCode"`
// // 	CharCode string `xml:"CharCode"`
// // 	Value    string `xml:"Value"`
// // }

// type ValuteCurs struct {
// 	Valutes []Valute `xml:"Valute"`
// }

// type Valute struct {
// 	NumCode  int     `json:"NumCode"  xml:"NumCode"`
// 	CharCode string  `json:"CharCode" xml:"CharCode"`
// 	Value    float64 `json:"Value" xml:"Value"`
// }

// func LoadConfigYaml() (Config, error) {
// 	configPath := flag.String("config", "", "path of config file")
// 	flag.Parse()

// 	dataConfig, err := os.ReadFile(*configPath)
// 	if err != nil {
// 		return Config{}, fmt.Errorf("error read config: %w", err)
// 	}

// 	var config Config

// 	err = yaml.Unmarshal(dataConfig, &config)
// 	if err != nil {
// 		return Config{}, fmt.Errorf("error parsing yaml: %w", err)
// 	}

// 	return config, nil
// }

// type Value float64

// func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	var s string
// 	if err := d.DecodeElement(&s, &start); err != nil {
// 		return err
// 	}

// 	s = strings.TrimSpace(s)
// 	s = strings.Replace(s, ",", ".", -1)

// 	val, err := strconv.ParseFloat(s, 64)
// 	if err != nil {
// 		return fmt.Errorf("parse Value '%s': %w", s, err)
// 	}

// 	*v = Value(val)
// 	return nil
// }

// type xmlValute struct {
// 	NumCode  int    `xml:"NumCode"`
// 	CharCode string `xml:"CharCode"`
// 	Value    Value  `xml:"Value"`
// }

// func (v *Valute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	var xv xmlValute
// 	if err := d.DecodeElement(&xv, &start); err != nil {
// 		return err
// 	}

// 	v.NumCode = int(xv.NumCode)
// 	v.CharCode = xv.CharCode
// 	v.Value = float64(xv.Value)
// 	return nil
// }

// func ParseValuteCursXML(inputFile string) (ValuteCurs, error) {
// 	var valCurs ValuteCurs

// 	file, err := os.Open(inputFile)
// 	if err != nil {
// 		return valCurs, fmt.Errorf("open input file: %w", err)
// 	}

// 	defer func() {
// 		file.Close()
// 	}()

// 	decoder := xml.NewDecoder(file)
// 	decoder.CharsetReader = charset.NewReaderLabel

// 	err = decoder.Decode(&valCurs)
// 	if err != nil {
// 		return valCurs, fmt.Errorf("error parsing XML: %w", err)
// 	}

// 	return valCurs, nil
// }

// // func CreateValuteCursJSON(valCurs ValuteCurs) ([]byte, error) {
// // 	var valutesOutput []Valute

// // 	for _, valute := range valCurs.Valutes {
// // 		numCode, _ := strconv.Atoi(valute.NumCode)
// // 		value, _ := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
// // 		valutesOutput = append(valutesOutput, Valute{
// // 			NumCode:  numCode,
// // 			CharCode: valute.CharCode,
// // 			Value:    value,
// // 		})
// // 	}

// // 	sort.Slice(valutesOutput, func(i, j int) bool {
// // 		return valutesOutput[i].Value > valutesOutput[j].Value
// // 	})

// // 	outputJSON, err := json.MarshalIndent(valutesOutput, "", "    ")
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error marshaling JSON: %w", err)
// // 	}

// // 	return outputJSON, nil
// // }

// // func CreateValuteCursJSON(valCurs ValuteCurs) ([]byte, error) {
// // 	sort.Slice(valCurs.Valutes, func(i, j int) bool {
// // 		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
// // 	})

// // 	outputJSON, err := json.MarshalIndent(valCurs.Valutes, "", "    ")
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error marshaling JSON: %w", err)
// // 	}

// // 	return outputJSON, nil
// // }

// func WriteFileJSON(outputFile string, value any) error {
// 	err := os.MkdirAll(filepath.Dir(outputFile), 0766)
// 	if err != nil {
// 		return fmt.Errorf("error creating directory: %w", err)
// 	}

// 	file, err := os.Create(outputFile)
// 	if err != nil {
// 		return fmt.Errorf("error open output file: %w", err)
// 	}

// 	defer func() {
// 		file.Close()
// 	}()

// 	dataJSON := json.NewEncoder(file)
// 	err = dataJSON.Encode(&value)
// 	if err != nil {
// 		return fmt.Errorf("erorr json encode  %w", err)
// 	}

// 	return nil
// }

// // func WriteFileJSON(outputFile string, outputJSON []byte) error {
// // 	err := os.MkdirAll(filepath.Dir(outputFile), 0755)
// // 	if err != nil {
// // 		return fmt.Errorf("error creating directory: %w", err)
// // 	}

// // 	err = os.WriteFile(outputFile, outputJSON, 0644)
// // 	if err != nil {
// // 		return fmt.Errorf("error writing output file: %w", err)
// // 	}

// // 	return nil
// // }

// func main() {
// 	config, err := LoadConfigYaml()
// 	if err != nil {
// 		panic(err)
// 	}

// 	valCurs, err := ParseValuteCursXML(config.InputFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// outputJSON, err := CreateValuteCursJSON(valCurs)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	err = WriteFileJSON(config.OutputFile, valCurs.Valutes)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type xmlValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlValCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

func LoadConfigYaml() (Config, error) {
	configPath := flag.String("config", "", "path of config file")
	flag.Parse()

	if *configPath == "" {
		return Config{}, fmt.Errorf("config path is required")
	}

	dataConfig, err := os.ReadFile(*configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error read config: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(dataConfig, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing yaml: %w", err)
	}

	return config, nil
}

func ParseValuteCursXML(inputFile string) ([]Valute, error) {
	var xmlData xmlValCurs

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&xmlData)
	if err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}

	var valutes []Valute
	for _, xmlVal := range xmlData.Valutes {
		numCode, err := strconv.Atoi(strings.TrimSpace(xmlVal.NumCode))
		if err != nil {
			return nil, fmt.Errorf("error parsing NumCode '%s': %w", xmlVal.NumCode, err)
		}

		cleanValue := strings.TrimSpace(xmlVal.Value)
		cleanValue = strings.Replace(cleanValue, ",", ".", -1)
		value, err := strconv.ParseFloat(cleanValue, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Value '%s': %w", xmlVal.Value, err)
		}

		valutes = append(valutes, Valute{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(xmlVal.CharCode),
			Value:    value,
		})
	}

	return valutes, nil
}

func CreateValuteCursJSON(valutes []Valute) ([]byte, error) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return outputJSON, nil
}

func WriteFileJSON(outputFile string, outputJSON []byte) error {
	err := os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

func main() {
	config, err := LoadConfigYaml()
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	valutes, err := ParseValuteCursXML(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error parsing XML: %v", err))
	}

	outputJSON, err := CreateValuteCursJSON(valutes)
	if err != nil {
		panic(fmt.Sprintf("Error creating JSON: %v", err))
	}

	err = WriteFileJSON(config.OutputFile, outputJSON)
	if err != nil {
		panic(fmt.Sprintf("Error writing file: %v", err))
	}
}
