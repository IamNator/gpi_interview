package buyandsell

import (
	"fmt"
)

type Cement struct {
	NoOfCement int
}

func (c *Cement) BuyCement(howmany int) {
	c.NoOfCement += howmany
}

func (c *Cement) SellCement(howmany int) {
	c.NoOfCement -= howmany // o was undefined here.
}

func (c Cement) String() string { //Cement is now passed by value here
	return fmt.Sprintf("%v", c.NoOfCement)
}
