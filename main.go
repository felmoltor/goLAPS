package main

import (
	"fmt"
	"log"
)

func getLapsPassword(args *GetArgs) {
	gl := GoLaps{
		Credentials: Credentials{
			Username: args.Username,
			Password: args.Password,
			Domain:   args.Domain,
		},
		TargetDC:           args.Dc,
		ComputernameFilter: args.Filter,
	}
	gl.InitLogging()
	l, err := gl.ConnectToDC()
	if err != nil {
		log.Fatal("Error connecting to DC" + err.Error())
	}
	err = gl.BindToDC(l)
	if err != nil {
		log.Fatal("Error binding to DC" + err.Error())
	}
	_, err = gl.GetDomainDN(l)
	if err != nil {
		log.Fatal("Error getting domain DN" + err.Error())
	}
	_, err = gl.SearchComputersWithLaps(l)
	if err != nil {
		log.Fatal("Error searching for computers with LAPS" + err.Error())
	}
	gl.CloseConnection(l)
}

func setLapsPassword(args *SetArgs) {
	gl := GoLaps{
		Credentials: Credentials{
			Username: args.Username,
			Password: args.Password,
			Domain:   args.Domain,
		},
		TargetDC:       args.Dc,
		Targetcomputer: args.Target,
	}
	gl.InitLogging()
	l, err := gl.ConnectToDC()
	if err != nil {
		log.Fatal("Error connecting to DC" + err.Error())
	}
	err = gl.BindToDC(l)
	if err != nil {
		log.Fatal("Error binding to DC" + err.Error())
	}
	_, err = gl.GetDomainDN(l)
	if err != nil {
		log.Fatal("Error getting domain DN" + err.Error())
	}
	// TODO: Complete the method to set the LAPS password of this computer
	err = gl.SetLapsPassword(l, args.Lapspass)
	if err != nil {
		log.Fatal("Error setting the LAPS password" + err.Error())
	}
	gl.CloseConnection(l)
}

// main
func main() {

	args := parseArgs()

	if args.get != nil {
		getLapsPassword(args.get)
	} else if args.set != nil {
		setLapsPassword(args.set)
	} else {
		fmt.Println("No command specified")
	}
}
