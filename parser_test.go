package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckErrorIfNil(t *testing.T) {
	var actualOutput string
	logFatalf := func(format string, args ...interface{}) {
		if len(args) > 0 {
			actualOutput = fmt.Sprintf(format, args)
		} else {
			actualOutput = format
		}
	}
	check(logFatalf, nil)

	if actualOutput != "" {
		t.Fatalf("Test Error")
	}
}

func TestErrorIsLogged(t *testing.T) {
	var actualOutput string
	logFatalf := func(format string, args ...interface{}) {
		if len(args) > 0 {
			actualOutput = fmt.Sprintf(format, args)
		} else {
			actualOutput = format
		}
	}
	errorCode := "Patrick"
	errors := errors.New(errorCode)

	check(logFatalf, errors)

	if actualOutput != errorCode {
		t.Fatalf("Test Error")
	}
}

func TestSnakHasDmg(t *testing.T) {
	lineItem := "[Combat] Snak uses Advanced Strafe on Acklay for 7328 points of damage!"

	people := parseLine(lineItem)

	assert.Equal(t, "Snak", people.name, "Should have snak as a person")
	assert.Equal(t, 7328, people.dmg, "Should have 7328 dmg")
}

func TestSnakHasATeam(t *testing.T) {
	snak := person{"Snak", 1000, ""}

	snakWithTeam := findMyTeam(snak)

	assert.Equal(t, "ev", snakWithTeam.team, "Should have snak as ev")
}

func TestAsifabHasATeam(t *testing.T) {
	asifab := person{"Asifab", 1000, ""}

	snakWithTeam := findMyTeam(asifab)

	assert.Equal(t, "sem", snakWithTeam.team, "Should have Asifab as SEM")
}

func TestRalphFromLiveHasATeam(t *testing.T) {
	liveTeamMember := person{"Kenshisan", 1000, ""}

	snakWithTeam := findMyTeam(liveTeamMember)

	assert.Equal(t, "live", snakWithTeam.team, "Should have Kenshisan as live")
}

func TestFindingTeamIsCaseInsensitive(t *testing.T) {
	asifab := person{"aSiFab", 1000, ""}

	snakWithTeam := findMyTeam(asifab)

	assert.Equal(t, "sem", snakWithTeam.team, "Should have Asifab as SEM")
}

func TestEVHasSumedDmg(t *testing.T) {
	snak := person{"snak", 1000, "ev"}
	wraken := person{"Wraken'", 3000, "ev"}

	teams := totalTeamDamage([]person{snak, wraken})

	assert.Equal(t, 4000, teams["ev"].damage, "Should have 4000 dmg")
}

func TestEVandSemHasSumedDmg(t *testing.T) {
	snak := person{"Snak", 1000, "ev"}
	wraken := person{"Wraken'", 3000, "ev"}
	asifab := person{"Asifab", 5000, "sem"}

	teams := totalTeamDamage([]person{snak, wraken, asifab})

	assert.Equal(t, 4000, teams["ev"].damage, "Should have 4000 dmg")
	assert.Equal(t, 5000, teams["sem"].damage, "Should have 5000 dmg")
}

func TestNoTeamHasSumedDmg(t *testing.T) {
	snak := person{"snak", 1000, ""}
	someone := person{"Wraken'", 3000, "ev"}
	wraken := person{"Wraken'", 3000, ""}

	teams := totalTeamDamage([]person{snak, someone, wraken})

	assert.Equal(t, 4000, teams[""].damage, "Should have 4000 dmg")
}

func TestTotalTeamDamageReturnsCountOfPeople(t *testing.T) {
	snak := person{"snak", 1000, "ev"}
	wraken := person{"Wraken'", 3000, "ev"}
	someone := person{"Someone", 3000, "ev"}

	teams := totalTeamDamage([]person{snak, wraken, someone})

	assert.Equal(t, 3, teams["ev"].count, "Should return total count of 2")
}

func TestTotalTeamDamageReturnsCountOfPeopleRegardlessOfDmgEntries(t *testing.T) {
	snak := person{"snak", 1000, "ev"}
	wraken := person{"Wraken'", 3000, "ev"}
	someone := person{"snak", 3000, "ev"}

	teams := totalTeamDamage([]person{snak, wraken, someone})

	assert.Equal(t, 2, teams["ev"].count, "Should return total count of 2")
}

func TestTotalTeamDamageReturnsUniqueCountByTeam(t *testing.T) {
	snak := person{"snak", 1000, "ev"}
	wraken := person{"Wraken'", 3000, "ev"}
	someone := person{"snak", 3000, "ev"}
	otherTeam := person{"joe", 3000, "ramrod"}

	teams := totalTeamDamage([]person{snak, wraken, someone, otherTeam})

	assert.Equal(t, 2, teams["ev"].count, "Should return total count of 2")
	assert.Equal(t, 1, teams["ramrod"].count, "Should return total count of 1")
}
