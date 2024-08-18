package main

type CooldownSchema struct {
	Total_Seconds     int    `json:"total_seconds"`
	Remaining_Seconds int    `json:"remaining_seconds"`
	Reason            string `json:"reason"`
}

type SkillInfoSchema struct {
	Xp int `json:"xp"`
}

type Character struct {
	Name                      string          `json:"name"`
	Skin                      string          `json:"skin"`
	Level                     int             `json:"level"`
	Xp                        int             `json:"xp"`
	Max_Xp                    int             `json:"max_xp"`
	Total_Xp                  int             `json:"total_xp"`
	Gold                      int             `json:"gold"`
	Speed                     int             `json:"speed"`
	Mining_Level              int             `json:"mining_level"`
	Mining_Xp                 int             `json:"mining_xp"`
	Mining_Max_Xp             int             `json:"mining_max_xp"`
	Woodcutting_Level         int             `json:"woodcutting_level"`
	Woodcutting_Xp            int             `json:"woodcutting_xp"`
	Woodcutting_Max_Xp        int             `json:"woodcutting_max_xp"`
	Fishing_Level             int             `json:"fishing_level"`
	Fishing_Xp                int             `json:"fishing_xp"`
	Fishing_Max_Xp            int             `json:"fishing_max_xp"`
	Weaponcrafting_Level      int             `json:"weaponcrafting_level"`
	Weaponcrafting_Xp         int             `json:"weaponcrafting_xp"`
	Weaponcrafting_Max_Xp     int             `json:"weaponcrafting_max_xp"`
	Gearcrafting_Level        int             `json:"gearcrafting_level"`
	Gearcrafting_Xp           int             `json:"gearcrafting_xp"`
	Gearcrafting_Max_Xp       int             `json:"gearcrafting_max_xp"`
	Jewelrycrafting_Level     int             `json:"jewelrycrafting_level"`
	Jewelrycrafting_Xp        int             `json:"jewelrycrafting_xp"`
	Jewelrycrafting_Max_Xp    int             `json:"jewelrycrafting_max_xp"`
	Cooking_Level             int             `json:"cooking_level"`
	Cooking_Xp                int             `json:"cooking_xp"`
	Cooking_Max_Xp            int             `json:"cooking_max_xp"`
	Hp                        int             `json:"hp"`
	Haste                     int             `json:"haste"`
	Critical_Strike           int             `json:"critical_strike"`
	Stamina                   int             `json:"stamina"`
	Attack_Fire               int             `json:"attack_fire"`
	Attack_Earth              int             `json:"attack_earth"`
	Attack_Water              int             `json:"attack_water"`
	Attack_Air                int             `json:"attack_air"`
	Dmg_Fire                  int             `json:"dmg_fire"`
	Dmg_Earth                 int             `json:"dmg_earth"`
	Dmg_Water                 int             `json:"dmg_water"`
	Dmg_Air                   int             `json:"dmg_air"`
	Res_Fire                  int             `json:"res_fire"`
	Res_Earth                 int             `json:"res_earth"`
	Res_Water                 int             `json:"res_water"`
	Res_Air                   int             `json:"res_air"`
	X                         int             `json:"x"`
	Y                         int             `json:"Y"`
	Cooldown                  int             `json:"cooldown"`
	Cooldown_Expiration       int             `json:"cooldown_expiration"`
	Weapon_Slot               string          `json:"weapon_slot"`
	Shield_Slot               string          `json:"shield_slot"`
	Helmet_Slot               string          `json:"helmet_slot"`
	Body_Armor_Slot           string          `json:"body_armor_slot"`
	Leg_Armor_Slot            string          `json:"leg_armor_slot"`
	Boots_Slot                string          `json:"boots_slot"`
	Ring1_Slot                string          `json:"ring1_slot"`
	Ring2_Slot                string          `json:"ring2_slot"`
	Amulet_Slot               string          `json:"amulet_slot"`
	Artifact1_Slot            string          `json:"artifact1_1lot"`
	Artifact2_Slot            string          `json:"artifact2_slot"`
	Artifact3_Slot            string          `json:"artifact3_slot"`
	Consumable1_Slot          string          `json:"consumable1_slot"`
	Consumable1_Slot_Quantity int             `json:"consumable1_slot_quantity"`
	Consumable2_Slot          string          `json:"consumable2_slot"`
	Consuable2_Slot_Quantity  int             `json:"consumable2_slot_quantity"`
	Task                      string          `json:"task"`
	Task_Type                 string          `json:"task_type"`
	Task_Progress             int             `json:"task_progress"`
	Task_Total                int             `json:"task_total"`
	Inventory_Max_Items       int             `json:"inventory_max_items"`
	Inventory                 []InventorySlot `json:"inventory"`
}

type InventorySlot struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type SkillDataSchema struct {
	Data struct {
		Cooldown  CooldownSchema  `json:"cooldown"`
		Details   SkillInfoSchema `json:"details"`
		Character Character       `json:"character"`
	} `json:"data"`
}

type CharacterSchema struct {
	Data struct {
		Character Character `json:"character"`
	} `json:"data"`
}
