package lenconv

const fToMFactor = 0.3048

func FToM(f Foot) Meter { return Meter(f * fToMFactor) }

func MToF(m Meter) Foot { return Foot(m / fToMFactor) }
