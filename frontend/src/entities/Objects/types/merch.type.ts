import type { BaseObjectType } from "./baseObjects.type";

export type MerchItem = BaseObjectType & {
	title: string;
	size: string;
	price: number;
	color: string;
};
