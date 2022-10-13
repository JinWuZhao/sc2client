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
	Keys      []*XMLKey      `xml:"Key,omitempty"`
	IndexData map[string]int `xml:"-"`
}

type XMLBank struct {
	XMLName   xml.Name       `xml:"Bank"`
	Version   string         `xml:"version,attr"`
	Sections  []*XMLSection  `xml:"Section,omitempty"`
	IndexData map[string]int `xml:"-"`
}

func (b *XMLBank) String() string {
	result, _ := xml.MarshalIndent(b, "", "    ")
	return string(result)
}

type Bank struct {
	path string
	data *XMLBank
}

func NewBank(name string) (*Bank, error) {
	bankFilePath, err := GetSC2BankPath(name)
	if err != nil {
		return nil, fmt.Errorf("GetSC2BankPath() error: %w", err)
	}
	return &Bank{
		path: bankFilePath,
		data: new(XMLBank),
	}, nil
}

func (b *Bank) Load() error {
	bankFile, err := os.ReadFile(b.path)
	if err != nil {
		return fmt.Errorf(" os.ReadFile() error: %w", err)
	}
	err = xml.Unmarshal(bankFile, b.data)
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

func (b *Bank) CreateSection(section string) {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		b.data.Sections = append(b.data.Sections, &XMLSection{
			Name:      section,
			IndexData: map[string]int{},
		})
		sectionIndex = len(b.data.Sections) - 1
		b.data.IndexData[section] = sectionIndex
	}
}

func (b *Bank) LoadSectionNames(index int, count int) []string {
	var result []string
	for _, s := range b.data.Sections[index : index+count] {
		result = append(result, s.Name)
	}
	return result
}

func (b *Bank) RemoveSection(section string) {
	sectionIndex, ok := b.data.IndexData[section]
	if ok {
		b.data.Sections = append(b.data.Sections[:sectionIndex], b.data.Sections[sectionIndex+1:]...)
		delete(b.data.IndexData, section)
	}
}

func (b *Bank) SectionExists(section string) bool {
	_, ok := b.data.IndexData[section]
	return ok
}

func (b *Bank) SectionsCount() int {
	return len(b.data.Sections)
}

func (b *Bank) StoreKey(section string, key string, value BankValue) {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		b.data.Sections = append(b.data.Sections, &XMLSection{
			Name:      section,
			IndexData: map[string]int{},
		})
		sectionIndex = len(b.data.Sections) - 1
		b.data.IndexData[section] = sectionIndex
	}
	sectionData := b.data.Sections[sectionIndex]
	keyIndex, ok := sectionData.IndexData[key]
	if !ok {
		sectionData.Keys = append(sectionData.Keys, &XMLKey{
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
}

func (b *Bank) LoadKey(section string, key string) (BankValue, bool) {
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

func (b *Bank) LoadKeys(section string, index int, count int) ([]string, []BankValue, bool) {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		return nil, nil, false
	}
	sectionData := b.data.Sections[sectionIndex]
	var keys []string
	var values []BankValue
	for _, k := range sectionData.Keys[index : index+count] {
		keys = append(keys, k.Name)
		values = append(values, BankValue{
			Type:  k.Value.Value.Name.Local,
			Value: k.Value.Value.Value,
		})
	}
	return keys, values, true
}

func (b *Bank) RemoveKey(section string, key string) {
	sectionIndex, ok := b.data.IndexData[section]
	if ok {
		sectionData := b.data.Sections[sectionIndex]
		keyIndex, ok := sectionData.IndexData[key]
		if ok {
			sectionData.Keys = append(sectionData.Keys[:keyIndex], sectionData.Keys[:keyIndex+1]...)
			delete(sectionData.IndexData, key)
		}
	}
}

func (b *Bank) KeyExists(section string, key string) bool {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		return false
	}
	sectionData := b.data.Sections[sectionIndex]
	_, ok = sectionData.IndexData[key]
	return ok
}

func (b *Bank) KeysCount(section string) int {
	sectionIndex, ok := b.data.IndexData[section]
	if !ok {
		return 0
	}
	sectionData := b.data.Sections[sectionIndex]
	return len(sectionData.Keys)
}
