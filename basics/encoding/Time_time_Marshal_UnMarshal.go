package main

import (
	"fmt"
	"time"
)

var marshalledBinaryTime []byte
var marshalledJsonTime []byte
var marshalledTextTime []byte
var marshalledGobTime []byte

var err error

func main() {
	// layout shows by example how the reference time should be represented.
	const layout = "July 16, 2014 at 12:00am (PST)"
	t := time.Date(2015, time.July, 16, 0, 0, 0, 0, time.Local)
	fmt.Println("****** Time Before Binary Decode ******8")

	fmt.Println(t.Format(layout))
	fmt.Println(t.UTC().Format(layout))

	if marshalledBinaryTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(marshalledBinaryTime)
	}
	t.UnmarshalBinary(marshalledBinaryTime)
	fmt.Println("UnMarshal Binary Time " + t.String())

	t = time.Date(2015, time.March, 4, 0, 0, 0, 0, time.Local)
	fmt.Println("****** Time Before JSON Decode ******8")

	fmt.Println(t.Format(layout))

	if marshalledJsonTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(marshalledJsonTime)
	}
	t.UnmarshalJSON(marshalledJsonTime)
	fmt.Println("UnMarshal JSON Time " + t.String())

	t = time.Date(2015, time.August, 22, 0, 0, 0, 0, time.Local)
	fmt.Println("****** Time Before Text Decode ******8")

	fmt.Println(t.Format(layout))
	if marshalledTextTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(marshalledTextTime)
	}
	t.UnmarshalText(marshalledTextTime)
	fmt.Println("UnMarshal Text Time " + t.String())

	t = time.Date(2015, time.July, 16, 0, 0, 0, 0, time.Local)
	fmt.Println("****** Time Before Gob Decode ******8")
	fmt.Println(t.Format(layout))
	//GobDecode implements the gob.GobDecoder interface
	if marshalledGobTime, err = t.GobEncode(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(marshalledGobTime)
	}
	t.GobDecode(marshalledGobTime)
	fmt.Println("UnMarshal Gob Time " + t.String())

}
