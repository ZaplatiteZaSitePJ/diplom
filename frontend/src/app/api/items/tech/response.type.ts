import type { TechItem } from "@entities/Objects/types/tech.type";
import type { Response } from "../../response.type";

export type TechResponse = Response & {
	data: TechItem;
};

export type TechResponseMulti = Response & {
	data: TechItem[];
};
