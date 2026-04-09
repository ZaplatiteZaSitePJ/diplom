import type { TechItem } from "@entities/Objects/types/tech.type";

export type TechResponse = Response & {
	data: TechItem;
};

export type TechResponseMulti = Response & {
	data: TechItem[];
};
