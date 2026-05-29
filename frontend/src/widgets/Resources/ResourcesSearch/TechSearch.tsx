import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, type FC } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/TechList";
import type { TechFilter, TechItem } from "@entities/Objects/types/tech.type";
import TransferSelect from "@features/formElements/ui/TransferSelect";
import QualitySelect from "@features/formElements/ui/QualitySelect";

type Props = {
	constFilter?: Partial<TechFilter>;
	isMe?: boolean;
};

const TechSearch: FC<Props> = ({ constFilter = {}, isMe = false }) => {
	const { register, watch } = useForm<Partial<TechItem>>();
	const [isWrapped, setWrapped] = useState(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const raw = watch();

	const isLocked = (field: keyof TechFilter) =>
		constFilter[field] !== undefined;

	const filter: TechFilter = {
		id: constFilter.id ?? raw.id ?? undefined,
		brand: constFilter.brand ?? raw.brand ?? undefined,
		model: constFilter.model ?? raw.model ?? undefined,
		category: constFilter.category ?? raw.category ?? undefined,

		last_worker: constFilter.last_worker ?? raw.last_worker ?? undefined,

		last_storage: constFilter.last_storage ?? raw.last_storage ?? undefined,

		quality_status:
			constFilter.quality_status ?? raw.quality_status ?? undefined,

		transfer_status:
			constFilter.transfer_status ?? raw.transfer_status ?? undefined,
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
						label="Бренд"
						register={register("brand")}
						width="240px"
						defaultValue={constFilter.brand}
						isAvailable={!isLocked("brand")}
					/>
					<Input
						label="Модель"
						register={register("model")}
						width="240px"
						type="string"
						defaultValue={constFilter.model}
						isAvailable={!isLocked("model")}
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
						label="Последний владелец"
						register={register("last_worker")}
						width="240px"
						type="string"
						defaultValue={constFilter.last_worker}
						isAvailable={!isLocked("last_worker")}
					/>

					<QualitySelect
						register={register("quality_status")}
						isAvailable={!isLocked("quality_status")}
						defaultValue={constFilter.quality_status}
					/>

					<Input
						label="Категория"
						register={register("category")}
						width="240px"
						type="string"
						defaultValue={constFilter.category}
						isAvailable={!isLocked("category")}
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

export default TechSearch;
