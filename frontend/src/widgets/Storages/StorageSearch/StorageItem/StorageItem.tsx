import { type FC } from "react";
import styles from "./StorageItem.module.scss";
import { Link } from "react-router-dom";
import type { StorageType } from "@entities/Storages/types/storages.type";

type Props = Pick<
	StorageType,
	"id" | "capacity" | "occupied_cells" | "storageName"
> & {
	mode?: "search" | "replace";
	onSelect?: (storage: StorageType) => void;
};

const StorageItem: FC<Props> = ({
	storageName,
	capacity,
	occupied_cells,
	id,
	mode = "search",
	onSelect,
}) => {
	const fillnes = Math.floor((occupied_cells / capacity) * 100);

	return (
		<div className={styles.replaceMode}>
			<Link
				to={`/storages/${id}`}
				className={styles.storageItem}
				onClick={(e) => {
					if (mode === "replace") {
						e.preventDefault();
						onSelect?.({
							id,
							storageName,
							capacity,
							occupied_cells,
						} as StorageType);
					}
				}}
			>
				<div className={styles.storageItem__box}>
					<div
						className={styles.storageItem__filling}
						style={{
							height: `${fillnes}%`,
							backgroundColor:
								fillnes < 90
									? "var(--light-blue-color)"
									: "var(--red-color)",
						}}
					>
						<p className={styles.storageItem__fillnes}>
							{fillnes}%
						</p>
					</div>
				</div>
				<h2 className={styles.storageItem__name}>{storageName}</h2>
			</Link>
		</div>
	);
};

export default StorageItem;
