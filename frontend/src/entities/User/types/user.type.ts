export type UserType = {
	id: string;

	name: string;
	lastname: string;

	email: string;

	post: string;
	grade: "intern" | "junior" | "middle" | "senior" | "team lead" | "manager";

	city: string;
};
