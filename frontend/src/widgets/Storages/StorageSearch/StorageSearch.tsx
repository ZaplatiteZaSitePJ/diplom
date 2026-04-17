import { useEffect, type FC } from "react";
import Input from "@shared/ui/ui-kit/inputs/ui/Input";
import styles from "./StorageSearch.module.scss";
import StorageList from "../StorageList/StorageList";
import { useForm } from "react-hook-form";
import type { SearchCriteria } from "./storageSearchCriteria";
import type { StorageType } from "@entities/Storages/types/storages.type";

type StorageSearchProps = {
	mode?: "search" | "replace";
	storage?: StorageType;
	onSelect?: (storage: StorageType) => void;
};

const StorageSearch: FC<StorageSearchProps> = ({
	mode = "search",
	storage,
	onSelect,
}) => {
	const { register, watch, reset } = useForm<SearchCriteria>();

	const name = watch("name");
	const city = watch("city");

	useEffect(() => {
		if (!storage) return;

		reset({
			city: storage.city || "",
		});
	}, [storage, reset]);

	return (
		<div
			className={styles.storageSearch}
			style={
				mode == "search"
					? { backgroundColor: "var(--blue-color)" }
					: { backgroundColor: "var(--dark-blue-color)" }
			}
		>
			<form className={styles.storageSearch__form}>
				<Input
					placeholder="Введите название хранилища"
					width={mode == "search" ? "540px" : "420px"}
					register={register("name")}
					label="Название хранилища"
				/>

				<Input
					placeholder="Город"
					width={mode == "search" ? "240px" : "200px"}
					register={register("city")}
					label="Город"
				/>
			</form>

			<div className={styles.storageSearch__list}>
				<StorageList
					name={name}
					city={city}
					occupied_cells={storage?.occupied_cells}
					mode={mode}
					onSelect={onSelect}
				/>
			</div>
		</div>
	);
};

export default StorageSearch;
