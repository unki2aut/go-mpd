[![Build Status](https://travis-ci.org/unki2aut/go-mpd.svg?branch=master)](https://travis-ci.org/unki2aut/go-mpd) [![Go Coverage](https://github.com/unki2aut/go-mpd/wiki/coverage.svg)](https://raw.githack.com/wiki/unki2aut/go-mpd/coverage.html) [![Go Report Card](https://goreportcard.com/badge/github.com/unki2aut/go-mpd)](https://goreportcard.com/report/github.com/unki2aut/go-mpd) [![GoDoc](https://godoc.org/github.com/unki2aut/go-mpd?status.svg)](https://godoc.org/github.com/unki2aut/go-mpd)
# go-mpd 

Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files.

This project is based on https://github.com/mc2soft/mpd.

## Usage

```go
package main

import (
	"fmt"
	"github.com/unki2aut/go-mpd"
)

func main() {
	mpd := new(mpd.MPD)
	mpd.Decode([]byte(`<MPD type="static" mediaPresentationDuration="PT3M30S">
  <Period>
    <AdaptationSet mimeType="video/mp4" codecs="avc1.42c00d">
      <SegmentTemplate media="../video/$RepresentationID$/dash/segment_$Number$.m4s" initialization="../video/$RepresentationID$/dash/init.mp4" duration="100000" startNumber="0" timescale="25000"/>
      <Representation id="180_250000" bandwidth="250000" width="320" height="180" frameRate="25"/>
      <Representation id="270_400000" bandwidth="400000" width="480" height="270" frameRate="25"/>
      <Representation id="360_800000" bandwidth="800000" width="640" height="360" frameRate="25"/>
      <Representation id="540_1200000" bandwidth="1200000" width="960" height="540" frameRate="25"/>
      <Representation id="720_2400000" bandwidth="2400000" width="1280" height="720" frameRate="25"/>
      <Representation id="1080_4800000" bandwidth="4800000" width="1920" height="1080" frameRate="25"/>
    </AdaptationSet>
    <AdaptationSet lang="en" mimeType="audio/mp4" codecs="mp4a.40.2" bitmovin:label="English stereo">
      <AudioChannelConfiguration schemeIdUri="urn:mpeg:dash:23003:3:audio_channel_configuration:2011" value="2"/>
      <SegmentTemplate media="../audio/$RepresentationID$/dash/segment_$Number$.m4s" initialization="../audio/$RepresentationID$/dash/init.mp4" duration="191472" startNumber="0" timescale="48000"/>
      <Representation id="1_stereo_128000" bandwidth="128000" audioSamplingRate="48000"/>
    </AdaptationSet>
  </Period>
</MPD>`))

	fmt.Println(mpd.MediaPresentationDuration)
}
```

## Related links
* https://en.wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
* [ISO_IEC_23009-1_2014 Standard](http://standards.iso.org/ittf/PubliclyAvailableStandards/c065274_ISO_IEC_23009-1_2014.zip)

## MPD parsing/generation in other languages
* Javascript - https://github.com/videojs/mpd-parser
* Python - https://github.com/sangwonl/python-mpegdash
* Cpp - https://github.com/bitmovin/libdash
* Java - https://github.com/carlanton/mpd-tools
