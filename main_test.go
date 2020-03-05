package main

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestConvertRow(t *testing.T) {
	validRow := "1,2019-06-01 00:09:57,2019-06-01 00:25:54,2,2.00,1,N,158,68,2,11.5,3,0.5,0,0,0.3,15.3,2.5"
	r := csv.NewReader(strings.NewReader(validRow))

	record, err := r.Read()

	if err != nil {
		t.Error("NOOOO")
	}

	yr := convertRow(record)

	if yr.VendorID != 1 {
		t.Errorf("VendorID not expected")
	}

	pickuptm := processTime("2019-06-01 00:09:57")
	if !yr.PickupDttm.Equal(pickuptm) {
		t.Errorf("PickupDttm not expected %s != %s", yr.PickupDttm, pickuptm)
	}

	dropofftm := processTime("2019-06-01 00:25:54")
	if !yr.DropoffDttm.Equal(dropofftm) {
		t.Errorf("dropoffDttm not expected %s != %s", yr.DropoffDttm, dropofftm)
	}

	if yr.PassengerCnt != 2 {
		t.Errorf("PassengerCnt not expected")
	}

	if yr.TripDistance != 2.00 {
		t.Errorf("TripDistance not expected")
	}

	if yr.RatecodeID != 1 {
		t.Errorf("RatecodeID not expected")
	}

	if yr.StoreAndFwdFlag != "N" {
		t.Errorf("StoreAndFwdFlag not expected")
	}

	if yr.PULocationID != 158 {
		t.Errorf("PULocationID not expected")
	}

	if yr.DOLocationID != 68 {
		t.Errorf("DOLocationID not expected")
	}

}

func TestValidateVendorId(t *testing.T) {
	/* Valid vendor Ids are 1 and 2 */

	t.Parallel()
	vendorIdTests := []struct {
		vid                 string
		valid               bool
		correctConvertedVID int
		shouldBeValid       bool
	}{
		{"1", true, 1, true},
		{"2", true, 2, true},
		{"foo", true, -1, false},
	}

	for _, test := range vendorIdTests {
		test := test
		testName := test.vid
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			vid := validateVendorId(&test.vid, &test.valid)

			if vid != test.correctConvertedVID || test.valid != test.shouldBeValid {
				t.Errorf("vid: %d != %d; %t != %t", vid, test.correctConvertedVID, test.valid, test.shouldBeValid)
			}
		})
	}
}

func TestPassengerCnt(t *testing.T) {
	/* Valid passenger cnts are between 1 and 4*/

	t.Parallel()
	passCntTests := []struct {
		pcnt          string
		valid         bool
		correctPCnt   int
		shouldBeValid bool
	}{
		{"1", true, 1, true},
		{"2", true, 2, true},
		{"0", true, -1, false},
		{"100", true, -1, false},
		{"foo", true, -1, false},
	}

	for _, test := range passCntTests {
		test := test
		testName := test.pcnt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			pcnt := validatePassengerCnt(&test.pcnt, &test.valid)

			if pcnt != test.correctPCnt || test.valid != test.shouldBeValid {
				t.Errorf("vid: %d != %d; %t != %t", pcnt, test.correctPCnt, test.valid, test.shouldBeValid)
			}
		})
	}
}

func TestValidateTripDistance(t *testing.T) {
	/* tripDistance should be a float in the range of (0.00, 100.00] */

	t.Parallel()
	tdTests := []struct {
		td                 string
		valid              bool
		correctConvertedTd float32
		shouldBeValid      bool
	}{
		{"1.00", true, 1.00, true},
		{"1.23", true, 1.23, true},
		{"foo", true, -1.0, false},
		{"101.00", true, -1.0, false},
		{"-1.00", true, -1.0, false},
		{"0.00", true, -1.0, false},
	}

	for _, test := range tdTests {
		test := test
		testName := test.td
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			convertedTd := validateTripDistance(&test.td, &test.valid)

			if convertedTd != test.correctConvertedTd || test.valid != test.shouldBeValid {
				t.Errorf("%s in should produce %f, not %f; valid should be %t, not %t",
					test.td, test.correctConvertedTd, convertedTd, test.shouldBeValid, test.valid)
			}
		})
	}
}

func TestValidateRatecodeID(t *testing.T) {
	/* RatecodeID is one of 1,2,3,4,5 or 6 according to data dictionary */

	t.Parallel()
	rcTests := []struct {
		rc                 string
		valid              bool
		correctConvertedRC int
		shouldBeValid      bool
	}{
		{"1", true, 1, true},
		{"2", true, 2, true},
		{"3", true, 3, true},
		{"4", true, 4, true},
		{"5", true, 5, true},
		{"6", true, 6, true},
		{"7", true, -1, false},
		{"0", true, -1, false},
		{"-2", true, -1, false},
		{"foo", true, -1, false},
	}

	for _, test := range rcTests {
		test := test
		testName := test.rc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			convertedRC := validateRatecodeID(&test.rc, &test.valid)

			if convertedRC != test.correctConvertedRC || test.valid != test.shouldBeValid {
				t.Errorf("%s in should produce %d, not %d; valid should be %t, not %t",
					test.rc, test.correctConvertedRC, convertedRC, test.shouldBeValid, test.valid)
			}
		})
	}
}

func TestValidateStoreAndFwdFlag(t *testing.T) {
	/* valid values are Y and N according to data dictionary */

	t.Parallel()
	sfTests := []struct {
		sf                 string
		valid              bool
		correctConvertedSF string
		shouldBeValid      bool
	}{
		{"Y", true, "Y", true},
		{"N", true, "N", true},
		{"foo", true, "Z", false},
	}

	for _, test := range sfTests {
		test := test
		testName := test.sf
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			convertedSF := validateStoreAndFwdFlag(&test.sf, &test.valid)

			if convertedSF != test.correctConvertedSF || test.valid != test.shouldBeValid {
				t.Errorf("%s in should produce %s, not %s; valid should be %t, not %t",
					test.sf, test.correctConvertedSF, convertedSF, test.shouldBeValid, test.valid)
			}
		})
	}
}

func TestLocationId(t *testing.T) {
	/* valid values are integers between 1 and 265 according to data dictionary */

	t.Parallel()
	puTests := []struct {
		lid                 string
		valid               bool
		correctConvertedLID int
		shouldBeValid       bool
	}{
		{"-12", true, -1, false},
		{"1", true, 1, true},
		{"4", true, 4, true},
		{"256", true, 256, true},
		{"foo", true, -1, false},
	}

	for _, test := range puTests {
		test := test
		testName := test.lid
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			convertedPULoc := validateLocationID(&test.lid, &test.valid)

			if convertedPULoc != test.correctConvertedLID || test.valid != test.shouldBeValid {
				t.Errorf("%s in should produce %d, not %d; valid should be %t, not %t",
					test.lid, test.correctConvertedLID, convertedPULoc, test.shouldBeValid, test.valid)
			}
		})
	}
}
