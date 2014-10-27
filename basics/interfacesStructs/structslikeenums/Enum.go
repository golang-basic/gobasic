package main

type Enum interface{
	name() string
	ordinal() int
	valueOf(struct{}, string)

}
