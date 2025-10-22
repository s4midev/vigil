package main

import "time"

// a stupid amount of api response typing

type APIResponse struct {
	Embedded Embedded `json:"_embedded"`
	Links    Links    `json:"_links"`
	Page     Page     `json:"page"`
}

type Embedded struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	Locale          string           `json:"locale"`
	Images          []Image          `json:"images"`
	Sales           Sales            `json:"sales"`
	Dates           Dates            `json:"dates"`
	Classifications []Classification `json:"classifications"`
	AgeRestrictions AgeRestriction   `json:"ageRestrictions"`
	Ticketing       Ticketing        `json:"ticketing"`
	Links           Links            `json:"_links"`
	Embedded        *EmbeddedNested  `json:"_embedded,omitempty"`
	Promoter        *Promoter        `json:"promoter,omitempty"`
	Promoters       []Promoter       `json:"promoters,omitempty"`
	Info            string           `json:"info,omitempty"`
	PleaseNote      string           `json:"pleaseNote,omitempty"`
	TicketLimit     *TicketLimit     `json:"ticketLimit,omitempty"`
}

type EmbeddedNested struct {
	Venues      []Venue      `json:"venues"`
	Attractions []Attraction `json:"attractions"`
}

type Image struct {
	Ratio    string `json:"ratio"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Fallback bool   `json:"fallback"`
}

type Sales struct {
	Public   Sale   `json:"public"`
	Presales []Sale `json:"presales,omitempty"`
}

type Sale struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
	StartTBD      bool      `json:"startTBD"`
	StartTBA      bool      `json:"startTBA"`
	Name          string    `json:"name,omitempty"`
}

type Dates struct {
	Start            DateStart `json:"start"`
	Timezone         string    `json:"timezone"`
	Status           Status    `json:"status"`
	SpanMultipleDays bool      `json:"spanMultipleDays"`
}

type DateStart struct {
	LocalDate      string    `json:"localDate"`
	LocalTime      string    `json:"localTime"`
	DateTime       time.Time `json:"dateTime"`
	DateTBD        bool      `json:"dateTBD"`
	DateTBA        bool      `json:"dateTBA"`
	TimeTBA        bool      `json:"timeTBA"`
	NoSpecificTime bool      `json:"noSpecificTime"`
}

type Status struct {
	Code string `json:"code"`
}

type Classification struct {
	Primary  bool     `json:"primary"`
	Segment  Category `json:"segment"`
	Genre    Category `json:"genre"`
	SubGenre Category `json:"subGenre"`
	Type     Category `json:"type"`
	SubType  Category `json:"subType"`
	Family   bool     `json:"family"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AgeRestriction struct {
	LegalAgeEnforced bool `json:"legalAgeEnforced"`
}

type Ticketing struct {
	SafeTix             FeatureFlag `json:"safeTix"`
	AllInclusivePricing FeatureFlag `json:"allInclusivePricing"`
}

type FeatureFlag struct {
	Enabled bool `json:"enabled"`
}

type Links struct {
	Self        Link   `json:"self"`
	Attractions []Link `json:"attractions,omitempty"`
	Venues      []Link `json:"venues,omitempty"`
}

type Link struct {
	Href string `json:"href"`
}

type Venue struct {
	Name                    string         `json:"name"`
	Type                    string         `json:"type"`
	ID                      string         `json:"id"`
	Test                    bool           `json:"test"`
	URL                     string         `json:"url"`
	Locale                  string         `json:"locale"`
	PostalCode              string         `json:"postalCode,omitempty"`
	Timezone                string         `json:"timezone"`
	City                    NamedEntity    `json:"city"`
	Country                 Country        `json:"country"`
	Address                 Address        `json:"address"`
	Location                Location       `json:"location"`
	Markets                 []NamedEntity  `json:"markets,omitempty"`
	DMAs                    []DMA          `json:"dmas,omitempty"`
	BoxOfficeInfo           *BoxOfficeInfo `json:"boxOfficeInfo,omitempty"`
	ParkingDetail           string         `json:"parkingDetail,omitempty"`
	AccessibleSeatingDetail string         `json:"accessibleSeatingDetail,omitempty"`
	GeneralInfo             *GeneralInfo   `json:"generalInfo,omitempty"`
	UpcomingEvents          EventCounts    `json:"upcomingEvents"`
	ADA                     *ADAInfo       `json:"ada,omitempty"`
	Links                   Links          `json:"_links"`
	Images                  []Image        `json:"images,omitempty"`
}

type Attraction struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	Locale          string           `json:"locale"`
	ExternalLinks   *ExternalLinks   `json:"externalLinks,omitempty"`
	Images          []Image          `json:"images"`
	Classifications []Classification `json:"classifications"`
	UpcomingEvents  EventCounts      `json:"upcomingEvents"`
	Links           Links            `json:"_links"`
}

type ExternalLinks struct {
	MusicBrainz []ExternalLink `json:"musicbrainz,omitempty"`
}

type ExternalLink struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type NamedEntity struct {
	Name string `json:"name"`
	ID   string `json:"id,omitempty"`
}

type Country struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

type Address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2,omitempty"`
}

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type DMA struct {
	ID int `json:"id"`
}

type BoxOfficeInfo struct {
	OpenHoursDetail       string `json:"openHoursDetail,omitempty"`
	PhoneNumberDetail     string `json:"phoneNumberDetail,omitempty"`
	AcceptedPaymentDetail string `json:"acceptedPaymentDetail,omitempty"`
	WillCallDetail        string `json:"willCallDetail,omitempty"`
}

type GeneralInfo struct {
	GeneralRule string `json:"generalRule,omitempty"`
	ChildRule   string `json:"childRule,omitempty"`
}

type ADAInfo struct {
	ADAPhones     string `json:"adaPhones,omitempty"`
	ADACustomCopy string `json:"adaCustomCopy,omitempty"`
	ADAHours      string `json:"adaHours,omitempty"`
}

type EventCounts struct {
	Ticketmaster int `json:"ticketmaster,omitempty"`
	Universe     int `json:"universe,omitempty"`
	Moshtix      int `json:"moshtix,omitempty"`
	Total        int `json:"_total"`
	Filtered     int `json:"_filtered"`
}

type Promoter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type TicketLimit struct {
	Info string `json:"info"`
}

type Page struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}
