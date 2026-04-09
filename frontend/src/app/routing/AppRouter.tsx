import Auth from "@pages/auth/Auth";
// import Storages from "@pages/storages/ui/Storages";
// import StorageUnit from "@pages/storageUnit/StorageUnit";
import AuthLayout from "@pages/auth/AuthLayout";
import Categories from "@pages/resources/Resources";
import Storages from "@pages/storages/Storages";
import StorageUnit from "@pages/storagesUnit/StorageUnit";
import MainLayout from "@shared/layouts/MainLayout";
// import MainLayout from "@shared/ui/layouts/MainLayout";
import {
	createBrowserRouter,
	createRoutesFromElements,
	Navigate,
	// Navigate,
	Outlet,
	// redirect,
	Route,
} from "react-router-dom";

const ProtectedLayout = () => {
	const token = localStorage.getItem("access");

	if (!token) {
		return <Navigate to="/auth" replace />;
	}

	return <Outlet />;
};

// const UnauthOnly = ({ children }: { children: JSX.Element }) => {
// 	const token = localStorage.getItem("accessToken");

// 	if (token) {
// 		return <Navigate to="/" replace />;
// 	}

// 	return children;
// };

const router = createBrowserRouter(
	createRoutesFromElements(
		<Route element={<Outlet />}>
			{/* <Route element={<ProtectedLayout />}> */}
			{/* <Route path="/" element={<MainLayout />}>
				<Route
					index
					element={<></>}
					loader={() => redirect("/storages")}
				/>

				<Route path="/storages/">
					<Route index element={<Storages />} />
					<Route path="/storages/:id" element={<StorageUnit />} />
				</Route>
			</Route> */}
			<Route element={<ProtectedLayout />}>
				<Route path="/" element={<MainLayout />}>
					<Route index element={<Categories />} />

					<Route path="/storages/">
						<Route index element={<Storages />} />
						<Route
							path="/storages/:id"
							element={<StorageUnit />}
						></Route>
					</Route>
				</Route>
			</Route>

			<Route
				path="/auth"
				element={
					// <UnauthOnly>
					<AuthLayout />
					// </UnauthOnly>
				}
			>
				<Route index element={<Auth />} />
			</Route>
		</Route>,
	),
);

export default router;
