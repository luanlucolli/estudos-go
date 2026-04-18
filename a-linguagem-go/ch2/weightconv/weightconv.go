package weightconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func (p Pound) String() string {
	if p == 1 || p == -1 {
		return fmt.Sprintf("%glb", p)
	}
	return fmt.Sprintf("%glbs", p)
}
