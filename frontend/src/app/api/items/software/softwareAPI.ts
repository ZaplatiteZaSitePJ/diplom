// entities/software/softwareApi.ts
import type {
	SoftwareFilter,
	SoftwareItemPublic,
} from "@entities/Objects/types/software.type";
import { baseApi } from "../../api";
import type { SoftwareResponse, SoftwareResponseMulti } from "./response.type";

export const softwareApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// GET ALL
		getSoftwares: builder.query<SoftwareResponseMulti, SoftwareFilter>({
			query: (filter) => {
				console.log("RTK Query filter:", filter);

				return {
					url: "/admin/items/software",
					method: "GET",
					params: filter, // 👈 сюда прокидывается объект
				};
			},
			providesTags: ["Software"],
		}),

		getMySoftwares: builder.query({
			query: () => ({
				url: "/me/items/software",
				method: "GET",
			}),
		}),

		// GET BY ID
		getSoftwareById: builder.query<SoftwareResponse, string>({
			query: (id) => `/admin/items/software/${id}`,
			providesTags: ["Software"],
		}),

		// CREATE
		createSoftware: builder.mutation<
			SoftwareResponse,
			Partial<SoftwareItemPublic>
		>({
			query: (body) => ({
				url: "/admin/items/software",
				method: "POST",
				body,
			}),
			invalidatesTags: ["Software"],
		}),

		//PATCH
		updateSoftware: builder.mutation<
			SoftwareResponse,
			{ id: string; body: Partial<SoftwareItemPublic> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/items/software/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["Software"],
		}),

		// DELETE
		deleteSoftware: builder.mutation<void, string>({
			query: (id) => ({
				url: `/admin/items/software/${id}`,
				method: "DELETE",
			}),
			invalidatesTags: ["Software"],
		}),
	}),
});

export const {
	useGetSoftwaresQuery,
	useLazyGetSoftwareByIdQuery,
	useCreateSoftwareMutation,
	useDeleteSoftwareMutation,
	useUpdateSoftwareMutation,
	useGetMySoftwaresQuery,
} = softwareApi;
