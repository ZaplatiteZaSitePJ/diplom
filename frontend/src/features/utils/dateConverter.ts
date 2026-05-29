export const toISODate = (date?: string | Date | null) => {
	if (!date) return undefined;

	const d = new Date(date);

	// убираем timezone → берем локальную дату
	const offset = d.getTimezoneOffset();
	const localDate = new Date(d.getTime() - offset * 60000);

	return localDate.toISOString().split("T")[0]; // YYYY-MM-DD
};

export const fromISODate = (date?: string | null) => {
	if (!date) return undefined;

	// date = "2026-04-11"
	const d = new Date(date);

	// превращаем в ISO строку (UTC)
	return d.toISOString();
};

export const toDateOnly = (date: Date) => {
	return date.toISOString().split("T")[0];
};
