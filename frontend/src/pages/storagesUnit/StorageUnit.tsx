import EntitiesLayout from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";
import styles from "./StorageUnit.module.scss";
import StorageForm from "@widgets/Storages/StorageForm/StorageForm";
import { useParams } from "react-router-dom";
import type { StorageType } from "@entities/Storages/types/storages.type";
import type { FC } from "react";
import { useGetStorageByIdQuery } from "@app/api/storage/storageAPI";

export default function StorageUnit() {
	const { id } = useParams();
	const { data, isLoading, isSuccess } = useGetStorageByIdQuery(id as string);

	console.log(data);

	const storage = data?.data;

	return (
		<EntitiesLayout
			treeLink={`/tree/${storage?.id}`}
			title={storage?.storageName}
			subTitle={`Хранилище / ${storage?.city}`}
			form={
				<StorageForm storage={storage} mode="save" onSaved={() => {}} />
			}
			statistic={<StorageStatistic storage={storage} />}
			entitie={storage}
		/>
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
