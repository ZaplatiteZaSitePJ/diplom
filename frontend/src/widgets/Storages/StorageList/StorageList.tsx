import type { FC } from "react";
import StorageItem from "../StorageSearch/StorageItem/StorageItem";
import styles from "./StorageList.module.scss";
import storageFiltration from "./storageFilter";
import { useGetStoragesQuery } from "@app/api/storage/storageAPI";

type StorageListProps = {
	name?: string;
	city?: string;
};

const StorageList: FC<StorageListProps> = ({ name, city }) => {
	const { data, isLoading } = useGetStoragesQuery();
	console.log(data);

	const storageFiltred = storageFiltration(name, city, data?.data || []);
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
					/>
				);
			})}
		</div>
	);
};

export default StorageList;
