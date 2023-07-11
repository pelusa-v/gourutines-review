package race_conditions

import (
	"fmt"
	"math/rand"
	"time"
)

func OgresLifetime(war chan string) {
	race := "ogros"
	defer wg.Done()
	ExploreNewTerrain(race)
	GatherMinerals(race, 3000)

	for {
		ExploreNewTerrain(race)
		GatherMinerals(race, 3000)

		warMessage := <-war
		fmt.Println(warMessage)

		GatherMinerals(race, 2500)
		TrainTroops(race, 100)

		war <- " -- Los ogros atacan por el frente a los humanos --"

	}
}

func HumansLifetime(war chan string) {
	race := "humanos"
	defer wg.Done()
	for {
		TrainTroops(race, 400)
		ExploreNewTerrain(race)
		ExploreNewTerrain(race)
		GatherMinerals(race, 1200)
		TrainTroops(race, 500)

		war <- " -- Los humanos atacan por emboscada a los ogros --"

		warMessage := <-war
		fmt.Println(warMessage)

	}
}

func DefenseAtack(race string) {
	fmt.Printf("La tribu %s se defiende\n", race)
}

func CounterAtack(race string, war chan string) {
	war <- " -- La tribu %s contraataca --"
	fmt.Printf("La tribu %s contraataca\n", race)
}

func GatherMinerals(race string, ammount int) {
	time.Sleep(time.Duration(ammount) * time.Millisecond)
	fmt.Printf("La tribu %s recolectó %d de minerales\n", race, ammount)
}

func TrainTroops(race string, number int) {
	time.Sleep(time.Duration(number*10) * time.Millisecond)
	fmt.Printf("La tribu %s entrenó %d soldados\n", race, number)
}

func ExploreNewTerrain(race string) {
	terrains := []string{"Egipto", "Tanzania", "Roma", "Estambul", "Kolkata", "Delhi", "Congo", "San Marino"}
	time.Sleep(time.Duration(rand.Int63n(3000)) * time.Millisecond)
	fmt.Printf("La tribu %s exploró el territorio %s\n", race, terrains[rand.Int63n(int64(len(terrains)))])
}

func StartWar() {
	warChannel := make(chan string)
	wg.Add(2)

	go OgresLifetime(warChannel)
	go HumansLifetime(warChannel)

	wg.Wait()
}
