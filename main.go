package main

import (
  "encoding/csv"
  "fmt"
  "log"
  "os"
  "io"
  "time"
  "strconv"
)

const (
  DTTM_FORMAT = "2006-01-02 15:04:05"
)

var (
)


type YellowRow struct {
  VendorID string
  PickupDttm time.Time
  DropoffDttm time.Time
  PassengerCnt int
  TripDistance string
  RatecodeID string
  StoreAndFwdFlag string
  PULocationID string
  DOLcationID string
  PaymentType string
  FareAmount string
  Extra string
  MTATax string
  TipAmt string
  TollsAmt string
  ImprovementSurcharge string
  TotalAmt string
  CongestionSurcharge string
}

func processTime(st string) time.Time {
  loc, _ := time.LoadLocation("America/New_York")

  // FIXME: handle error!
  dttm, _ := time.ParseInLocation(DTTM_FORMAT, st, loc)
  return dttm
  //dropoffDttm, err := time.ParseInLocation(DTTM_FORMAT, row[2], loc)
}

func processInt(st string) int {
  i, err := strconv.Atoi(st)

  if err != nil {
    return -1
  }

  return i
}

func convertRow(row []string) YellowRow {
  // FIXME: shortcut for header for now
  if row[0] == "VendorID" {
    return YellowRow{}
  }

  return YellowRow{
    VendorID: row[0],
    PickupDttm: processTime(row[1]),
    DropoffDttm: processTime(row[2]),
    PassengerCnt: processInt(row[3]),
    TripDistance: row[4],
    RatecodeID: row[5],
    StoreAndFwdFlag: row[6],
    PULocationID: row[7],
    DOLcationID: row[8],
    PaymentType: row[9],
    FareAmount: row[10],
    Extra: row[11],
    MTATax: row[12],
    TipAmt: row[13],
    TollsAmt: row[14],
    ImprovementSurcharge: row[15],
    TotalAmt: row[16],
    CongestionSurcharge: row[17],
  }
}

func processRow(row []string) bool {
  if len(row) != 18 {
    return false
  }

  yrow := convertRow(row)

  fmt.Println(yrow)

  return true
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

    processRow(record)
  }
}


func main() {

  //readFile("data/yellow_tripdata_2019-06.csv")
  readFile("data/yellow_head.csv")
}
