import type { StorageType } from "@entities/Storages/types/storages.type";
import type { Response } from "../response.type";

export type StorageResponse = Response & {
	data: StorageType;
};

export type StorageResponseMulti = Response & {
	data: StorageType[];
};
