import { type FC, useRef, useState } from "react";
import Papa from "papaparse";
import { enqueueSnackbar } from "notistack";

import { ButtonFilled } from "@shared/ui/ui-kit";
import { useCreateTechMutation } from "@app/api/items/tech/techAPI";
import Modal from "@features/modal/Modal";

import type {
	QualityStatus,
	TransferStatus,
} from "@entities/Objects/types/baseObjects.type";

import type { TechItem } from "@entities/Objects/types/tech.type";

type Props = {
	storageName?: string;
	free_cells_ammount?: number;
};

type CsvTechItem = {
	universal_name: string;
	category?: string;
	transfer_status: string;
	quality_status: string;
	purchase_price: string;
	occupied_cells: string;
	post_number?: string;
	brand: string;
	model: string;
	warranty_started_at?: string;
	warranty_end_at?: string;
	last_worker_email: string;
};

export const ImportTechCsvButton: FC<Props> = ({
	storageName,
	free_cells_ammount,
}) => {
	const inputRef = useRef<HTMLInputElement>(null);

	const [createTech] = useCreateTechMutation();

	const [failedImports, setFailedImports] = useState<
		Array<{
			rowNumber: number;
			row: CsvTechItem;
			error?: any;
		}>
	>([]);

	const [isFailedModalOpen, setIsFailedModalOpen] = useState(false);

	const handleSelectFile = () => {
		inputRef.current?.click();
	};

	const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
		const file = e.target.files?.[0];
		if (!file) return;

		Papa.parse<CsvTechItem>(file, {
			header: true,
			skipEmptyLines: true,

			complete: async ({ data }) => {
				let successCount = 0;
				let failedCount = 0;
				let neededCells = 0;

				// считаем нужные ячейки
				for (const row of data) {
					const cells = Number(row.occupied_cells);
					if (!Number.isNaN(cells)) {
						neededCells += cells;
					}
				}

				if (!free_cells_ammount || neededCells > free_cells_ammount) {
					enqueueSnackbar(
						`Недостаточно места. Необходимо ${neededCells} ячеек, в хранилище ${storageName} доступно ${free_cells_ammount}`,
						{
							variant: "error",
							autoHideDuration: 7000,
						},
					);
					return;
				}

				const failedRows: {
					rowNumber: number;
					row: CsvTechItem;
					error?: any;
				}[] = [];

				for (const [index, row] of data.entries()) {
					try {
						const techItem: Partial<TechItem> = {
							type_id: 0,
							universal_name: row.universal_name,
							category: row.category,
							transfer_status:
								row.transfer_status as TransferStatus,
							quality_status: row.quality_status as QualityStatus,
							purchase_price: Number(row.purchase_price),
							occupied_cells: Number(row.occupied_cells),
							post_number: row.post_number,
							brand: row.brand,
							model: row.model,
							warranty_started_at: row.warranty_started_at
								? `${row.warranty_started_at}T00:00:00Z`
								: undefined,
							warranty_end_at: row.warranty_end_at
								? `${row.warranty_end_at}T00:00:00Z`
								: undefined,
							last_worker_email:
								row.last_worker_email || "admin@company.com",
							last_storage: storageName,
						};

						await createTech(techItem).unwrap();
						successCount++;
					} catch (error: any) {
						failedCount++;

						failedRows.push({
							rowNumber: index + 1,
							row,
							error: error?.data ?? error,
						});
					}
				}

				// сохраняем ошибки
				setFailedImports(failedRows);

				// открываем модалку если есть ошибки
				if (failedRows.length > 0) {
					setIsFailedModalOpen(true);
				}

				enqueueSnackbar(
					`Импорт завершён. Успешно: ${successCount}, ошибок: ${failedCount}`,
					{
						variant: failedCount > 0 ? "warning" : "success",
						autoHideDuration: 7000,
					},
				);
			},

			error: (error) => {
				console.error("Ошибка чтения CSV:", error);

				enqueueSnackbar("Не удалось прочитать CSV", {
					variant: "error",
					autoHideDuration: 5000,
				});
			},
		});

		e.target.value = "";
	};

	return (
		<>
			<ButtonFilled onClick={handleSelectFile}>
				Импорт объектов (csv)
			</ButtonFilled>

			<input
				ref={inputRef}
				type="file"
				accept=".csv"
				hidden
				onChange={handleFileChange}
			/>

			{isFailedModalOpen && (
				<Modal
					title={`Не удалось импортировать ${failedImports.length} объектов`}
					onClose={() => setIsFailedModalOpen(false)}
				>
					<div
						style={{
							maxHeight: "500px",
							overflowY: "auto",
						}}
					>
						{failedImports.map((item) => (
							<div
								key={item.rowNumber}
								style={{
									padding: "12px",
									borderBottom:
										"1px solid rgba(255,255,255,.1)",
								}}
							>
								<b>Строка {item.rowNumber}</b>

								<div>{item.row.universal_name}</div>

								<div>Категория: {item.row.category}</div>

								<div>
									Ошибка:
									<pre>
										{JSON.stringify(
											item?.error?.data,
											null,
											2,
										)}
									</pre>
								</div>
							</div>
						))}
					</div>
				</Modal>
			)}
		</>
	);
};
