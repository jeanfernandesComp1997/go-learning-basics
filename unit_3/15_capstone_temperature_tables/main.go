package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 1.8) + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

func drawTable(header1, header2 string, rows int, getRow getRowFn) {
	drawTableHeader(header1, header2)
	drawTableValues(rows, getRow)
	drawTableFooter()
}

func drawTableHeader(header1, header2 string) {
	fmt.Println(line)
	fmt.Printf(rowFormat, header1, header2)
	fmt.Println(line)
}

func drawTableValues(rows int, getRow getRowFn) {
	for row := 0; row <= rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
}

func drawTableFooter() {
	fmt.Println(line)
}

func celsiusToFahrenheit(row int) (string, string) {
	celsiusDregrees := celsius(row*5 - 40)
	fahrenheitDregrees := celsiusDregrees.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, celsiusDregrees)
	cell2 := fmt.Sprintf(numberFormat, fahrenheitDregrees)
	return cell1, cell2
}

func fahrenheitToCelsius(row int) (string, string) {
	fahrenheitDegrees := fahrenheit(row*5 - 40)
	celsiusDregrees := fahrenheitDegrees.celsius()
	cell1 := fmt.Sprintf(numberFormat, fahrenheitDegrees)
	cell2 := fmt.Sprintf(numberFormat, celsiusDregrees)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 28, celsiusToFahrenheit)
	println()
	drawTable("°F", "°C", 28, fahrenheitToCelsius)
}
