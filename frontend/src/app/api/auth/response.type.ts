import type { Response } from "../response.type";

type TokkensData = {
	access: string;
	role: "admin" | "user";
};

export type authResponse = Response & {
	data: TokkensData;
};
