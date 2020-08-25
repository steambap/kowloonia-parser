export type SectionBody = Map<string, string | string[]>;
export type Section = {
	name: string;
	order: number;
	dictionary: SectionBody;
};

export type EngineConfig = {
	sound: number;
	battle: number;
	drama: boolean;
	autosave: boolean;
	corp: string;
	stretch: boolean;
	campaign: number;
	host: string;
	textSpeed: number;
	formLeft: number;
	formTop: number;
	formMaximum: number;
};
