import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const rawBaseQuery = fetchBaseQuery({
	baseUrl: "http://localhost:8080/api/v1",
	prepareHeaders: (headers) => {
		const token = localStorage.getItem("access");

		if (token) {
			headers.set("Authorization", `Bearer ${token}`);
		}

		return headers;
	},
});

const baseQueryWithAuthRedirect = async (args, api, extraOptions) => {
	const result = await rawBaseQuery(args, api, extraOptions);

	if (result.error) {
		if (result.error.status === 401) {
			window.location.href = "/auth";
		}
	}

	return result;
};

export const baseApi = createApi({
	reducerPath: "api",
	baseQuery: baseQueryWithAuthRedirect,
	refetchOnMountOrArgChange: true,
	tagTypes: ["Storage", "User", "Tech"],
	endpoints: () => ({}),
});
