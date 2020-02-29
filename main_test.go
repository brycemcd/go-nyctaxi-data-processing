package main

import (
  "testing"
  //"fmt"
)

func TestValidateVendorId(t *testing.T) {
  /* Valid vendor Ids are 1 and 2 */

  vendorIdTests := []struct {
    vid string
    valid bool
    correctConvertedVID int
    shouldBeValid bool
  }{
    {"1", true, 1, true},
    {"2", true, 2, true},
    {"foo", true, -1, false},
  }

  for _, test := range vendorIdTests {
    vid := validateVendorId(&test.vid, &test.valid)

    if(vid != test.correctConvertedVID || test.valid != test.shouldBeValid) {
      t.Errorf("vid: %d != %d; %t != %t", vid, test.correctConvertedVID, test.valid, test.shouldBeValid)
    }
  }
}

func TestValidateTripDistance(t *testing.T) {
  /* tripDistance should be a float in the range of (0.00, 100.00] */

  tdTests := []struct {
    td string
    valid bool
    correctConvertedTd float32
    shouldBeValid bool
  }{
    {"1.00", true, 1.00, true},
    {"1.23", true, 1.23, true},
    {"foo", true, -1.0, false},
    {"101.00", true, -1.0, false},
    {"-1.00", true, -1.0, false},
    {"0.00", true, -1.0, false},
  }

  for _, test := range tdTests {
    convertedTd := validateTripDistance(&test.td, &test.valid)

    if convertedTd != test.correctConvertedTd || test.valid != test.shouldBeValid {
      t.Errorf("%s in should produce %f, not %f; valid should be %t, not %t",
        test.td, test.correctConvertedTd, convertedTd, test.shouldBeValid, test.valid)
    }
  }
}
