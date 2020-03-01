package main

import (
	"testing"
	//"fmt"
)

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
