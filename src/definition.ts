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
