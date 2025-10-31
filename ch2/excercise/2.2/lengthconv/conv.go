package lengthconv

const feetPerMeter = 3.28084

func MToFt(m Meter) Foot {
	return Foot(m * feetPerMeter)
}

func FtToM(f Foot) Meter {
	return Meter(f / feetPerMeter)
}
