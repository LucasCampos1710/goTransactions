package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// func main() {
// 	result, err := soma(20,2)
// 	fmt.Println(result, err)
// }

// func soma(a, b int) (int, error) {
// 	if a+b > 10	{
// 		return 0, fmt.Errorf("Soma maior que 10")
// 	}
// 	return a + b, nil
// }

type Car struct {
	Name string
	Model string
	Price float64
}
var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Ferrai", Model: "S200", Price: 100})
	cars = append(cars, Car{Name: "Porsche", Model: "Cayene", Price: 200})
	cars = append(cars, Car{Name: "BMW", Model: "S250", Price: 300})
	cars = append(cars, Car{Name: "Mercedes-Bens", Model: "Sw300", Price: 400})
}

func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Logger.Fatal(e.Start(":8080"))
	// carro := Car{"Fusca", "Vw"}
	// fmt.Println(carro.Name)
	// fmt.Println(carro.Model)
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, model, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	} 
	_, err = stmt.Exec(car.Name, car.Model, car.Price)
	if err != nil {
		return err
	}
	return nil
}
