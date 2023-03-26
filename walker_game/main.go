// TODO: добавить замки

package main

import (
	"strings"
)

var (
	player                            Player
	kitchen, hall, bedroom, street    Room
	backpack                          Equipment
	keys, books, tea                  Item
	door                              Door
	kitchenTable, chair, bedroomTable Furniture
	rooms                             map[string]*Room
)

type Player struct {
	Location  *Room
	Equipment *Equipment // может быть только 1 рюкзак (или что-то подобное)
	Actions   map[string]interface{}
}

type Room struct { // кухня, коридор, улица и тд
	Name        string
	EnterMsg    string
	LookUpMsg   string
	MissionMsg  []string
	OutsideFlag bool
	Furniture   []*Furniture
	Doors       []*Door
	Ways        []*Room
}

type Furniture struct { // стол, стул, шкаф и тд
	Name      string
	Items     []*Item
	Equipment []*Equipment
}

type Equipment struct { // рюкзак
	Name      string
	Inventory map[string]*Item
}

type Item struct { // ключи и тд
	Name   string
	Usages map[string]*Door
}

type Door struct {
	Name            string
	ActivatedStatus bool
	Key             *Item
}

func main() {
	/*
		в этой функции можно ничего не писать
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
}

func (r *Room) initRoom(name, enterMsg, lookUpMsg string, missionMsg []string,
	ways []*Room, furnit []*Furniture, doors []*Door) {
	r.Name = name
	r.EnterMsg = enterMsg
	r.LookUpMsg = lookUpMsg
	r.MissionMsg = missionMsg
	r.Ways = ways
	r.Furniture = furnit
	r.Doors = doors
}

func (i *Item) initItem(name string, usages map[string]*Door) {
	i.Name = name
	i.Usages = usages
}

func (f *Furniture) initFurniture(name string, items []*Item, eqiup []*Equipment) {
	f.Name = name
	f.Items = items
	f.Equipment = eqiup
}

func (eq *Equipment) initEqiupment(name string, inventory map[string]*Item) {
	eq.Name = name
	eq.Inventory = inventory
}

func (p *Player) initPlayer(location *Room, equip *Equipment, actions map[string]interface{}) {
	p.Location = location
	p.Equipment = equip
	p.Actions = actions
}

func initGame() {
	hall.initRoom("коридор", "ничего интересного", "", nil,
		[]*Room{&kitchen, &bedroom, &street}, nil, []*Door{&door})

	kitchen.initRoom("кухня", "кухня, ничего интересного",
		"ты находишься на кухне", []string{"собрать рюкзак", "идти в универ"},
		[]*Room{&hall}, []*Furniture{&kitchenTable}, nil)

	bedroom.initRoom("комната", "ты в своей комнате", "", nil,
		[]*Room{&hall}, []*Furniture{&bedroomTable, &chair}, nil)

	street.initRoom("улица", "на улице весна", "", nil,
		[]*Room{&hall}, nil, []*Door{&door})
	street.OutsideFlag = true

	rooms = map[string]*Room{kitchen.Name: &kitchen,
		hall.Name: &hall, bedroom.Name: &bedroom, street.Name: &street}

	tea.initItem("чай", nil)
	keys.initItem("ключи", map[string]*Door{"дверь": &door})
	books.initItem("конспекты", nil)

	kitchenTable.initFurniture("стол", []*Item{&tea}, nil)
	bedroomTable.initFurniture("стол", []*Item{&keys, &books}, nil)
	chair.initFurniture("стул", nil, []*Equipment{&backpack})

	backpack.initEqiupment("рюкзак", nil)

	door.ActivatedStatus = false
	door.Name = "дверь"
	door.Key = &keys

	player.initPlayer(&kitchen, nil,
		map[string]interface{}{"осмотреться": player.LookUp,
			"идти": player.GoRoomString, "надеть": player.WearEquipString,
			"взять": player.TakeItemString, "применить": player.UseItemString})
}

func handleCommand(command string) string {
	var res = ""
	var commandFound bool

	commandSlice := strings.Split(command, " ")

	for actionName, action := range player.Actions {
		if actionName == commandSlice[0] {
			commandFound = true
			switch len(commandSlice) {
			case 1:
				res = action.(func() string)()
			case 2:
				res = action.(func(string) string)(commandSlice[1])
			case 3:
				res = action.(func(string, string) string)(commandSlice[1], commandSlice[2])
			}
		}
	}
	if commandFound {
		return res
	} else {
		return "неизвестная команда"
	}
}
