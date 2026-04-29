import styles from "./StorageUnit.module.scss";
import StorageForm from "@widgets/Storages/StorageForm/StorageForm";
import { useParams } from "react-router-dom";
import type { StorageType } from "@entities/Storages/types/storages.type";
import type { FC } from "react";
import { useGetStorageByIdQuery } from "@app/api/storage/storageAPI";
import ResourcesSearch from "@widgets/Resources/ResourcesSearch/ResourcesSearch";
import { StorageLayout } from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";

export default function StorageUnit() {
	const { id } = useParams();
	const { data } = useGetStorageByIdQuery(id as string);

	const storage = data?.data;

	if (!storage) return null;

	return (
		<div className={styles.page}>
			<StorageLayout
				title={storage.storageName}
				subTitle={`Хранилище / ${storage.city}`}
				form={
					<StorageForm
						storage={storage}
						mode="save"
						onSaved={() => {}}
					/>
				}
				statistic={<StorageStatistic storage={storage} />}
				entitie={storage}
			/>

			<div className={styles.page__resources}>
				<h2>Ресурсы на хранении</h2>
				<ResourcesSearch
					callPlace="storage"
					name={storage.storageName}
				/>
			</div>
		</div>
	);
}

const StorageStatistic: FC<{ storage?: StorageType }> = ({ storage }) => {
	return (
		<div className={styles.statistic}>
			<h2>Статистика</h2>

			<div className={styles.statistic__container}>
				<h3 className={styles.statistic__title}>Свободных ячеек</h3>
				<div>
					<ul>
						<li className={styles.statistic__textContainer}>
							<p className={styles.statistic__infoOnly}>
								{(storage?.capacity || 0) -
									(storage?.occupied_cells || 0)}
							</p>
						</li>
					</ul>
				</div>
			</div>

			<div className={styles.statistic__container}>
				<h3 className={styles.statistic__title}>Объекты</h3>
				<p className={styles.statistic__infoOnly}>
					{storage?.items_amount ?? 0} шт.
				</p>
			</div>

			<div className={styles.statistic__container}>
				<h3 className={styles.statistic__title}>Заполненность:</h3>
				<p className={styles.statistic__infoOnly}>
					{Math.floor(
						((storage?.occupied_cells || 0) /
							(storage?.capacity || 1)) *
							100,
					)}{" "}
					%
				</p>
			</div>
		</div>
	);
};
