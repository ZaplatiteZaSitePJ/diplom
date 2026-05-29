import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, type FC } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/DocsList";
import TransferSelect from "@features/formElements/ui/TransferSelect";

import type { DocsFilter, DocsItem } from "@entities/Objects/types/docs.type";

type Props = {
	constFilter?: Partial<DocsFilter>;
	isMe?: boolean;
};

const DocsSearch: FC<Props> = ({ constFilter = {}, isMe }) => {
	const { register, watch } = useForm<Partial<DocsItem>>();
	const [isWrapped, setWrapped] = useState(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const raw = watch();

	const isLocked = (field: keyof DocsFilter) =>
		constFilter[field] !== undefined;

	const filter: DocsFilter = {
		id: constFilter.id ?? raw.id ?? undefined,
		doc_number: constFilter.doc_number ?? raw.doc_number ?? undefined,
		category: constFilter.category ?? raw.category ?? undefined,

		last_storage: constFilter.last_storage ?? raw.last_storage ?? undefined,

		transfer_status:
			constFilter.transfer_status ?? raw.transfer_status ?? undefined,

		responsible_worker_email:
			constFilter.last_worker_email ??
			raw.responsible_worker_email ??
			undefined,
	};

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
						defaultValue={constFilter.id}
						isAvailable={!isLocked("id")}
					/>

					<Input
						label="Категория"
						register={register("category")}
						width="240px"
						defaultValue={constFilter.category}
						isAvailable={!isLocked("category")}
					/>

					<Input
						label="Номер"
						register={register("doc_number")}
						width="240px"
						type="string"
						defaultValue={constFilter.doc_number}
						isAvailable={!isLocked("doc_number")}
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
						isAvailable={!isLocked("transfer_status")}
						defaultValue={constFilter.transfer_status}
					/>

					<Input
						label="Последнее хранилище"
						register={register("last_storage")}
						width="240px"
						type="string"
						defaultValue={constFilter.last_storage}
						isAvailable={!isLocked("last_storage")}
					/>

					<Input
						label="Подписант (email)"
						register={register("responsible_worker_email")}
						width="240px"
						type="string"
						defaultValue={constFilter.last_worker_email}
						isAvailable={!isLocked("responsible_worker_email")}
					/>
				</div>

				<div
					className={styles.objectFormSearch__additionalContainer}
					style={
						isWrapped ? { display: "block" } : { display: "none" }
					}
				/>
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
				<ObjectList filter={filter} isMe={isMe} />
			</div>
		</div>
	);
};

export default DocsSearch;
