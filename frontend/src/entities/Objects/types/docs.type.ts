import type { BaseObjectType } from "./baseObjects.type";

export type DocsItem = BaseObjectType & {
	responsible_worker: string;
	full_signed_at: string;
	responsible_worker_email: string;
	needed_signs: boolean;
	received_signs: boolean;
	doc_number: string;
	doc_type: string;
};
