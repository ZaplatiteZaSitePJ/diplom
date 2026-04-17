import Auth from "@pages/auth/Auth";
// import Storages from "@pages/storages/ui/Storages";
// import StorageUnit from "@pages/storageUnit/StorageUnit";
import AuthLayout from "@pages/auth/AuthLayout";
import ItemUnit from "@pages/itemUnit/ItemUnit";
import Categories from "@pages/resources/Resources";
import Storages from "@pages/storages/Storages";
import StorageUnit from "@pages/storagesUnit/StorageUnit";
import MainLayout from "@shared/layouts/MainLayout";
import type { JSX } from "react";
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

const UnauthOnly = ({ children }: { children: JSX.Element }) => {
	const token = localStorage.getItem("access");

	if (token) {
		return <Navigate to="/" replace />;
	}

	return children;
};

const router = createBrowserRouter(
	createRoutesFromElements(
		<Route element={<Outlet />}>
			<Route element={<ProtectedLayout />}>
				<Route path="/" element={<MainLayout />}>
					{/* 👉 ВОТ ЭТО ДОБАВЛЯЕМ */}
					<Route index element={<Navigate to="items" replace />} />

					<Route path="items" element={<Categories />} />
					<Route path="items/:id" element={<ItemUnit />} />

					<Route path="storages">
						<Route index element={<Storages />} />
						<Route path=":id" element={<StorageUnit />} />
					</Route>
				</Route>
			</Route>

			<Route
				path="/auth"
				element={
					<UnauthOnly>
						<AuthLayout />
					</UnauthOnly>
				}
			>
				<Route index element={<Auth />} />
			</Route>
		</Route>,
	),
);

export default router;
