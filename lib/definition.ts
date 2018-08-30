export type SectionBody = Map<string, string|string[]>;
export type Section = {
	name: string,
	order: number,
	dictionary: SectionBody,
};
