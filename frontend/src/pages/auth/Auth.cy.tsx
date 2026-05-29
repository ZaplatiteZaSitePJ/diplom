import Auth from "./Auth";

describe("Auth page", () => {
	it("wrong email", () => {
		cy.mount(<Auth />);
		cy.get('[name="email"]').type("admin@vcompany.com");
		cy.get('[name="password"]').type("StrongPass123!");
		cy.get("button").click();
	});

	it("wrong password", () => {
		cy.mount(<Auth />);
		cy.get('[name="email"]').type("admin@company.com");
		cy.get('[name="password"]').type("StrogngPass123!");
		cy.get("button").click();
	});

	it("correct data", () => {
		cy.mount(<Auth />);
		cy.get('[name="email"]').type("admin@company.com");
		cy.get('[name="password"]').type("StrongPass123!");
		cy.get("button").click();
	});
});
