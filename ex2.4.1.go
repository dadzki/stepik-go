package main

type Terminator struct {
	On    bool
	Ammo  int
	Power int
}

func (c *Terminator) Shoot() bool {
	if !c.On || c.Ammo < 1 {
		return false
	}
	c.Ammo--
	return true
}

func (c *Terminator) RideBike() bool {
	if !c.On || c.Power < 1 {
		return false
	}
	c.Power--
	return true
}

func main() {
	testStruct := Terminator{true, 3, 3}
	testStruct.Shoot()
	testStruct.RideBike()
}
