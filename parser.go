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
	count  int
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
	teamMap["iamalegend"] = "ev"
	teamMap["lykomedes"] = "ev"
	teamMap["netoo"] = "ev"
	teamMap["moeko"] = "ev"
	teamMap["okto"] = "ev"
	teamMap["brump"] = "ev"
	teamMap["wolfsheen"] = "ev"
	teamMap["ussoc"] = "ev"
	teamMap["thecolorred"] = "ev"
	teamMap["thecolorred's"] = "ev"
	teamMap["fujikano"] = "ev"

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
	teamMap["aspect"] = "sem"
	teamMap["adrell"] = "sem"
	teamMap["innova"] = "sem"
	teamMap["debree"] = "sem"
	teamMap["prusa"] = "sem"
	teamMap["morbol"] = "sem"
	teamMap["sirbeanie"] = "sem"
	teamMap["rubedaddy"] = "sem"
	teamMap["daddyap"] = "sem"
	teamMap["z'oso"] = "sem"
	teamMap["wololo"] = "sem"
	teamMap["radial"] = "sem"
	teamMap["'anita"] = "sem"
	teamMap["mallicoy"] = "sem"
	teamMap["rand-e"] = "sem"
	teamMap["tecate"] = "sem"
	teamMap["'archer"] = "sem"
	teamMap["kolera"] = "sem"

	teamMap["mynionsss"] = "live"
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
	teamMap["salfo"] = "live"
	teamMap["villion"] = "live"
	teamMap["tehmob"] = "live"
	teamMap["adderall"] = "live"
	teamMap["stilgarr"] = "live"
	teamMap["theopaway"] = "live"
	teamMap["lusitania"] = "live"
	teamMap["taliyah"] = "live"
	teamMap["salfo"] = "live"
	teamMap["ziyad"] = "live"
	teamMap["variac"] = "live"
	teamMap["kalpo"] = "live"
	teamMap["insignificance"] = "live"
	teamMap["fisile"] = "live"
	teamMap["behelds"] = "live"
	teamMap["notamea"] = "live"
	teamMap["adrelle"] = "live"
	teamMap["juzzy"] = "live"
	teamMap["xtz"] = "live"
	teamMap["pouhaa"] = "live"
	teamMap["kura"] = "live"
	teamMap["kraytpolice"] = "live"
	teamMap["psychopath"] = "live"

	whoami.team = teamMap[strings.ToLower(whoami.name)]
	return whoami
}

func totalTeamDamage(people []person) map[string]team {
	membersSeenBefore := make(map[string]int)
	teams := make(map[string]team)
	for _, teamMem := range people {
		timesSeen := membersSeenBefore[teamMem.name]
		membersSeenBefore[teamMem.name] = timesSeen + 1
		if currentTeam, ok := teams[teamMem.team]; ok {
			currentTeam.damage = currentTeam.damage + teamMem.dmg
			if timesSeen == 0 {
				currentTeam.count = currentTeam.count + 1
			}
			teams[teamMem.team] = currentTeam
		} else {
			teams[teamMem.team] = team{
				name:   teamMem.team,
				damage: teamMem.dmg,
				count:  1,
			}
		}

		if teamMem.team == "" {
			fmt.Println(teamMem.name)
		}
	}
	return teams
}

func main() {
	f, err := os.Open("./6320.txt")
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
