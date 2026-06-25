import type { DocsItem } from "@entities/Objects/types/docs.type";
import type { TechItem } from "@entities/Objects/types/tech.type";
import { fromISODate } from "@features/utils/dateConverter";
import type { SoftwareItemPublic } from "@entities/Objects/types/software.type";

export const techDataForPush = (data: TechItem) => {
	return {
		...data,

		occupied_cells: data.occupied_cells
			? Number(data.occupied_cells)
			: undefined,

		purchase_price: data.purchase_price
			? Number(data.purchase_price)
			: undefined,

		// убираем пустые строки
		last_worker_email: data.last_worker_email || undefined,

		// даты → ISO через helper
		warranty_started_at: fromISODate(data.warranty_started_at),
		warranty_end_at: fromISODate(data.warranty_end_at),
		sended_at: fromISODate(data.sended_at),
		arrived_at: fromISODate(data.arrived_at),
	};
};

export const docsDataForPush = (data: DocsItem): DocsItem => {
	return {
		...data,
		received_signs: data.received_signs ? Number(data.received_signs) : 0,

		needed_signs: data.needed_signs ? Number(data.needed_signs) : 0,

		// убираем пустые строки
		last_worker_email: data.last_worker_email || undefined,
	};
};

export const softwareDataForPush = (data: SoftwareItemPublic) => {
	return {
		...data,

		occupied_cells: data.occupied_cells
			? Number(data.occupied_cells)
			: undefined,

		purchase_price: data.purchase_price
			? Number(data.purchase_price)
			: undefined,

		last_worker_email: data.last_worker_email || undefined,

		started_at: fromISODate(data.started_at),
		expired_at: fromISODate(data.expired_at),
	};
};
