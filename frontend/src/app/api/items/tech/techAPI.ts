// entities/tech/techApi.ts
import type { TechItem } from "@entities/Objects/types/tech.type";
import { baseApi } from "../../api";
import type { TechResponse, TechResponseMulti } from "./response.type";

export const techApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// GET ALL
		getTechs: builder.query<TechResponseMulti, void>({
			query: () => "/admin/items/tech",
			providesTags: ["Tech"],
		}),

		// GET BY ID
		getTechById: builder.query<TechResponse, string>({
			query: (id) => `/admin/tech/${id}`,
			providesTags: ["Tech"],
		}),

		// CREATE
		createTech: builder.mutation<TechResponse, Partial<TechItem>>({
			query: (body) => ({
				url: "/admin/items/tech",
				method: "POST",
				body,
			}),
			invalidatesTags: ["Tech"],
		}),

		updateTech: builder.mutation<
			TechResponse,
			{ id: string; body: Partial<TechItem> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/tech/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["Tech"],
		}),

		// DELETE
		deleteTech: builder.mutation<void, string>({
			query: (id) => ({
				url: `/admin/tech/${id}`,
				method: "DELETE",
			}),
			invalidatesTags: ["Tech"],
		}),
	}),
});

export const {
	useGetTechsQuery,
	useGetTechByIdQuery,
	useCreateTechMutation,
	useDeleteTechMutation,
	useUpdateTechMutation,
} = techApi;
