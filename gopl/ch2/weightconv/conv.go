package weightconv

const lbToKgFactor = 0.45359237

func PToK(p Pound) Kilogram { return Kilogram(p * lbToKgFactor) }
func KToP(k Kilogram) Pound { return Pound(k / lbToKgFactor) }
