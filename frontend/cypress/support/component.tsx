import { store } from "@app/api/store";
import "./commands";
import { mount } from "cypress/react";
import { MemoryRouter } from "react-router-dom";
import { Provider } from "react-redux";
import "../../src/app/styles/reset.css";
import "../../src/app/styles/fonts/fonts.css";
import "../../src/app/styles/variables.css";
import "../../src/app/styles/globalStyles.css";

Cypress.Commands.add("mount", (component, options = {}) => {
	return mount(
		<div id="root">
			<Provider store={store}>
				<MemoryRouter>{component}</MemoryRouter>
			</Provider>
		</div>,
		options,
	);
});

declare global {
	namespace Cypress {
		interface Chainable {
			mount: typeof mount;
		}
	}
}

export {};
