package fury

import (
	"time"

	"github.com/wowsims/cata/sim/core"
	"github.com/wowsims/cata/sim/warrior"
)

func (war *FuryWarrior) RegisterRagingBlow() {

	ragingBlowActionID := core.ActionID{SpellID: 85288}

	ohRagingBlow := war.RegisterSpell(core.SpellConfig{
		ActionID:       ragingBlowActionID.WithTag(2),
		SpellSchool:    core.SpellSchoolPhysical,
		ProcMask:       core.ProcMaskMeleeOHSpecial,
		ClassSpellMask: warrior.SpellMaskRagingBlowOh,
		Flags:          core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage | core.SpellFlagNoOnCastComplete,

		DamageMultiplier: 1.0,
		CritMultiplier:   war.DefaultMeleeCritMultiplier(),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			ohBaseDamage := spell.Unit.OHNormalizedWeaponDamage(sim, spell.MeleeAttackPower())

			spell.CalcAndDealDamage(sim, target, ohBaseDamage*war.EnrageEffectMultiplier, spell.OutcomeMeleeSpecialBlockAndCrit)
		},
	})

	war.RegisterSpell(core.SpellConfig{
		ActionID:       ragingBlowActionID.WithTag(1),
		SpellSchool:    core.SpellSchoolPhysical,
		ProcMask:       core.ProcMaskMeleeSpecial,
		Flags:          core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage | core.SpellFlagAPL,
		ClassSpellMask: warrior.SpellMaskRagingBlow,

		RageCost: core.RageCostOptions{
			Cost:   20,
			Refund: 0.8,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,

			CD: core.Cooldown{
				Timer:    war.NewTimer(),
				Duration: 6 * time.Second,
			},
		},

		DamageMultiplier: 1.0,
		CritMultiplier:   war.DefaultMeleeCritMultiplier(),

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return war.StanceMatches(warrior.BerserkerStance) && war.HasActiveAuraWithTag(warrior.EnrageTag) && war.HasOHWeapon()
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMeleeSpecialHit)

			if !result.Landed() {
				spell.IssueRefund(sim)
				return
			}
			//remove hit metric from the no damage hit roll
			spell.SpellMetrics[target.UnitIndex].Hits--

			// 1 hit roll then 2 damage events
			mhBaseDamage := spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower())
			spell.CalcAndDealDamage(sim, target, mhBaseDamage*war.EnrageEffectMultiplier, spell.OutcomeMeleeSpecialBlockAndCrit)

			ohRagingBlow.Cast(sim, result.Target)
		},
	})
}
