import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, type FC } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/TechList";
import type { TechFilter, TechItem } from "@entities/Objects/types/tech.type";
import type {
	QualityStatus,
	TransferStatus,
} from "@entities/Objects/types/baseObjects.type";
import TransferSelect from "@features/formElements/ui/TransferSelect";
import QualitySelect from "@features/formElements/ui/QualitySelect";

type ResourcesProps = {
	callPlace?: Extract<TransferStatus, "worker" | "storage">;
	name?: string;
	isBroken?: boolean;
};

const TechSearch: FC<ResourcesProps> = ({
	callPlace,
	name,
	isBroken = undefined,
}) => {
	const { register, watch } = useForm<Partial<TechItem>>();

	const [isWrapped, setWrapped] = useState<boolean>(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const rawValues = watch();

	const filter: TechFilter = {
		id: rawValues.id ?? undefined,
		brand: rawValues.brand ?? undefined,
		model: rawValues.model ?? undefined,
		last_worker:
			callPlace == "worker" ? name : (rawValues.last_worker ?? undefined),
		last_storage:
			callPlace == "storage"
				? name
				: (rawValues.last_storage ?? undefined),
		category: rawValues.category ?? undefined,
		quality_status: rawValues.quality_status ?? undefined,
		transfer_status: rawValues.transfer_status ?? undefined,
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
						label="Бренд"
						register={register("brand")}
						width="240px"
					/>
					<Input
						label="Модель"
						register={register("model")}
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

					<QualitySelect
						register={register("quality_status")}
						isAvailable={isBroken ? false : true}
						defaultValue={isBroken ? "faulty" : undefined}
					/>

					<Input
						label="Категория"
						register={register("category")}
						width="240px"
						type="string"
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

export default TechSearch;
