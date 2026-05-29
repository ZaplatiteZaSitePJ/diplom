import styles from "./DeleteModal.module.scss";
import { useState, type FC } from "react";
import { ButtonText } from "@shared/ui/ui-kit";
import type { StorageType } from "@entities/Storages/types/storages.type";
import { useNavigate } from "react-router-dom";
import StorageSearch from "@widgets/Storages/StorageSearch/StorageSearch";
import { useDeleteStorageMutation } from "@app/api/storage/storageAPI";
import { enqueueSnackbar } from "notistack";

type StorageDeleteModalProps = {
	storage: StorageType;
	handleClose?: () => void;
};

const StorageDeleteModal: FC<StorageDeleteModalProps> = ({
	storage,
	handleClose,
}) => {
	const [newPlace, setNewPlace] = useState<StorageType | undefined>();
	const [openTree, setOpenTree] = useState<boolean>(false);
	const [triggerDelete] = useDeleteStorageMutation();

	const navigate = useNavigate();

	const handleSelectStorage = (storage: StorageType) => {
		setNewPlace(storage);
		setOpenTree(false);
	};

	const handleDelete = async () => {
		try {
			console.log(newPlace);
			await triggerDelete({
				id: storage.id,
				newStorageName: newPlace?.storageName,
			}).unwrap();
			enqueueSnackbar(`Хранилище ${storage.id} Успешно удалено!`, {
				variant: "success",
				autoHideDuration: 7000,
			});
			handleClose?.();
			navigate("/storages");
		} catch (error: any) {
			enqueueSnackbar(
				`Ошибка! Хранилище ${storage.id} не удалось удалить! Попробуйте позже.`,
				{
					variant: "error",
					autoHideDuration: 7000,
				},
			);
			console.log(error);
		}
	};

	return (
		<div className={styles.deleteModule}>
			<div className={styles.deleteModule__infoContainer}>
				<p className={styles.deleteModule__info}>
					{storage.storageName} ({storage.occupied_cells} занято)
				</p>

				<p className={styles.deleteModule__info}>→</p>

				<p className={styles.deleteModule__info}>
					{newPlace
						? `${newPlace.storageName} (${
								(newPlace.capacity || 0) -
								(newPlace.occupied_cells || 0)
							} свободно)`
						: "Потеря объектов"}
				</p>
			</div>

			<div className={styles.deleteModule__replaceContainer}>
				<ButtonText
					textSize="var(--smallest-font-size)"
					textWeight={100}
					className={styles.deleteModule__openTreeButton}
					onClick={() => setOpenTree((prev) => !prev)}
				>
					Выбрать храналище для перемещания предметов ↓
				</ButtonText>
			</div>

			{openTree && (
				<StorageSearch
					mode="replace"
					storage={storage}
					onSelect={handleSelectStorage}
				/>
			)}

			<div className={styles.deleteModule__buttonContainer}>
				<ButtonText
					type="button"
					textSize="var(--normal-font-size)"
					textColor="var(--white-color)"
					onClick={handleClose}
				>
					Отмена
				</ButtonText>

				<ButtonText
					textWeight={100}
					textSize="var(--normal-font-size)"
					textColor="var(--white-color)"
					type="button"
					onClick={handleDelete}
				>
					Удалить
				</ButtonText>
			</div>
		</div>
	);
};

export default StorageDeleteModal;
