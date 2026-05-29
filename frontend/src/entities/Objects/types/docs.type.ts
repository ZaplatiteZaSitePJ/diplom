import type { BaseObjectType } from "./baseObjects.type";

export type DocsItem = BaseObjectType & {
	responsible_worker: string;
	full_signed_at: string;
	responsible_worker_email: string;
	needed_signs: number;
	received_signs: number;
	doc_number: string;
};

export interface DocsFilter {
	id?: string;
	doc_number?: string;
	last_worker_email?: string;
	last_storage?: string;
	category?: string;
	transfer_status?: string;
	responsible_worker_email?: string;
}
