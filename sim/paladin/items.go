package paladin

import (
	"time"

	"github.com/wowsims/cata/sim/common/cata"
	"github.com/wowsims/cata/sim/core"
	"github.com/wowsims/cata/sim/core/stats"
)

// Tier 11 ret
var ItemSetReinforcedSapphiriumBattleplate = core.NewItemSet(core.ItemSet{
	Name: "Reinforced Sapphirium Battleplate",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			paladin.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Flat,
				ClassMask:  SpellMaskTemplarsVerdict,
				FloatValue: 0.1,
			})
		},
		4: func(agent core.Agent) {
			// Handled in inquisition.go
		},
	},
})

// Tier 12 ret
var ItemSetBattleplateOfImmolation = core.NewItemSet(core.ItemSet{
	Name: "Battleplate of Immolation",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()
			cata.RegisterIgniteEffect(&paladin.Unit, cata.IgniteConfig{
				ActionID:           core.ActionID{SpellID: 35395}.WithTag(3), // actual 99092
				DisableCastMetrics: true,
				DotAuraLabel:       "Flames of the Faithful",
				IncludeAuraDelay:   true,

				ProcTrigger: core.ProcTrigger{
					Name:           "Flames of the Faithful",
					Callback:       core.CallbackOnSpellHitDealt,
					ClassSpellMask: SpellMaskCrusaderStrike,
					Outcome:        core.OutcomeLanded,
				},

				DamageCalculator: func(result *core.SpellResult) float64 {
					return result.Damage * 0.15
				},
			})
		},
		4: func(agent core.Agent) {
			// Handled in talents_retribution.go
		},
	},
})

// Tier 13 ret
var ItemSetBattleplateOfRadiantGlory = core.NewItemSet(core.ItemSet{
	Name: "Battleplate of Radiant Glory",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in judgement.go
		},
		4: func(agent core.Agent) {
			// Handled in talents_retribution.go
		},
	},
})

// PvP set
var ItemSetGladiatorsVindication = core.NewItemSet(core.ItemSet{
	ID:   917,
	Name: "Gladiator's Vindication",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			paladin.AddStat(stats.Strength, 70)
		},
		4: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			paladin.AddStat(stats.Strength, 90)
			paladin.AddStaticMod(core.SpellModConfig{
				Kind:      core.SpellMod_Cooldown_Flat,
				ClassMask: SpellMaskJudgementBase,
				TimeValue: -1 * time.Second,
			})
		},
	},
})

func (paladin *Paladin) addBloodthirstyGloves() {
	switch paladin.Hands().ID {
	case 64844, 70649, 60414, 65591, 72379, 70250, 70488, 73707, 73570:
		paladin.AddStaticMod(core.SpellModConfig{
			Kind:       core.SpellMod_DamageDone_Pct,
			ClassMask:  SpellMaskCrusaderStrike,
			FloatValue: 0.05,
		})
	default:
		break
	}
}

// Tier 11 prot
var ItemSetReinforcedSapphiriumBattlearmor = core.NewItemSet(core.ItemSet{
	Name: "Reinforced Sapphirium Battlearmor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			paladin.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Flat,
				ClassMask:  SpellMaskCrusaderStrike,
				FloatValue: 0.1,
			})
		},
		4: func(agent core.Agent) {
			// Handled in guardian_of_ancient_kings.go
		},
	},
})

// Tier 12 prot
var ItemSetBattlearmorOfImmolation = core.NewItemSet(core.ItemSet{
	Name: "Battlearmor of Immolation",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			procDamage := 0.0

			righteousFlames := paladin.RegisterSpell(core.SpellConfig{
				ActionID:    core.ActionID{SpellID: 53600}.WithTag(3), // actual 99075
				SpellSchool: core.SpellSchoolFire,
				ProcMask:    core.ProcMaskEmpty,
				Flags: core.SpellFlagIgnoreModifiers |
					core.SpellFlagBinary |
					core.SpellFlagNoOnCastComplete |
					core.SpellFlagNoOnDamageDealt,

				DamageMultiplier: 1,
				ThreatMultiplier: 1,

				ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
					spell.CalcAndDealDamage(sim, target, procDamage, spell.OutcomeAlwaysHit)
				},
			})

			core.MakeProcTriggerAura(&paladin.Unit, core.ProcTrigger{
				Name:           "Righteous Flames",
				Callback:       core.CallbackOnSpellHitDealt,
				ClassSpellMask: SpellMaskShieldOfTheRighteous,
				Outcome:        core.OutcomeLanded,

				ProcChance: 1,

				Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					procDamage = result.Damage * 0.2
					righteousFlames.Cast(sim, result.Target)
				},
			})
		},
		4: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			paladin.FlamingAegis = paladin.GetOrRegisterAura(core.Aura{
				Label:    "Flaming Aegis",
				ActionID: core.ActionID{SpellID: 99090},
				Duration: time.Second * 10,

				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					paladin.PseudoStats.BaseParryChance += 0.12
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					paladin.PseudoStats.BaseParryChance -= 0.12
				},
			})

			// Trigger in divine_protection.go
		},
	},
})
