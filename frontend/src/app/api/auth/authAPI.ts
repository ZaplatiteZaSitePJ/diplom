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
			{
				email: string;
				password: string;
			}
		>({
			query: (body) => ({
				url: "/auth/login",
				method: "POST",
				body,
			}),
		}),

		// LOGOUT
		logout: builder.mutation<authResponse, void>({
			query: () => ({
				url: "/me/logout",
				method: "POST",
			}),
		}),
	}),
});

export const { useLoginMutation, useLogoutMutation } = authApi;
