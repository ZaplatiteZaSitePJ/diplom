export const normalize = <T>(v: T | undefined | null) => {
	if (v === "" || v === undefined || v === null) return undefined;
	return v;
};
