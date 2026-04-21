import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, type FC } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/DocsList";
import type { TransferStatus } from "@entities/Objects/types/baseObjects.type";
import TransferSelect from "@features/formElements/ui/TransferSelect";
import type { DocsFilter, DocsItem } from "@entities/Objects/types/docs.type";

type ResourcesProps = {
	callPlace?: Extract<TransferStatus, "worker" | "storage">;
	name?: string;
};

const DocsSearch: FC<ResourcesProps> = ({ callPlace, name }) => {
	const { register, watch } = useForm<Partial<DocsItem>>();

	const [isWrapped, setWrapped] = useState<boolean>(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const rawValues = watch();

	const filter: DocsFilter = {
		id: rawValues.id ?? undefined,
		doc_number: rawValues.doc_number ?? undefined,

		last_worker_email:
			callPlace === "worker"
				? name
				: (rawValues.last_worker_email ?? undefined),

		last_storage:
			callPlace === "storage"
				? name
				: (rawValues.last_storage ?? undefined),

		category: rawValues.category ?? undefined,

		transfer_status:
			callPlace === "worker"
				? "worker"
				: callPlace === "storage"
					? "storage"
					: (rawValues.transfer_status ?? undefined),
	};

	console.log("Form changed:", filter);

	return (
		<div className={styles.objectFormSearch}>
			<form
				className={styles.objectFormSearch__filter}
				style={
					isWrapped
						? {
								height: "368px",
								overflow: "auto",
								boxShadow: "0 12px 30px rgba(0, 0, 0, 0.75)",
							}
						: { height: "138px", overflow: "hidden" }
				}
			>
				<div className={styles.objectFormSearch__filterContainer}>
					<Input
						label="Артикул"
						register={register("id")}
						width="240px"
					/>
					<Input
						label="Категория"
						register={register("category")}
						width="240px"
					/>
					<Input
						label="Номер"
						register={register("doc_number")}
						width="240px"
						type="string"
					/>
				</div>

				<div
					className={styles.objectFormSearch__filterContainer}
					style={
						isWrapped ? { display: "flex" } : { display: "none" }
					}
				>
					<TransferSelect
						register={register("transfer_status")}
						isAvailable={callPlace ? false : true}
						defaultValue={
							callPlace
								? callPlace == "storage"
									? "storage"
									: "worker"
								: undefined
						}
					/>

					<Input
						label="Последнее хранилище"
						register={register("last_storage")}
						width="240px"
						type="string"
						value={callPlace == "storage" ? name : undefined}
						isAvailable={callPlace == "storage" ? false : true}
					/>

					<Input
						label="Последний владелец"
						register={register("last_worker")}
						width="240px"
						type="string"
						value={callPlace == "worker" ? name : undefined}
						isAvailable={callPlace == "worker" ? false : true}
					/>
				</div>

				<div
					className={styles.objectFormSearch__additionalContainer}
					style={
						isWrapped ? { display: "block" } : { display: "none" }
					}
				>
					{" "}
				</div>
			</form>

			<button
				className={cn(styles.objectFormSearch__unwrapButton, {
					[styles.wrapped]: isWrapped,
				})}
				type="button"
				onClick={handleWrap}
			>
				↓
			</button>

			<div className={styles.objectFormSearch__objectListPlace}>
				<ObjectList filter={filter} />
			</div>
		</div>
	);
};

export default DocsSearch;
