package weightconv

const poundsPerKilogram = 2.20462

func KgToLb(k Kilogram) Pound {
	return Pound(k * poundsPerKilogram)
}

func LbToKg(p Pound) Kilogram {
	return Kilogram(p / poundsPerKilogram)
}
