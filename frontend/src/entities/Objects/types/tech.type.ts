import type { BaseObjectType } from "./baseObjects.type";

export type TechItem = BaseObjectType & {
	brand: string;
	model: string;
	warranty_started_at: string;
	warranty_end_at: string;
	last_storage_id?: string;
	post_number?: string;
};

export interface TechFilter {
	id?: string;
	brand?: string;
	model?: string;
	last_worker?: string;
	last_storage?: string;
	category?: string;
	quality_status?: string;
	transfer_status?: string;
}
