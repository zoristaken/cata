import{m as t,aa as e,a8 as a}from"./preset_utils-BOrpehHr.chunk.js";import{bI as o,A as i,S as l,a4 as c}from"./detailed_results-BnSf1ELV.chunk.js";const n=()=>e({fieldName:"summon",values:[{value:o.NoSummon,tooltip:"No Pet"},{actionId:i.fromSpellId(691),value:o.Felhunter},{actionId:i.fromSpellId(30146),value:o.Felguard,showWhen:t=>t.getSpec()==l.SpecDemonologyWarlock},{actionId:i.fromSpellId(688),value:o.Imp},{actionId:i.fromSpellId(712),value:o.Succubus}],changeEmitter:t=>t.changeEmitter}),s=()=>a({fieldName:"detonateSeed",label:"Detonate Seed on Cast",labelTooltip:"Simulates raid doing damage to targets such that seed detonates immediately on cast."}),m=t({fieldName:"prepullMastery",label:"Prepull Mastery",labelTooltip:"Mastery in the prepull set equipped at the start. Only applies if it's value is > 0 and only before combat."}),r=new Map([[c.StatSpellHaste,new Map([["13-tick - Agony",538],["6-tick - UA/Immo/DS",1280],["14-tick - Agony",1603],["4-tick - SF",2133],["15-tick - Agony",2665],["16-tick - Agony",3734],["7-tick - UA/Immo/DS",3844],["17-tick - Agony",4803],["18-tick - Agony",5869],["8-tick - UA/Immo/DS",6399],["5-tick - SF",6401],["19-tick - Agony",6934],["20-tick - Agony",8009],["9-tick - UA/Immo/DS",8967],["21-tick - Agony",9076],["22-tick - Agony",10134],["6-tick - SF",10681],["23-tick - Agony",11209],["10-tick - UA/Immo/DS",11533],["12-tick - Corruption",11735],["24-tick - Agony",12267],["25-tick - Agony",13342],["13-tick - Corruption",13883],["11-tick - UA/Immo/DS",14088]])]]);export{s as D,n as P,r as W,m as a};
