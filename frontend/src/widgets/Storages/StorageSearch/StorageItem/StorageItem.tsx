import { type FC } from "react";
import styles from "./StorageItem.module.scss";
import { Link } from "react-router-dom";
import type { StorageType } from "@entities/Storages/types/storages.type";

const StorageItem: FC<
	Pick<StorageType, "id" | "capacity" | "occupied_cells" | "storageName">
> = ({ storageName, capacity, occupied_cells, id }) => {
	const fillnes = Math.floor((occupied_cells / capacity) * 100);

	return (
		<Link to={`/storages/${id}`} className={styles.storageItem}>
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
					<p className={styles.storageItem__fillnes}>{fillnes}%</p>
				</div>
			</div>
			<h2 className={styles.storageItem__name}>{storageName}</h2>
		</Link>
	);
};

export default StorageItem;
