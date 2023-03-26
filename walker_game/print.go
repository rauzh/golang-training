package main

import "fmt"

func PrintWays(r Room) string {
	var result = ""

	if r.OutsideFlag {
		return ". можно пройти - домой"
	}

	if r.Ways != nil && len(r.Ways) > 0 {
		result += ". можно пройти - "
	}
	for wayIdx, way := range r.Ways {
		if wayIdx == 0 {
			result += way.Name
		} else {
			result += fmt.Sprintf(", %s", way.Name)
		}
	}

	return result
}

func PrintMission(misList []string) (result string) {
	result = ""
	if misList != nil {
		result += ", надо "
		for missionIdx, mission := range misList {
			if missionIdx == 0 {
				result += mission
			} else {
				result += fmt.Sprintf(" и %s", mission)
			}
		}
	}

	return
}

func PrintItems(furnit Furniture) (result string) {
	result = ""
	for itemIdx, item := range furnit.Items {
		if itemIdx == 0 {
			result += item.Name
		} else {
			result += fmt.Sprintf(", %s", item.Name)
		}
	}

	return
}

func PrintEquipment(furnit Furniture) (result string) {
	result = ""
	for equipIdx, equip := range furnit.Equipment {
		if equipIdx == 0 {
			result += equip.Name
		} else {
			result += fmt.Sprintf(", %s", equip.Name)
		}
	}

	return
}
