// entities/tech/techApi.ts
import type { TechFilter, TechItem } from "@entities/Objects/types/tech.type";
import { baseApi } from "../../api";
import type { TechResponse, TechResponseMulti } from "./response.type";
import { type Response } from "@app/api/response.type";

export const techApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// GET ALL
		getTechs: builder.query<TechResponseMulti, TechFilter>({
			query: (filter) => {
				console.log("RTK Query filter:", filter);

				return {
					url: "/admin/items/tech",
					method: "GET",
					params: filter, // 👈 сюда прокидывается объект
				};
			},
			providesTags: ["Tech"],
		}),

		// GET BY ID
		getTechById: builder.query<TechResponse, string>({
			query: (id) => `/admin/items/tech/${id}`,
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

		//PATCH
		updateTech: builder.mutation<
			TechResponse,
			{ id: string; body: Partial<TechItem> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/items/tech/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["Tech"],
		}),

		// DELETE
		deleteTech: builder.mutation<void, string>({
			query: (id) => ({
				url: `/admin/items/${id}`,
				method: "DELETE",
			}),
			invalidatesTags: ["Tech"],
		}),

		getCategories: builder.query<Response & { data: string[] }, string>({
			query: (type_id) => `/admin/categories/${type_id}`,
		}),
	}),
});

export const {
	useGetTechsQuery,
	useLazyGetTechByIdQuery,
	useCreateTechMutation,
	useDeleteTechMutation,
	useUpdateTechMutation,
	useGetCategoriesQuery,
} = techApi;
