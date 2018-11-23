// Package mpd implements parsing and generating of MPEG-DASH Media Presentation Description (MPD) files.
package mpd

import (
	"bytes"
	"encoding/xml"
	"io"
	"regexp"
)

// http://mpeg.chiariglione.org/standards/mpeg-dash
// https://www.brendanlong.com/the-structure-of-an-mpeg-dash-mpd.html
// http://standards.iso.org/ittf/PubliclyAvailableStandards/MPEG-DASH_schema_files/DASH-MPD.xsd

var emptyElementRE = regexp.MustCompile(`></[A-Za-z]+>`)

// MPD represents root XML element.
type MPD struct {
	XMLNS                      *string    `xml:"xmlns,attr"`
	Type                       *string    `xml:"type,attr"`
	MinimumUpdatePeriod        *string    `xml:"minimumUpdatePeriod,attr"`
	AvailabilityStartTime      *string    `xml:"availabilityStartTime,attr"`
	MediaPresentationDuration  *string    `xml:"mediaPresentationDuration,attr"`
	MinBufferTime              *string    `xml:"minBufferTime,attr"`
	SuggestedPresentationDelay *string    `xml:"suggestedPresentationDelay,attr"`
	TimeShiftBufferDepth       *string    `xml:"timeShiftBufferDepth,attr"`
	PublishTime                *string    `xml:"publishTime,attr"`
	Profiles                   string     `xml:"profiles,attr"`
	BaseURL                    []*BaseURL `xml:"BaseURL,omitempty"`
	Period                     *Period    `xml:"Period,omitempty"`
}

// Do not try to use encoding.TextMarshaler and encoding.TextUnmarshaler:
// https://github.com/golang/go/issues/6859#issuecomment-118890463

// Encode generates MPD XML.
func (m *MPD) Encode() ([]byte, error) {
	x := new(bytes.Buffer)
	e := xml.NewEncoder(x)
	e.Indent("", "  ")
	err := e.Encode(m)
	if err != nil {
		return nil, err
	}

	// hacks for self-closing tags
	res := new(bytes.Buffer)
	res.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	res.WriteByte('\n')
	for {
		s, err := x.ReadString('\n')
		if s != "" {
			s = emptyElementRE.ReplaceAllString(s, `/>`)
			res.WriteString(s)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	res.WriteByte('\n')
	return res.Bytes(), err
}

// Decode parses MPD XML.
func (m *MPD) Decode(b []byte) error {
	return xml.Unmarshal(b, m)
}

// Period represents XSD's PeriodType.
type Period struct {
	Start          *string          `xml:"start,attr"`
	ID             *string          `xml:"id,attr"`
	Duration       *string          `xml:"duration,attr"`
	AdaptationSets []*AdaptationSet `xml:"AdaptationSet,omitempty"`
	BaseURL        []*BaseURL       `xml:"BaseURL,omitempty"`
}

// BaseURL represents XSD's BaseURLType.
type BaseURL struct {
	Value                    string  `xml:",chardata"`
	ServiceLocation          *string `xml:"serviceLocation,attr"`
	ByteRange                *string `xml:"byteRange,attr"`
	AvailabilityTimeOffset   *uint64 `xml:"availabilityTimeOffset,attr"`
	AvailabilityTimeComplete *bool   `xml:"availabilityTimeComplete,attr"`
}

// AdaptationSet represents XSD's AdaptationSetType.
type AdaptationSet struct {
	MimeType                string           `xml:"mimeType,attr"`
	ContentType             *string          `xml:"contentType,attr"`
	SegmentAlignment        ConditionalUint  `xml:"segmentAlignment,attr"`
	SubsegmentAlignment     ConditionalUint  `xml:"subsegmentAlignment,attr"`
	StartWithSAP            *uint64          `xml:"startWithSAP,attr"`
	SubsegmentStartsWithSAP *uint64          `xml:"subsegmentStartsWithSAP,attr"`
	BitstreamSwitching      *bool            `xml:"bitstreamSwitching,attr"`
	Lang                    *string          `xml:"lang,attr"`
	Par                     *string          `xml:"par,attr"`
	BaseURL                 []*BaseURL       `xml:"BaseURL,omitempty"`
	SegmentTemplate         *SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	ContentProtections      []Descriptor     `xml:"ContentProtection,omitempty"`
	Representations         []Representation `xml:"Representation,omitempty"`
}

// Representation represents XSD's RepresentationType.
type Representation struct {
	ID                 *string          `xml:"id,attr"`
	Width              *uint64          `xml:"width,attr"`
	Height             *uint64          `xml:"height,attr"`
	FrameRate          *string          `xml:"frameRate,attr"`
	Bandwidth          *uint64          `xml:"bandwidth,attr"`
	AudioSamplingRate  *string          `xml:"audioSamplingRate,attr"`
	Codecs             *string          `xml:"codecs,attr"`
	SAR                *string          `xml:"sar,attr"`
	ScanType           *string          `xml:"scanType,attr"`
	ContentProtections []Descriptor     `xml:"ContentProtection,omitempty"`
	SegmentTemplate    *SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	BaseURL            []*BaseURL       `xml:"BaseURL,omitempty"`
}

// Descriptor represents XSD's DescriptorType.
type Descriptor struct {
	SchemeIDURI *string `xml:"schemeIdUri,attr"`
	Value       *string `xml:"value,attr"`
}

// SegmentTemplate represents XSD's SegmentTemplateType.
type SegmentTemplate struct {
	Duration               *uint64          `xml:"duration,attr"`
	Timescale              *uint64          `xml:"timescale,attr"`
	Media                  *string          `xml:"media,attr"`
	Initialization         *string          `xml:"initialization,attr"`
	StartNumber            *uint64          `xml:"startNumber,attr"`
	PresentationTimeOffset *uint64          `xml:"presentationTimeOffset,attr"`
	SegmentTimeline        *SegmentTimeline `xml:"SegmentTimeline,omitempty"`
}

type SegmentTimeline struct {
	S []*SegmentTimelineS `xml:"S"`
}

// SegmentTimelineS represents XSD's SegmentTimelineType's inner S elements.
type SegmentTimelineS struct {
	T *uint64 `xml:"t,attr"`
	D uint64  `xml:"d,attr"`
	R *int64  `xml:"r,attr"`
}
