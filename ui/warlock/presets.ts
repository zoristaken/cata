import { Stat } from '../core/proto/common';

export const WARLOCK_BREAKPOINTS = new Map([
	[
		Stat.StatHasteRating,
		new Map([
			// Picked from Warlock Discord
			// Sources:
			// https://docs.google.com/spreadsheets/d/1plPkROOiRxDO3lBNc6ejowUgXDrt_zXd7FaYPU1wyVQ/edit?gid=497073075#gid=497073075
			['13-tick - Agony', 4.19381],
			['7-tick - Corruption', 8.32281],
			['6-tick - UA/Immo/DS', 9.99084],
			['14-tick - Agony', 12.51759],
			['4-tick - SF', 16.65209],
			['15-tick - Agony', 20.80943],
			['8-tick - Corruption', 24.97397],
			['16-tick - Agony', 29.15726],
			['7-tick - UA/Immo/DS', 30.01084],
			['17-tick - Agony', 37.50431],
			['9-tick - Corruption', 41.67651],
			['18-tick - Agony', 45.82575],
			['8-tick - UA/Immo/DS', 49.96252],
			['5-tick - SF', 49.98126],
			['19-tick - Agony', 54.14259],
			['10-tick - Corruption', 58.35314],
			['20-tick - Agony', 62.53557],
			['9-tick - UA/Immo/DS', 70.01985],
			['21-tick - Agony', 70.86717],
			['11-tick - Corruption', 74.97814],
			['22-tick - Agony', 79.13123],
			['6-tick - SF', 83.40213],
			['23-tick - Agony', 87.52932],
			['10-tick - UA/Immo/DS', 90.05386],
			['12-tick - Corruption', 91.63208],
			['24-tick - Agony', 95.79052],
			['25-tick - Agony', 104.18583],
			['13-tick - Corruption', 108.40571],
			['11-tick - UA/Immo/DS', 110.01052],
			// ['26-tick - Agony', 112.427],
			// ['7-tick - SF', 116.56743],
			// ['27-tick - Agony', 120.87247],
			// ['14-tick - Corruption', 124.9719],
			// ['28-tick - Agony', 129.22639],
			// ['12-tick - UA/Immo/DS', 129.97319],
			// ['29-tick - Agony', 137.38875],
			// ['15-tick - Corruption', 141.64319],
			// ['30-tick - Agony', 145.85129],
			// ['8-tick - SF', 149.84388],
			// ['13-tick - UA/Immo/DS', 150.10423],
			// ['31-tick - Agony', 154.2912],
			// ['16-tick - Corruption', 158.28672],
			// ['32-tick - Agony', 162.63956],
			// ['14-tick - UA/Immo/DS', 169.90556],
			// ['33-tick - Agony', 170.81926],
			// ['17-tick - Corruption', 175.10319],
			// ['34-tick - Agony', 179.13472],
			// ['9-tick - SF', 183.48693],
			// ['35-tick - Agony', 187.56295],
			// ['15-tick - UA/Immo/DS', 189.99519],
			// ['36-tick - Agony', 195.63936],
			// ['37-tick - Agony', 204.18256],
			// ['38-tick - Agony', 212.2561],
			// ['39-tick - Agony', 220.7699],
			// ['40-tick - Agony', 229.21816],
			// ['41-tick - Agony', 237.5528],
			// ['42-tick - Agony', 245.72175],
			// ['43-tick - Agony', 254.2959],
			// ['44-tick - Agony', 262.64739],
			// ['45-tick - Agony', 270.71369],
			// ['46-tick - Agony', 279.14699],
			// ['47-tick - Agony', 287.22176],
		]),
	],
]);
