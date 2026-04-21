import type { UserType } from "@entities/User/types/user.type";
import type { Response } from "../response.type";

export type UserResponse = Response & {
	data: UserType;
};

export type UserResponseMulti = Response & {
	data: UserType[];
};
