package main

/*
* periodically wake up 12:00-4:00 and scrap
* if scrap failed retry exponentially 4, 16, 128
* and report/alert error
 */

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"regexp"
	"strings"
	"time"
)

type point struct {
	tags   map[string]string
	fields map[string]interface{}
}

func mp(murl string) ([]point, error) {
	pts := []point{}

	dataLabels := map[string]bool{
		"Apartment":      true,
		"Sq. Ft.":        true,
		"Rent":           true,
		"Date Available": true,
	}
	//url := "https://www.missionpointebywindsor.com/availableunits.aspx?myOlePropertyId=422086&MoveInDate=&t=0.8235535301794832&floorPlans=1914340"
	url := "https://www.missionpointebywindsor.com/availableunits.aspx?myOlePropertyId=422086"
	resp, err := soup.Get(url)
	if err != nil {
		return pts, err
	}

	doc := soup.HTMLParse(resp)
	apts := doc.FindAll("td")
	scrapMap := map[string]map[string]string{}

	re := regexp.MustCompile("[0-9]+")
	for _, apt := range apts {
		attr := apt.Attrs()
		// find id
		id := re.FindAllString(attr["data-selenium-id"], 1)[0]
		if _, ok := scrapMap[id]; !ok {
			scrapMap[id] = map[string]string{"name": "MP"}
		}

		if key, ok := attr["data-label"]; ok {
			if _, ok := dataLabels[key]; ok {
				label := strings.Replace(apt.Text(), "#", "", 1)
				if len(label) <= 0 {
					availDate := apt.Find("span")
					label = availDate.Text()
				}
				scrapMap[id][key] = label
			}
		}
	}

	// normalize
	for _, v := range scrapMap {
		l := strings.Replace(strings.Split(v["Rent"], "-")[0], "$", "", 1)
		r := strings.Replace(l, ",", "", 1)
		pts = append(pts, point{
			tags: v,
			fields: map[string]interface{}{
				"rent-low": r},
		})
	}

	return pts, nil
}

func main() {
	db, err := DbInit()

	if err != nil {
		fmt.Errorf("failed to init %s \n", err)
		os.Exit(1)
	}

	for {
		p, err := mp()
		if err != nil {
			fmt.Errorf("failed to get points, %s", err)
			continue
		}
		for _, p1 := range p {
			fmt.Printf("points %+v \n", p1)
		}
		db.Write(p)
		time.Sleep(time.Hour * 12)

	}
}
