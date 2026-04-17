import type { FC } from "react";
import StorageItem from "../StorageSearch/StorageItem/StorageItem";
import styles from "./StorageList.module.scss";
import storageFiltration from "./storageFilter";
import { useGetStoragesQuery } from "@app/api/storage/storageAPI";
import type { StorageType } from "@entities/Storages/types/storages.type";

type StorageListProps = {
	name?: string;
	city?: string;
	occupied_cells?: number;
	mode?: "search" | "replace";
	onSelect?: (storage: StorageType) => void;
};

const StorageList: FC<StorageListProps> = ({
	name,
	city,
	occupied_cells,
	mode,
	onSelect,
}) => {
	const { data, isLoading } = useGetStoragesQuery();
	console.log(data);

	const storageFiltred = storageFiltration(
		name,
		city,
		data?.data || [],
		occupied_cells,
	);
	return (
		<div className={styles.storageList}>
			{storageFiltred.map((element) => {
				return (
					<StorageItem
						key={element.id}
						storageName={element.storageName}
						capacity={element.capacity}
						id={element.id}
						occupied_cells={element.occupied_cells}
						mode={mode}
						onSelect={onSelect}
					/>
				);
			})}
		</div>
	);
};

export default StorageList;
