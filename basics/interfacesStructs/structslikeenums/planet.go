package main

import "strconv"
//http://play.golang.org/p/nFZgQ2mpDx
type surfaceGravity interface {
	surfaceGravity() float64
	surfaceWeight(float64) float64
}
type PlanetName string

func (pn PlanetName) name() *PlanetName {
	return &pn
}
func (pn PlanetName) String() string {
	return string(pn)
}

type planet struct {
	Name   PlanetName
	mass   float64
	radius float64
}
var Planet []*planet = []*planet{}
func NewPlanet(name PlanetName, mass float64, radius float64) (planetMass *planet) {
	planetMass = &planet{Name: name, mass: mass, radius: radius}
	Planet = append(Planet, planetMass)
	return
}

func CreateEnumPlanets() {
	NewPlanet("MERCURY", 3.303e+23, 2.4397e6)
	NewPlanet("VENUS", 4.869e+24, 6.0518e6)
	NewPlanet("EARTH", 5.976e+24, 6.37814e6)
	NewPlanet("MARS", 6.421e+23, 3.3972e6)
	NewPlanet("JUPITER", 1.9e+27, 7.1492e7)
	NewPlanet("SATRUN", 5.688e+26, 6.0268e7)
	NewPlanet("URANUS", 8.686e+25, 2.5559e7)
	NewPlanet("NEPTUNE", 2.024e+26, 2.4746e7)
	NewPlanet("PLUTO", 1.142e+26, 1.5223e9)
}

func (planet planet) surfaceGravity() float64 {
	return G * planet.mass / (planet.radius * planet.radius)
}

func (planet planet) surfaceWeight(otherMass float64) float64 {
	return otherMass * planet.surfaceGravity()
}

func (planet planet) String() string {
	mf := strconv.FormatFloat(planet.mass, 'e', 3, 64)
	rf := strconv.FormatFloat(planet.radius, 'G', 3, 64)
	return planet.Name.String() + " : [ mass : " + mf + ", radius :" + rf + " ]"
}



// universal gravitational constant  (m3 kg-1 s-2)
const (
	G float64 = 6.67300E-11
)

func main() {
	mainDayTest()
//	CreateEnumPlanets()
//	fmt.Println("Planet Mass size >> ")
//	fmt.Println(len(Planet))
//
//	fmt.Println("Enter the weight of earth")
//	var weight float64
//	fmt.Scanf("%f", &weight)
//	if weight < 100 {
//		fmt.Println("Usage: planet <earth_weight>")
//		os.Exit(1)
//	}

}
