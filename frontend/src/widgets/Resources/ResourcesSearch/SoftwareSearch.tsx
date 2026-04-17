import { useForm } from "react-hook-form";
import { useState, type FC } from "react";
import styles from "./FormSearch.module.scss";

import type {
	SoftwareFilter,
	SoftwareItemPublic,
} from "@entities/Objects/types/software.type";
import cn from "classnames";

import { Input } from "@shared/ui/ui-kit";
import ObjectList from "../ObjectList/SoftwareList";

type Props = {
	callPlace?: "worker" | "storage";
	name?: string;
};

const SoftwareSearch: FC<Props> = ({ callPlace, name }) => {
	const { register, watch } = useForm<Partial<SoftwareItemPublic>>();
	const [isWrapped, setWrapped] = useState(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const raw = watch();

	const filter: SoftwareFilter = {
		id: raw.id ?? undefined,
		category: raw.category ?? undefined,
		vendor: raw.vendor ?? undefined,
		license_key: raw.license_key ?? undefined,
		title: raw.title ?? undefined,
		purchase_price: raw.purchase_price ?? undefined,

		last_worker_email:
			callPlace === "worker"
				? name
				: (raw.last_worker_email ?? undefined),
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
					/>
					<Input
						label="Вендор"
						register={register("vendor")}
						width="240px"
					/>
					<Input
						label="Продукт"
						register={register("title")}
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
					<Input
						label="Категория"
						register={register("category")}
						width="240px"
						type="string"
					/>

					<Input
						label="Последний владелец"
						register={register("last_worker")}
						width="240px"
						type="string"
						value={callPlace == "worker" ? name : undefined}
						isAvailable={callPlace == "worker" ? false : true}
					/>

					<Input
						label="Цена"
						register={register("purchase_price")}
						width="240px"
						type="number"
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

export default SoftwareSearch;
