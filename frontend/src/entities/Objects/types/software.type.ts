import type { BaseObjectType } from "./baseObjects.type";

export type SoftwareItem = BaseObjectType & {
	vendor: string;
	license_key: string;
	title: string;
	responsible_worker: string;
	started_at: string;
	expired_at: string;
	updated_at: string;
};
