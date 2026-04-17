export type CategoriesList =
	| "Документы"
	| "Программное обеспечение"
	| "Техника";

export const TRANSFER_STATUS = {
	WORKER: "worker",
	STORAGE: "storage",
	TRANSFERING_TO_WORKER: "transfering_to_worker",
	TRANSFERING_TO_STORAGE: "transfering_to_storage",
} as const;

export type TransferStatus =
	(typeof TRANSFER_STATUS)[keyof typeof TRANSFER_STATUS];

export type TransferStatusRuLabel =
	| "У работника"
	| "На складе"
	| "В пути к работнику"
	| "В пути на склад";

export const TransferStatusList: {
	label: TransferStatusRuLabel | "Сбросить";
	value: TransferStatus | undefined;
}[] = [
	{ label: "Сбросить", value: undefined },
	{ label: "У работника", value: TRANSFER_STATUS.WORKER },
	{ label: "На складе", value: TRANSFER_STATUS.STORAGE },
	{
		label: "В пути к работнику",
		value: TRANSFER_STATUS.TRANSFERING_TO_WORKER,
	},
	{
		label: "В пути на склад",
		value: TRANSFER_STATUS.TRANSFERING_TO_STORAGE,
	},
];

export const QUALITY_STATUS = {
	NEW: "new",
	USED: "used",
	DAMAGED: "damaged",
	FAULTY: "faulty",
} as const;

export type QualityStatus =
	(typeof QUALITY_STATUS)[keyof typeof QUALITY_STATUS];

export type QualityStatusRuLabel =
	| "Новый"
	| "Пользованный"
	| "Повреждён"
	| "Неисправен";

export const QualityStatusList: {
	label: QualityStatusRuLabel | "Сбросить";
	value: QualityStatus | undefined;
}[] = [
	{ label: "Сбросить", value: undefined },
	{ label: "Новый", value: QUALITY_STATUS.NEW },
	{ label: "Пользованный", value: QUALITY_STATUS.USED },
	{ label: "Повреждён", value: QUALITY_STATUS.DAMAGED },
	{ label: "Неисправен", value: QUALITY_STATUS.FAULTY },
];

export type BaseObjectType = {
	id: string;
	universal_name?: string;
	type: CategoriesList;

	category: string;

	last_storage?: string | null;
	last_worker?: string | null;
	last_worker_email?: string;

	transfer_status: TransferStatus;
	quality_status: QualityStatus;

	purchase_price?: number;

	occupied_cells: number;
};

export const getTransferLabel = (value: TransferStatus) =>
	TransferStatusList.find((item) => item.value === value)?.label ?? "Все";

export const getQualityLabel = (value: QualityStatus) =>
	QualityStatusList.find((item) => item.value === value)?.label ?? undefined;

export const makeLightID = (value: string) => {
	if (!value) return "";

	if (value.length <= 6) return value;

	return `${value.slice(0, 3)}...${value.slice(-3)}`;
};
