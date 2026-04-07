// entities/storage/storageApi.ts
import type { StorageType } from "@entities/Storages/types/storages.type";
import { baseApi } from "../api";
import type { UserType } from "@entities/User/types/user.type";
import type { authResponse } from "./response.type";

export const authApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// LOGIN
		login: builder.mutation<
			authResponse,
			Pick<UserType, "email" | "password">
		>({
			query: (body) => ({
				url: "/auth/login",
				method: "POST",
				body,
			}),
		}),

		// LOGOUT
		logout: builder.mutation<authResponse, Partial<StorageType>>({
			query: (body) => ({
				url: "/auth/logout/me",
				method: "POST",
				body,
			}),
		}),
	}),
});

export const { useLoginMutation } = authApi;
