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

func Test_Time_Marshal_Unmarshall() {

	loc, _ := time.LoadLocation("Europe/Berlin")
	//Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location)

	t := time.Date(2013, time.July, 16, 1, 30, 26, 0, loc)
	// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
	// As this time is unlikely to come up in practice, the IsZero method gives
	// a simple way of detecting a time that has not been initialized explicitly.
	fmt.Println("Is time Not Initialized : ", t.IsZero()) // returns false

	fmt.Println("****** Time Before Binary Decode ******8")
	fmt.Println(t.Local()) // not the zero value time

	/**
	Following gives only Zero time
	 */

	const layout = "Jul 16, 2013 at 12:00am (PST)"

	if marshalledBinaryTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("****** Time After Binary Decode ******8")
		fmt.Println("Binary Encoded time ", marshalledBinaryTime)
	}
	t.UnmarshalBinary(marshalledBinaryTime)
	fmt.Println("Binary Decoded Time " + t.Local().String())

	loc, _ = time.LoadLocation("America/Toronto")
	t = time.Date(2013, time.March, 4, 0, 0, 0, 0, loc)
	t, _ = time.ParseInLocation(layout, "Jul 16, 2013 at 12:00am (EST)", loc)
	fmt.Println("****** Time Before JSON Decode ******8")
	fmt.Println(t.Format(layout))

	if marshalledJsonTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("****** Time Before JSON Decode ******8")
		fmt.Println(marshalledJsonTime)
	}
	t.UnmarshalJSON(marshalledJsonTime)
	fmt.Println("UnMarshal JSON Time " + t.String())


	loc, _ = time.LoadLocation("America/Phoenix")
	t = time.Date(2013, time.August, 22, 0, 0, 0, 0, loc)
	fmt.Println("****** Time Before Text Decode ******8")

	fmt.Println(t.Format(layout))
	if marshalledTextTime, err = t.MarshalBinary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("****** Time Before TEXT Decode ******8")
		fmt.Println(marshalledTextTime)
	}
	t.UnmarshalText(marshalledTextTime)
	fmt.Println("UnMarshal Text Time " + t.String())


	loc, _ = time.LoadLocation("America/Los_Angeles")
	t = time.Date(2013, time.July, 16, 0, 0, 0, 0, loc)
	fmt.Println("****** Time Before Gob Decode ******8")
	fmt.Println(t.Format(layout))
	//GobDecode implements the gob.GobDecoder interface
	if marshalledGobTime, err = t.GobEncode(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("****** Time Before GOB/Binary Decode ******8")
		fmt.Println(marshalledGobTime)
	}
	t.GobDecode(marshalledGobTime)
	fmt.Println("UnMarshal Gob Time " + t.String())

}

/**


Flaky Code :

// layout shows by example how the reference time should be represented.
	//	const layout = "Jul 16, 2013 at 12:00am (PST)"
	//	//	loc, _ = time.LoadLocation("America/Chicago")
	//
	//	ParseInLocation is like Parse but differs in two important ways.
	//	First, in the absence of time zone information, Parse interprets a time as UTC;
	//	ParseInLocation interprets the time as in the given location.
	//	Second, when given a zone offset or abbreviation, Parse tries to match it
	//	against the Local location; ParseInLocation uses the given location.
	//
	t2, _ := time.ParseInLocation(layout, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t2)
	fmt.Println(t2.Local())

	t2, _ = time.ParseInLocation(layout, t.Format(layout), loc)
	fmt.Println(t2.IsZero()) // returns true
	fmt.Println(t2)
	fmt.Println(t2.Local())

	t2, _ = time.ParseInLocation(layout, t.String(), loc)
	fmt.Println(t2.IsZero()) // returns true
	fmt.Println(t2)

	fmt.Println(t2.Local())

	fmt.Println(t.Format(layout))
	fmt.Println(t.UTC().Format(layout))

*/
