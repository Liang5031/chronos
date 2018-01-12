package lunar

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetSolarTerm(t *testing.T) {
	//GetSolarTerm(time.Now(), 0)
}

func TestYearDays(t *testing.T) {
	//log.Println(YearDays(1900))
}

func TestGetDayString(t *testing.T) {
	//log.Println(GetDayString(20))
	//reader := transform.NewReader(strings.NewReader(, simplifiedchinese.GB18030.NewDecoder())
	//all, _ := ioutil.ReadAll(reader)
	fmt.Println(strconv.Itoa(0x97783), strconv.Itoa(0x97bd0), strconv.Itoa(0x97c36), strconv.Itoa(0xb0b6f), strconv.Itoa(0xc9274), strconv.Itoa(0xc91aa))
}

func TestGetTerm(t *testing.T) {
	i := GetTerm(2018, 3)
	log.Println(i)
}

func TestGetZodiac(t *testing.T) {
	log.Println(GetZodiac(time.Now()))
}

func TestSolar2Lunar(t *testing.T) {
	log.Println(Solar2Lunar(time.Now()))
	log.Println(Solar2Lunar(time.Date(1900, 2, 18, 12, 12, 12, 0, time.Local)))

}

func TestStemBranchYear(t *testing.T) {
	log.Println(StemBranchYear(2017))

}

func TestStemBranchMonth(t *testing.T) {
	log.Println(StemBranchMonth(2017, 9))
}

func TestStemBranchDay(t *testing.T) {
	log.Println(StemBranchDay(2017, 11, 14))

}

func TestStemBranchHour(t *testing.T) {
	for i := 1; i <= 23; i++ {
		log.Println(StemBranchHour(i))
	}
}
