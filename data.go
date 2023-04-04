package main

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target pybadge -ldflags="-X main.YourName=@myontwitter -X main.YourTitleA1='Amazing human' -X main.YourTitleA2='also kind'"
//
// See Makefile for more info.
var (
	YourName, YourTitleA1, YourTitleA2, YourTitleB1, YourTitleB2     string
	YourMarqueeTop, YourMarqueeMiddle, YourMarqueeBottom, YourQRText string
)

const (
	DefaultName          = "@_CONEJO"
	DefaultTitleA1       = "Go Compiler"
	DefaultTitleA2       = "Small Places"
	DefaultMarqueeTop    = "This badge"
	DefaultMarqueeMiddle = "runs"
	DefaultMarqueeBottom = "TINYGO"
	DefaultQRText        = "https://gopherbadge.com"
	DefaultTitleB1       = "I enjoy"
	DefaultTitleB2       = "TINYGO"
)
