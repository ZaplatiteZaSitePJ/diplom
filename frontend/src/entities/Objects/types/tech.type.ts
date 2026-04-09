import type { BaseObjectType } from "./baseObjects.type";

export type TechItem = BaseObjectType & {
	brand: string;
	model: string;
	warranty_started_at: string;
	warranty_end_at: string;
};
