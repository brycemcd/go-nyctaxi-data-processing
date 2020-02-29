package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	DTTM_FORMAT = "2006-01-02 15:04:05"
)

var ()

type YellowRow struct {
	rawData              []string
	VendorID             int
	PickupDttm           time.Time
	DropoffDttm          time.Time
	PassengerCnt         int
	TripDistance         float32
	RatecodeID           string
	StoreAndFwdFlag      string
	PULocationID         string
	DOLcationID          string
	PaymentType          string
	FareAmount           string
	Extra                string
	MTATax               string
	TipAmt               string
	TollsAmt             string
	ImprovementSurcharge string
	TotalAmt             string
	CongestionSurcharge  string
	Validated            bool
	Valid                bool
}

func processTime(st string) time.Time {
	loc, _ := time.LoadLocation("America/New_York")

	// FIXME: handle error!
	dttm, _ := time.ParseInLocation(DTTM_FORMAT, st, loc)
	return dttm
	//dropoffDttm, err := time.ParseInLocation(DTTM_FORMAT, row[2], loc)
}

func convertRow(row []string) YellowRow {

	// NOTE: we don't actually know it's valid yet. When setting the values below
	// any of the validation methods may return false to invalidate the record
	isValid := true

	return YellowRow{
		rawData:              row,
		VendorID:             validateVendorId(&row[0], &isValid),
		PickupDttm:           processTime(row[1]),
		DropoffDttm:          processTime(row[2]),
		PassengerCnt:         validatePassengerCnt(&row[3], &isValid),
		TripDistance:         validateTripDistance(&row[4], &isValid),
		RatecodeID:           row[5],
		StoreAndFwdFlag:      row[6],
		PULocationID:         row[7],
		DOLcationID:          row[8],
		PaymentType:          row[9],
		FareAmount:           row[10],
		Extra:                row[11],
		MTATax:               row[12],
		TipAmt:               row[13],
		TollsAmt:             row[14],
		ImprovementSurcharge: row[15],
		TotalAmt:             row[16],
		CongestionSurcharge:  row[17],
		Validated:            false,
		Valid:                isValid,
	}
}

func processRow(row []string) YellowRow {
	if len(row) != 18 {
		return YellowRow{Valid: false, Validated: true}
	}

	// FIXME: shortcut for header for now
	if row[0] == "VendorID" {
		return YellowRow{Valid: false, Validated: true}
	}

	yrow := convertRow(row)

	return yrow
}

func validatePassengerCnt(passCnt *string, valid *bool) int {
	i, err := strconv.Atoi(*passCnt)

	if err != nil {
		*valid = false
		return -1
	}

	if i <= 0 || i > 5 {
		*valid = false
		return -1
	}

	return i
}

func validateTripDistance(tripDist *string, valid *bool) float32 {

	fmt.Sprintf("tripDist is %s", *tripDist)
	tripDistf, err := strconv.ParseFloat(*tripDist, 32)

	if err != nil {
		*valid = false
		return -1.0
	}

	if tripDistf > 100.0 || tripDistf <= 0 {
		*valid = false
		return -1.0
	}

	return float32(tripDistf)
}

func validateVendorId(vendorId *string, valid *bool) int {

	vid, err := strconv.Atoi(*vendorId)
	if err != nil {
		*valid = false
		return -1
	}

	// NOTE: according to the data dictionary, the only valid values are 1 or 2
	// https://www1.nyc.gov/assets/tlc/downloads/pdf/data_dictionary_trip_records_yellow.pdf
	if vid != 1 && vid != 2 {
		*valid = false
		return -1
	}

	return vid
}

func invalidateRecord(record *YellowRow) *YellowRow {
	record.Valid = false
	record.Validated = false

	return record
}

func validateRecord(record *YellowRow) *YellowRow {
	//validateVendorId(record)

	return record
}

func readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		processedRecord := processRow(record)
		//validatedRecord := validateRecord(&processedRecord)
		fmt.Println(processedRecord)
	}
}

func main() {

	//readFile("data/yellow_tripdata_2019-06.csv")
	readFile("data/yellow_head.csv")

	//i, j := 42, 2701

	//p := &i         // point to i
	//fmt.Println(p) // read i through the pointer
	//*p = 21         // set i through the pointerinter
	//fmt.Println(i)  // see the new value of i

	//p = &j         // pointerint to j
	//*p = *p / 37   // divide j through the pointer
	//fmt.Println(j) // see the new value of j
}
