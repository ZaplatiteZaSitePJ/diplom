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
import { normalize } from "@features/utils/normalizeRequests";

type Props = {
	constFilter?: Partial<SoftwareFilter>;
	isMe?: boolean;
};

const SoftwareSearch: FC<Props> = ({ constFilter = {}, isMe }) => {
	const { register, watch } = useForm<Partial<SoftwareItemPublic>>();
	const [isWrapped, setWrapped] = useState(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const raw = watch();

	const isLocked = (field: keyof SoftwareFilter) =>
		constFilter[field] !== undefined;

	const filter: SoftwareFilter = {
		id: normalize(constFilter.id ?? raw.id),
		category: normalize(constFilter.category ?? raw.category),
		vendor: normalize(constFilter.vendor ?? raw.vendor),
		license_key: normalize(constFilter.license_key ?? raw.license_key),
		title: normalize(constFilter.title ?? raw.title),
		purchase_price: normalize(
			constFilter.purchase_price ?? raw.purchase_price,
		),

		last_worker_email: normalize(
			constFilter.last_worker_email ?? raw.last_worker_email,
		),

		transfer_status: normalize(constFilter.transfer_status),
		last_storage: normalize(constFilter.last_storage),

		// 🔥 ВАЖНО: берём raw, а не constFilter
		expired_at: raw.expired_at ?? constFilter.expired_at,
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
						label="Вендор"
						register={register("vendor")}
						width="240px"
						defaultValue={constFilter.vendor}
						isAvailable={!isLocked("vendor")}
					/>
					<Input
						label="Продукт"
						register={register("title")}
						width="240px"
						type="string"
						defaultValue={constFilter.title}
						isAvailable={!isLocked("title")}
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
						defaultValue={constFilter.category}
						isAvailable={!isLocked("category")}
					/>

					<Input
						label="Последний владелец"
						register={register("last_worker_email")}
						width="240px"
						type="string"
						defaultValue={constFilter.last_worker_email}
						isAvailable={!isLocked("last_worker_email")}
					/>

					<Input
						label="Цена"
						register={register("purchase_price")}
						width="240px"
						type="number"
						defaultValue={constFilter.purchase_price}
						isAvailable={!isLocked("purchase_price")}
					/>
				</div>

				<div
					className={styles.objectFormSearch__additionalContainer}
					style={
						isWrapped ? { display: "block" } : { display: "none" }
					}
				></div>
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

export default SoftwareSearch;
