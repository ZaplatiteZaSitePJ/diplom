import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { Provider } from "react-redux";
import { store } from "./api/store.ts";
import { SnackbarProvider } from "notistack";

createRoot(document.getElementById("root")!).render(
	<StrictMode>
		<SnackbarProvider
			maxSnack={3}
			anchorOrigin={{
				vertical: "bottom",
				horizontal: "center",
			}}
		>
			<Provider store={store}>
				<App />
			</Provider>
		</SnackbarProvider>
	</StrictMode>,
);
