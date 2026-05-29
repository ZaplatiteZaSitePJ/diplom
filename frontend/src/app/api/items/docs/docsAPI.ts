// entities/docs/docsApi.ts
import type { DocsFilter, DocsItem } from "@entities/Objects/types/docs.type";
import { baseApi } from "../../api";
import type { DocsResponse, DocsResponseMulti } from "./response.type";

export const docsApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// GET ALL
		getDocss: builder.query<DocsResponseMulti, DocsFilter>({
			query: (filter) => {
				console.log("RTK Query filter:", filter);

				return {
					url: "/admin/items/docs",
					method: "GET",
					params: filter, // 👈 сюда прокидывается объект
				};
			},
			providesTags: ["Docs"],
		}),

		getMyDocs: builder.query({
			query: () => ({
				url: "/me/items/docs",
				method: "GET",
			}),
		}),

		// GET BY ID
		getDocsById: builder.query<DocsResponse, string>({
			query: (id) => `/admin/items/docs/${id}`,
			providesTags: ["Docs"],
		}),

		// CREATE
		createDocs: builder.mutation<DocsResponse, Partial<DocsItem>>({
			query: (body) => ({
				url: "/admin/items/docs",
				method: "POST",
				body,
			}),
			invalidatesTags: ["Docs"],
		}),

		//PATCH
		updateDocs: builder.mutation<
			DocsResponse,
			{ id: string; body: Partial<DocsItem> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/items/docs/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["Docs"],
		}),

		// DELETE
		deleteDocs: builder.mutation<void, string>({
			query: (id) => ({
				url: `/admin/items/docs/${id}`,
				method: "DELETE",
			}),
			invalidatesTags: ["Docs"],
		}),
	}),
});

export const {
	useGetDocssQuery,
	useGetMyDocsQuery,
	useLazyGetDocsByIdQuery,
	useCreateDocsMutation,
	useDeleteDocsMutation,
	useUpdateDocsMutation,
} = docsApi;
