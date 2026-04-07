import type { Response } from "../response.type";

type TokkensData = {
	access: string;
	refresh: string;
};

export type authResponse = Response & {
	data: TokkensData;
};
