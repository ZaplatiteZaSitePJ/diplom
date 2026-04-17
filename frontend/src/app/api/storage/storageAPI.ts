// entities/storage/storageApi.ts
import type { StorageType } from "@entities/Storages/types/storages.type";
import { baseApi } from "../api";
import type { StorageResponse, StorageResponseMulti } from "./response.type";

export const storageApi = baseApi.injectEndpoints({
	endpoints: (builder) => ({
		// GET ALL
		getStorages: builder.query<StorageResponseMulti, void>({
			query: () => "/admin/storages",
			providesTags: ["Storage"],
		}),

		// GET BY ID
		getStorageById: builder.query<StorageResponse, string>({
			query: (id) => `/admin/storages/${id}`,
			providesTags: ["Storage"],
		}),

		// CREATE
		createStorage: builder.mutation<StorageResponse, Partial<StorageType>>({
			query: (body) => ({
				url: "/admin/storages",
				method: "POST",
				body,
			}),
			invalidatesTags: ["Storage"],
		}),

		// PATCH
		updateStorage: builder.mutation<
			StorageResponse,
			{ id: string; body: Partial<StorageType> }
		>({
			query: ({ id, body }) => ({
				url: `/admin/storages/${id}`,
				method: "PATCH",
				body,
			}),
			invalidatesTags: ["Storage"],
		}),

		// DELETE
		deleteStorage: builder.mutation<
			void,
			{ id: string; newStorageName: string | undefined }
		>({
			query: ({ id, newStorageName }) => ({
				url: `/admin/storages/${id}`,
				method: "DELETE",
				params: newStorageName ? { newStorageName } : undefined,
			}),
			invalidatesTags: ["Storage"],
		}),
	}),
});

export const {
	useGetStoragesQuery,
	useGetStorageByIdQuery,
	useCreateStorageMutation,
	useDeleteStorageMutation,
	useUpdateStorageMutation,
} = storageApi;
