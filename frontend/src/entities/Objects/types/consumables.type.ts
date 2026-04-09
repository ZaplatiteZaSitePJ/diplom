import type { BaseObjectType } from "./baseObjects.type";

export type ConsumableItem = BaseObjectType & {
	title: string;
	quantity: number;
	unit: string;
};
