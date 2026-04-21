import type { BaseObjectType } from "./baseObjects.type";

export type SoftwareItemPublic = BaseObjectType & {
	vendor: string;
	license_key: string;
	title: string;

	started_at: string | null;
	expired_at: string | null;
	updated_at: string | null;
};

export type SoftwareFilter = {
	id?: string;

	category?: string;
	last_worker_email?: string;

	vendor?: string;
	license_key?: string;
	title?: string;

	purchase_price?: number;
	transfer_status?: string;
	last_storage?: string;
};
