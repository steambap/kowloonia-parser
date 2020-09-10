export type Section = {
	name: string;
	order: number;
	lines: string[];
};

export interface EngineConfig extends Record<string, any> {
	sound: number;
	battle: number;
	drama: boolean;
	autoSave: boolean;
	// corp: string;
	stretch: boolean;
	campaign: number;
	host: string;
	textSpeed: number;
	formLeft: number;
	formTop: number;
	formMaximum: number;
};

export interface unitClass {
	name: string;
	animePath: string;
	meleeAttack: number;
	rangeAttack: number;
	initiative: number;
	iconIndex: number;
	huge: boolean;
	levelMelee: number;
	levelRange: number;
	levelInit: number;
	levelHP: number;
	levelSP: number;
	hp: number;
	meleeTarget: number;
	rangeTarget: number;
	unitCost: number;
	levelCost: number;
	unitSkill: string;
	unitSkillLevel: number;
	
}

export interface GameSetup extends Record<string, any> {
	mapWidth: number;
	mapHeight: number;
	startYear: number;
	startMonth: number;
	applicationTitle: string;
	yearPrefix: string;
	battleBGM: string;
	defaultBGM: string;
	aiBGM: string;
	tpLimit: number;
	spLimit: number;
	reformLimit: number;
	randomName1: string[];
	randomName2: string[];
	randomeName3: string[];
}
