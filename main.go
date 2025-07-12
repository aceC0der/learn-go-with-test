package main

import "fmt"

type RockClimber struct {
	rocksClimbed int
}

func (rc *RockClimber) placeSafeties() {
	fmt.Println("placing my safties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.placeSafeties()
	}
}


func main() {
	rc := &RockClimber{}
	for range 11 {
		rc.climbRock()
	}
}