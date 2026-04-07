import Input from "@shared/ui/ui-kit/inputs/ui/Input";
import styles from "./StorageSearch.module.scss";
import StorageList from "../StorageList/StorageList";
import { useForm } from "react-hook-form";
import type { SearchCriteria } from "./storageSearchCriteria";

export default function StorageSearch() {
	const { register, watch } = useForm<SearchCriteria>();

	const name = watch("name");
	const city = watch("city");

	return (
		<div className={styles.storageSearch}>
			<form className={styles.storageSearch__form}>
				<Input
					placeholder="Введите название хранилища"
					width="540px"
					register={register("name")}
					label="Название хранилища"
				/>

				<Input
					placeholder="Город"
					width="240px"
					register={register("city")}
					label="Город"
				/>
			</form>

			<div className={styles.storageSearch__list}>
				<StorageList name={name} city={city} />
			</div>
		</div>
	);
}
