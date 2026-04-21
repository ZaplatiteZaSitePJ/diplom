import type { UserType } from "@entities/User/types/user.type";
import { baseApi } from "../api";
import type { UserResponse, UserResponseMulti } from "./response.type";

export const storageApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// =========================
		// GET ALL
		// =========================
		getUsers: builder.query<UserResponseMulti, Partial<UserType>>({
			query: (filter) => ({
				url: "/admin/users",
				method: "GET",
				params: filter,
			}),
			providesTags: ["User"],
		}),

		// =========================
		// GET BY ID
		// =========================
		getUserById: builder.query<UserResponse, string>({
			query: (id) => ({
				url: `/admin/users/${id}`,
				method: "GET",
			}),
			providesTags: ["User"],
		}),

		// =========================
		// CREATE
		// =========================
		createUser: builder.mutation<UserResponse, Partial<UserType>>({
			query: (body) => ({
				url: "/admin/users",
				method: "POST",
				body,
			}),
			invalidatesTags: ["User"],
		}),

		// =========================
		// UPDATE (PATCH)
		// =========================
		updateUser: builder.mutation<
			UserResponse,
			{ id: string; body: Partial<UserType> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/users/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["User"],
		}),

		// =========================
		// DELETE
		// =========================
		deleteUser: builder.mutation<
			void,
			{ id: string; newUserName?: string }
		>({
			query: ({ id, newUserName }) => ({
				url: `/admin/users/${id}`,
				method: "DELETE",
				params: newUserName ? { newUserName } : undefined,
			}),
			invalidatesTags: ["User"],
		}),
	}),
});

export const {
	useGetUsersQuery,
	useGetUserByIdQuery,
	useCreateUserMutation,
	useUpdateUserMutation,
	useDeleteUserMutation,
} = storageApi;
