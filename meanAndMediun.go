package main

import (
	"fmt"
	"math"
	"sort"
	// "reflect"
	"encoding/csv"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	readFromFile()
	// var mediun int
	// s := []int {1,2,3,4,5,6,7,8,9,10} // unsorted
	// sum := 0;
	// for index, value := range s {
	//   fmt.Println(index)
	//   sum += value
	// }
	//
	// avg := sum / len(s)
	// fmt.Printf("The Mean is ... %d ", avg)
	//
	// sort.Ints(s)
	// fmt.Println("len of S", len(s))
	// if (len(s) % 2 == 0) {
	//   fl := (s[(len(s) / 2) + 1] + s[(len(s) / 2)]) / 2
	//   fmt.Println(fl)
	// } else {
	//   temp := float64(len(s)) / 2.0
	//   temp1 := math.Floor(temp)
	//   // fmt.Println(temp1)
	//   fmt.Println("%T", reflect.TypeOf(temp1))
	//   res := int64(temp1)
	//   mediun = s[res]
	// }
	//
	// fmt.Printf("The mediun ...%d", mediun)
}

func readFromFile() {
	byteData, err := ioutil.ReadFile("../Ex_Files_UaR_Go/Ex-Lynda/Environmental_Data_Deep_Moor_2015.csv")
	checkErrors(err)
	stringData := BytesToString(byteData)
	readr := csv.NewReader(strings.NewReader(stringData))
	fmt.Println(readr)
	data, err := readr.ReadAll()
	checkErrors(err)
	fmt.Println(data[1])

	airTemp := 0.0
	airTempAry := make([]float64, 0)
	Barometric_Press := 0.0
	Barometric_PressAry := make([]float64, 0)
	Wind_Speed := 0.0
	Wind_SpeedAry := make([]float64, 0)

	fmt.Printf("The length %d ", len(data), "\n")
	for index, value := range data {
		if index != 0 {
			temp, err := strconv.ParseFloat(value[7], 64)
			checkErrors(err)
			Wind_Speed += float64(temp)
			Wind_SpeedAry = append(Wind_SpeedAry, float64(temp))
			temp2, err := strconv.ParseFloat(value[2], 64)
			checkErrors(err)
			Barometric_Press += float64(temp2)
			Barometric_PressAry = append(Barometric_PressAry, float64(temp2))
			temp3, err := strconv.ParseFloat(value[1], 64)
			checkErrors(err)
			airTemp += float64(temp3)
			airTempAry = append(airTempAry, float64(temp3))
		}
	}

	datLen := float64(len(data))
	airTemp = airTemp / datLen
	Barometric_Press = Barometric_Press / datLen
	Wind_Speed = Wind_Speed / datLen

	fmt.Println("\n Air Temp : ", airTemp)
	fmt.Println("\n Barometric Press: ", Barometric_Press)
	fmt.Println("\n Wind Speed : ", Wind_Speed)

	airTempMean := mean(airTempAry)
	Barometric_PressMean := mean(Barometric_PressAry)
	Wind_SpeedMean := mean(Wind_SpeedAry)

	fmt.Println("\n Air Temp Mean: ", airTempMean)
	fmt.Println("\n Barometric Press Mean: ", Barometric_PressMean)
	fmt.Println("\n Wind Speed Mean: ", Wind_SpeedMean)

}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func mean(data []float64) float64 {
	dataLen := len(data)
	var mediun float64
	sort.Float64sAreSorted(data)
	if dataLen%2 == 0 {
		fl := (data[(dataLen/2)+1] + data[(dataLen/2)]) / 2
		fmt.Println(fl)
	} else {
		temp := float64(dataLen) / 2.0
		temp1 := math.Floor(temp)
		// fmt.Println(temp1)
		// fmt.Println("%T", reflect.TypeOf(temp1))
		res := int64(temp1)
		mediun = data[res]
	}
	return mediun
}

func BytesToString(data []byte) string {
	return string(data[:])
}
