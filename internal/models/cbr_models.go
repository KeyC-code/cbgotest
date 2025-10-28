package models

import (
	"encoding/xml"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

var CurrencyMapping = map[string]struct {
	ID      string
	NumCode string
	Name    string
}{
	"AUD": {"R01010", "036", "Австралийский доллар"},
	"GBP": {"R01035", "826", "Фунт стерлингов"},
	"BYR": {"R01090", "974", "Белорусских рублей"},
	"DKK": {"R01215", "208", "Датских крон"},
	"USD": {"R01235", "840", "Доллар США"},
	"EUR": {"R01239", "978", "Евро"},
	"ISK": {"R01310", "352", "Исландских крон"},
	"KZT": {"R01335", "398", "Тенге"},
	"CAD": {"R01350", "124", "Канадский доллар"},
	"NOK": {"R01535", "578", "Норвежских крон"},
	"XDR": {"R01589", "960", "СДР (специальные права заимствования)"},
	"SGD": {"R01625", "702", "Сингапурский доллар"},
	"TRL": {"R01700", "792", "Турецких лир"},
	"UAH": {"R01720", "980", "Гривен"},
	"SEK": {"R01770", "752", "Шведских крон"},
	"CHF": {"R01775", "756", "Швейцарский франк"},
	"JPY": {"R01820", "390", "Иен"},
}
