# Exercise 2.2

Write a general-purpose unit-conversion program analogous to `cf` that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.

## Solution Implementation

This solution implements a comprehensive unit conversion program that converts any given number into all available units across three different measurement categories.

### Supported Conversions

#### Temperature
- Celsius (°C) ↔ Fahrenheit (°F)
- Celsius (°C) ↔ Kelvin (K)
- Fahrenheit (°F) ↔ Kelvin (K)

#### Length
- Meters (m) ↔ Feet (ft)

#### Weight
- Kilograms (kg) ↔ Pounds (lb)

### Package Structure

The program is organized into three main packages:

- `tempconv`: Temperature conversion utilities
  - Types: `Celsius`, `Fahrenheit`, `Kelvin`
  - Conversions between all temperature units
- `lengthconv`: Length conversion utilities
  - Types: `Meter`, `Foot`
  - Bidirectional conversions
- `weightconv`: Weight conversion utilities
  - Types: `Kilogram`, `Pound`
  - Bidirectional conversions

### Usage

Run the program with one or more numerical arguments:

```bash
go run main.go 100 200 300
```

For each input number, the program will display all possible unit conversions grouped by category.

### Example Output

```
Temperature:
        100°F = 37.78°C = 310.93K
        100°C = 212°F = 373.15K
        100K = -173.15°C = -279.67°F

Length:
        100 m = 328.08 ft
        100 ft = 30.48 m

Weight:
        100 kg = 220.46 lb
        100 lb = 45.36 kg
```

### Conversion Factors Used

- Temperature:
  - °F to °C: (°F - 32) × 5/9
  - °C to K: °C + 273.15
- Length:
  - Meters to Feet: 1 m = 3.28084 ft
- Weight:
  - Kilograms to Pounds: 1 kg = 2.20462 lb