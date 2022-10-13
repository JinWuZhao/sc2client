package sc2client

import (
	"encoding/xml"
	"fmt"
	"os"
)

const (
	BankValueTypeString = "string"
	BankValueTypeFixed  = "fixed"
	BankValueTypeInt    = "int"
	BankValueTypeBool   = "bool"
)

type BankValue struct {
	Type  string
	Value string
}

type XMLValue struct {
	XMLName xml.Name `xml:"Value"`
	Value   xml.Attr `xml:",any,attr"`
}

type XMLKey struct {
	XMLName xml.Name `xml:"Key"`
	Name    string   `xml:"name,attr"`
	Value   XMLValue `xml:"Value"`
}

type XMLSection struct {
	XMLName   xml.Name       `xml:"Section"`
	Name      string         `xml:"name,attr"`
	Keys      []XMLKey       `xml:"Key,omitempty"`
	IndexData map[string]int `xml:"-"`
}

type XMLBank struct {
	XMLName   xml.Name       `xml:"Bank"`
	Version   string         `xml:"version,attr"`
	Sections  []XMLSection   `xml:"Section,omitempty"`
	IndexData map[string]int `xml:"-"`
}

type Bank struct {
	path string
	data XMLBank
}

func NewBank(name string) (*Bank, error) {
	bankFilePath, err := GetSC2BankPath(name)
	if err != nil {
		return nil, fmt.Errorf("GetSC2BankPath() error: %w", err)
	}
	return &Bank{
		path: bankFilePath,
	}, nil
}

func (b *Bank) Load() error {
	bankFile, err := os.ReadFile(b.path)
	if err != nil {
		return fmt.Errorf(" os.ReadFile() error: %w", err)
	}
	err = xml.Unmarshal(bankFile, &b.data)
	if err != nil {
		return fmt.Errorf("xml.Unmarshal() error: %w", err)
	}
	b.data.IndexData = map[string]int{}
	for i, s := range b.data.Sections {
		b.data.IndexData[s.Name] = i
		s.IndexData = map[string]int{}
		for j, k := range s.Keys {
			s.IndexData[k.Name] = j
		}
		b.data.Sections[i] = s
	}
	return nil
}

func (b *Bank) Save() error {
	content, err := xml.MarshalIndent(b.data, "", "    ")
	if err != nil {
		return fmt.Errorf("xml.Marshal() error: %w", err)
	}
	err = os.WriteFile(b.path, append([]byte(xml.Header), content...), 0644)
	if err != nil {
		return fmt.Errorf("os.WriteFile() error: %w", err)
	}
	return nil
}

func (b *Bank) StoreValue(section string, key string, value BankValue) {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		b.data.Sections = append(b.data.Sections, XMLSection{
			Name:      section,
			IndexData: map[string]int{},
		})
		sectionIndex = len(b.data.Sections) - 1
		b.data.IndexData[section] = sectionIndex
	}
	sectionData := b.data.Sections[sectionIndex]
	keyIndex, ok := sectionData.IndexData[key]
	if !ok {
		sectionData.Keys = append(sectionData.Keys, XMLKey{
			Name: key,
		})
		keyIndex = len(sectionData.Keys) - 1
		sectionData.IndexData[key] = keyIndex
	}
	sectionData.Keys[keyIndex].Value.Value = xml.Attr{
		Name: xml.Name{
			Local: value.Type,
		},
		Value: value.Value,
	}
	b.data.Sections[sectionIndex] = sectionData
}

func (b *Bank) LoadValue(section string, key string) (BankValue, bool) {
	var result BankValue
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		return result, false
	}
	sectionData := b.data.Sections[sectionIndex]
	keyIndex, ok := sectionData.IndexData[key]
	if !ok {
		return result, false
	}
	valueData := sectionData.Keys[keyIndex].Value
	result.Type = valueData.Value.Name.Local
	result.Value = valueData.Value.Value
	return result, true
}
