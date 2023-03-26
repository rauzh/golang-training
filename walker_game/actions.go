package main

import "fmt"

func (p *Player) GoRoomString(destRoomName string) string {
	var result = ""
	for roomName, room := range rooms {
		if roomName == destRoomName {
			result += p.GoRoom(*room)
		}
	}
	if result == "" { // такой комнаты не существует
		result = fmt.Sprintf("нет пути в %s", destRoomName)
	}

	return result
}

func (p *Player) handleDoor(destRoom Room) string {
	var result = ""
	for _, door := range p.Location.Doors {
		for _, destDoor := range destRoom.Doors {
			if destDoor == door {
				if door.ActivatedStatus {
					p.Location = &destRoom
					result += p.Location.EnterMsg
					result += PrintWays(*p.Location)
				} else {
					return "дверь закрыта"
				}
			}
		}
	}

	return result
}

func (p *Player) GoRoom(destRoom Room) (result string) {
	result = ""

	for _, way := range p.Location.Ways {
		if destRoom.Name == way.Name {
			if p.Location.Doors != nil && destRoom.Doors != nil {
				result += p.handleDoor(destRoom)
				if result == "дверь закрыта" {
					return
				}
			} else {
				p.Location = &destRoom
				result += p.Location.EnterMsg
				result += PrintWays(*p.Location)
			}
		}
	}
	if p.Location.Name != destRoom.Name { // нет пути из этой комнаты в destRoom
		result = fmt.Sprintf("нет пути в %s", destRoom.Name)
	}

	return
}

func (p *Player) LookUp() string {
	var result = ""
	if p.Location.LookUpMsg != "" {
		result += fmt.Sprintf("%s, ", p.Location.LookUpMsg)
	}

	emptyFlag := true

	for furnitIdx, furnit := range p.Location.Furniture {
		if furnitIdx == 0 && (len(furnit.Items) > 0 || len(furnit.Equipment) > 0) {
			result += fmt.Sprintf("на %sе: ", furnit.Name)
			emptyFlag = false
		} else if len(furnit.Items) > 0 || len(furnit.Equipment) > 0 {
			result += fmt.Sprintf(", на %sе: ", furnit.Name)
		}

		result += PrintItems(*furnit)
		result += PrintEquipment(*furnit)
	}

	if emptyFlag {
		result += "пустая комната"
	}

	if p.Equipment != nil && len(p.Location.MissionMsg) == 2 {
		p.Location.MissionMsg = []string{p.Location.MissionMsg[1]}
	}

	result += PrintMission(p.Location.MissionMsg)
	result += PrintWays(*p.Location)

	return result
}

func (p *Player) WearEquipString(equipName string) string {
	var result = ""
	for _, furnit := range p.Location.Furniture {
		for equipIdx, equip := range furnit.Equipment {
			if equipName == equip.Name {
				result += p.WearEquip(equip)
				furnit.Equipment = append(furnit.Equipment[:equipIdx],
					furnit.Equipment[equipIdx+1:]...)
			}
		}
	}
	if result == "" { // нет такого предмета в комнате
		result = fmt.Sprintf("нет предмета - %s", equipName)
	}

	return result
}

func (p *Player) WearEquip(equip *Equipment) string {
	p.Equipment = equip

	return fmt.Sprintf("вы надели: %s", equip.Name)
}

func (p *Player) TakeItemString(itemName string) string {
	var result = ""
	for _, furnit := range p.Location.Furniture {
		for itemIdx, item := range furnit.Items {
			if itemName == item.Name {
				result += p.TakeItem(item)
				if result != "некуда класть" {
					furnit.Items = append(furnit.Items[:itemIdx],
						furnit.Items[itemIdx+1:]...)
				}
			}
		}
	}
	if result == "" { // нет такого предмета в комнате
		result = "нет такого"
	}

	return result
}

func (p *Player) TakeItem(item *Item) string {
	if p.Equipment == nil {
		return "некуда класть"
	} else {
		if p.Equipment.Inventory != nil {
			p.Equipment.Inventory[item.Name] = item
		} else {
			p.Equipment.Inventory = map[string]*Item{item.Name: item}
		}
	}

	return fmt.Sprintf("предмет добавлен в инвентарь: %s", item.Name)
}

func (p *Player) UseItemString(itemName string, doorName string) string {
	var result = ""
	var itemFound, doorFound bool
	var item *Item
	if p.Equipment == nil || p.Equipment.Inventory == nil {
		return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
	}

	for playerItemName, playerItem := range p.Equipment.Inventory {
		if playerItemName == itemName {
			itemFound = true
			item = playerItem
		}
	}

	if !itemFound {
		return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
	}

	for _, locationDoor := range p.Location.Doors {
		if locationDoor.Name == doorName {
			doorFound = true
			if locationDoor.Key == item {
				locationDoor.ActivatedStatus = true
				return fmt.Sprintf("%s открыта", locationDoor.Name)
			}
		}
	}

	if !doorFound {
		return "не к чему применить"
	}

	return result
}
