package main

type record struct {
	IP string
	ID string
}

type myipApi struct {
	v4 string
	v6 string
}

type aliApi struct {
	accessKeyId     string
	accessKeySecret string
}

type enableFlag struct {
	v4 string
	v6 string
}

type config struct {
	currentIP string
	domain    string
	subDomain string
	record    record
}

var (
	connTestAPI string = "https://connect.rom.miui.com/generate_204"
	myip        myipApi
	aliApiToken aliApi
	flag        enableFlag
	v4          config
	v6          config
)
