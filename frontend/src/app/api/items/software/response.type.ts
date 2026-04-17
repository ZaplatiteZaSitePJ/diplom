import type { SoftwareItemPublic } from "@entities/Objects/types/software.type";
import type { Response } from "../../response.type";

export type SoftwareResponse = Response & {
	data: SoftwareItemPublic;
};

export type SoftwareResponseMulti = Response & {
	data: SoftwareItemPublic[];
};
