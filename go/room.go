package main

import "strings"

type Room struct {
	Name     string
	Monsters []Monster
	Item     Item
}

// NewRoom parses a string and returns a Room
func NewRoom(s string) Room {
	tokens := strings.Split(s, ",")
	if len(tokens) != 11 {
		panic("invalid room")
	}

	monsterHealth := parseIntAttribute(tokens[2])

	monster := Monster{
		CharacterAttributes: &CharacterAttributes{
			Name:         tokens[1],
			Health:       monsterHealth,
			AttackDamage: parseIntAttribute(tokens[3]),
			Speed:        parseIntAttribute(tokens[4]),
		},
		SpeedDamage:   parseIntAttribute(tokens[5]),
		Clonable:      parseBoolAttribute(tokens[6]),
		initialHealth: monsterHealth,
	}

	item := Item{
		CharacterAttributes: CharacterAttributes{
			Name:         tokens[7],
			Health:       parseIntAttribute(tokens[8]),
			AttackDamage: parseIntAttribute(tokens[9]),
			Speed:        parseIntAttribute(tokens[10]),
		},
	}

	return Room{
		Name:     tokens[0],
		Monsters: []Monster{monster},
		Item:     item,
	}
}

func (r *Room) AliveMonsters() []Monster {
	monsters := []Monster{}
	for _, m := range r.Monsters {
		if m.IsAlive() {
			monsters = append(monsters, m)
		}
	}

	return monsters
}

func (r *Room) Combat(h *Hero, m *Monster) {
	for {
		if !h.IsAlive() {
			tv.Show("💀 Hero " + h.Name + " dies!")
			break
		}
		if !m.IsAlive() {
			tv.Show("💀 Monster " + m.Name + " is dead")
			break
		}

		if h.Speed > m.Speed {
			tv.Show("🗡️ Hero " + h.Name + " fights " + m.Name)
			h.Hit(m)

			if m.IsAlive() {
				tv.Show("🧌 Monster " + m.Name + " attacks: " + m.Roar())
				m.Hit(h)
			}
		} else {
			tv.Show("🧌 Monster " + m.Name + " attacks: " + m.Roar())
			m.Hit(h)
			if h.IsAlive() {
				tv.Show("🗡️ Hero " + h.Name + " fights " + m.Name)
				h.Hit(m)
			}
		}

		if m.CanBeCloned() {
			cloned := m.clone()
			r.Monsters = append(r.Monsters, cloned)
			tv.Show("👥 Monster " + m.Name + " cloned itself!")
		}
	}
}