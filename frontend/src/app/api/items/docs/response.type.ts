import type { Response } from "../../response.type";
import type { DocsItem } from "@entities/Objects/types/docs.type";

export type DocsResponse = Response & {
	data: DocsItem;
};

export type DocsResponseMulti = Response & {
	data: DocsItem[];
};
