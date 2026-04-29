import { fetchBaseQuery } from "@reduxjs/toolkit/query";

import type {
	BaseQueryFn,
	FetchArgs,
	FetchBaseQueryError,
} from "@reduxjs/toolkit/query";
import { createApi } from "@reduxjs/toolkit/query/react";

// ================= BASE QUERY =================

const rawBaseQuery = fetchBaseQuery({
	baseUrl: "http://localhost:8080/api/v1",
	credentials: "include",
	prepareHeaders: (headers) => {
		const token = localStorage.getItem("access");

		if (token) {
			headers.set("Authorization", `Bearer ${token}`);
		}

		return headers;
	},
});

// ================= TYPES =================

type RefreshResponse = {
	access: string;
};

// ================= MUTEX =================

let isRefreshing = false;
let refreshPromise: Promise<RefreshResponse | void> | null = null;

// ================= BASE QUERY WITH REAUTH =================

const baseQueryWithAuthRedirect: BaseQueryFn<
	string | FetchArgs,
	unknown,
	FetchBaseQueryError
> = async (args, api, extraOptions) => {
	let result = await rawBaseQuery(args, api, extraOptions);

	if (result.error?.status === 401) {
		if (!isRefreshing) {
			isRefreshing = true;

			refreshPromise = rawBaseQuery(
				{
					url: "/auth/refresh",
					method: "POST",
				},
				api,
				extraOptions,
			)
				.then((res) => {
					if (res.data) {
						const data = res.data as { data: { access: string } };

						localStorage.setItem("access", data.data.access);

						return data;
					}

					throw new Error("refresh failed");
				})
				.catch(() => {
					// refresh умер → разлогиниваем
					localStorage.removeItem("access");
					window.location.href = "/auth";
				})
				.finally(() => {
					isRefreshing = false;
					refreshPromise = null;
				});
		}

		await refreshPromise;

		const newAccess = localStorage.getItem("access");

		if (!newAccess) {
			return result;
		}

		const retryArgs: FetchArgs =
			typeof args === "string" ? { url: args } : { ...args };

		retryArgs.headers = {
			...(retryArgs.headers || {}),
			Authorization: `Bearer ${newAccess}`,
		};

		result = await rawBaseQuery(retryArgs, api, extraOptions);
	}

	return result;
};

// ================= API =================

export const baseApi = createApi({
	reducerPath: "api",
	baseQuery: baseQueryWithAuthRedirect,
	refetchOnMountOrArgChange: true,
	tagTypes: ["Storage", "User", "Tech", "Docs", "Software"],
	endpoints: () => ({}),
});
