package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type team struct {
	name   string
	damage int
}

type person struct {
	name string
	dmg  int
	team string
}

func check(logger func(format string, v ...interface{}), e error) {
	if e != nil {
		logger(e.Error())
	}
}

func parseLine(text string) person {
	//Dmg only number
	damage := regexp.MustCompile("[0-9]+")
	damageStr := damage.FindString(text)
	damageNum, _ := strconv.Atoi(damageStr)
	charName := strings.Fields(text)

	return person{
		name: charName[1],
		dmg:  damageNum,
		team: "",
	}
}

func findMyTeam(whoami person) person {
	teamMap := make(map[string]string)

	teamMap["snak"] = "ev"
	teamMap["ruan'"] = "ev"
	teamMap["willson"] = "ev"
	teamMap["dogcompany"] = "ev"
	teamMap["jesskill"] = "ev"
	teamMap["wwwwwwwwwwwwwww"] = "ev"
	teamMap["spamkry"] = "ev"
	teamMap["wraken'"] = "ev"
	teamMap["ghenisha"] = "ev"
	teamMap["pea'shooter"] = "ev"
	teamMap["vanona"] = "ev"
	teamMap["anakaren"] = "ev"
	teamMap["mechalis"] = "ev"
	teamMap["fuji-kan"] = "ev"
	teamMap["wisha"] = "ev"
	teamMap["lelaebai"] = "ev"
	teamMap["ixin"] = "ev"
	teamMap["moonsnipe"] = "ev"
	teamMap["you"] = "ev"

	teamMap["asifab"] = "sem"
	teamMap["sylenth"] = "sem"
	teamMap["mareek"] = "sem"
	teamMap["innova"] = "sem"
	teamMap["kayros"] = "sem"
	teamMap["loubean"] = "sem"
	teamMap["thegreatunclean"] = "sem"
	teamMap["glenndanzig"] = "sem"
	teamMap["exact"] = "sem"
	teamMap["'aspect"] = "sem"
	teamMap["machru"] = "sem"
	teamMap["nemento"] = "sem"
	teamMap["oiz"] = "sem"
	teamMap["raae"] = "sem"
	teamMap["ra-"] = "sem"
	teamMap["lepen"] = "sem"
	teamMap["serato"] = "sem"
	teamMap["kayros"] = "sem"
	teamMap["juyn"] = "sem"
	teamMap["moosia"] = "sem"
	teamMap["silene"] = "sem"
	teamMap["seradek"] = "sem"
	teamMap["eekoo"] = "sem"

	teamMap["mynionss"] = "live"
	teamMap["kenshisan"] = "live"
	teamMap["melinoe"] = "live"
	teamMap["mynionss"] = "live"
	teamMap["lyinglow"] = "live"
	teamMap["adderall"] = "live"
	teamMap["lying"] = "live"
	teamMap["solamante"] = "live"
	teamMap["squabble"] = "live"
	teamMap["dieselrs"] = "live"
	teamMap["dapperdan"] = "live"
	teamMap["dapper'"] = "live"
	teamMap["valleypower"] = "live"
	teamMap["doggunner"] = "live"
	teamMap["variac"] = "live"
	teamMap["salfo"] = "live"
	teamMap["villion"] = "live"
	teamMap["tehmob"] = "live"
	teamMap["stilgarr"] = "live"

	whoami.team = teamMap[strings.ToLower(whoami.name)]
	return whoami
}

func totalTeamDamage(people []person) map[string]team {
	teams := make(map[string]team)
	for _, teamMem := range people {
		if currentTeam, ok := teams[teamMem.team]; ok {
			currentTeam.damage = currentTeam.damage + teamMem.dmg
			teams[teamMem.team] = currentTeam
		} else {
			teams[teamMem.team] = team{
				name:   teamMem.team,
				damage: teamMem.dmg,
			}
		}
		fmt.Println(teamMem)
	}
	return teams
}

func main() {
	f, err := os.Open("./new_chatlog.txt")
	logger := log.Fatalf
	check(logger, err)
	defer f.Close()
	allLines := []person{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		if !strings.Contains(string(line), "Acklay") {
			continue
		}

		acklayLine := string(line)

		character := parseLine(acklayLine)

		character = findMyTeam(character)

		allLines = append(allLines, character)
	}

	teamDmg := totalTeamDamage(allLines)

	fmt.Println(teamDmg)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
