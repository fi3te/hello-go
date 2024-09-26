package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

const sheetName1 = "Sheet1"
const fileName = "demo.xlsx"

func main() {
	animals := []string{"Cat", "Cattle", "Dog", "Donkey", "Goat", "Horse", "Pig", "Rabbit", "Sheep", "Buffalo", "Chicken", "Duck", "Goose", "Pigeon", "Turkey", "Elephant", "Leopard", "Albatross", "Alligator", "Alpaca", "Anaconda", "Ant", "Antelope", "Antlion", "Ape", "Aphid", "Arctic Fox", "Arctic Wolf"}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	random := rand.New(rand.NewSource(time.Now().Unix()))
	f.SetCellValue(sheetName1, "A1", "Animal")
	f.SetCellValue(sheetName1, "B1", "Count")
	for i := 0; i < 20; i++ {
		f.SetCellValue(sheetName1, "A"+strconv.Itoa(i+2), animals[i])
		f.SetCellValue(sheetName1, "B"+strconv.Itoa(i+2), random.Intn(100))
	}

	if err := f.AddChart(sheetName1, "C1", &excelize.Chart{
		Type: excelize.Col3DClustered,
		Series: []excelize.ChartSeries{
			{
				Name:       sheetName1 + "!$A$1",
				Categories: sheetName1 + "!$A$2:$A$21",
				Values:     sheetName1 + "!$B$2:$B$21",
			},
		},
		Title: []excelize.RichTextRun{
			{
				Text: "Animal count",
			},
		},
	}); err != nil {
		panic(err)
	}

	if err := f.SaveAs(fileName); err != nil {
		panic(err)
	}
}
