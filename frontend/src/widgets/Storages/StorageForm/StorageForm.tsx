import { useForm } from "react-hook-form";
import styles from "./StorageForm.module.scss";
import type { StorageType } from "@entities/Storages/types/storages.type";
import NameField from "@features/formElements/NameField";
import { nameOption } from "@features/formElements/options/name.options";
import { ButtonText, Input } from "@shared/ui/ui-kit";
import { useEffect, type FC } from "react";
import CityField from "@features/formElements/CityField";
import {
	useCreateStorageMutation,
	useUpdateStorageMutation,
} from "@app/api/storage/storageAPI";
// import postStorage from "@features/api/axios/requests/storages/postStorage";
// import putStorage from "@features/api/axios/requests/storages/putStorage";

type StorageFormProps = {
	storage?: StorageType | undefined;
	isCell?: boolean;
	parentStorageId?: number | null;
	mode: "save" | "create";
	handleClose?: () => void;
	isReadOnly?: boolean;
	setReadOnly?: () => void;
	onSaved?: () => void;
};

type StorageFormType = StorageType & {
	increaseValue: number;
	decreaseValue: number;
};

const StorageForm: FC<StorageFormProps> = ({
	storage,
	isReadOnly = false,
	mode,
	handleClose,
	setReadOnly,
	onSaved,
}) => {
	const { register, reset, getValues, handleSubmit, setValue } =
		useForm<StorageFormType>();

	useEffect(() => {
		reset(storage);
	}, [storage, reset]);

	const [triggerPost] = useCreateStorageMutation();
	const [trigerUpdate] = useUpdateStorageMutation();

	const onSubmit = async () => {
		const capacity =
			Number(storage?.capacity || 0) +
			Number(getValues("increaseValue")) -
			Number(getValues("decreaseValue"));
		if (mode === "create" && handleClose) {
			await triggerPost({
				...getValues(),
				capacity,
			});
			handleClose();
		}
		if (mode === "save") {
			if (!storage?.id) return;

			trigerUpdate({
				id: storage.id,
				body: {
					...getValues(),
					capacity,
				},
			});
			setReadOnly?.();
			onSaved?.();
		}
	};

	return (
		<>
			<form
				className={styles.storageForm}
				onSubmit={handleSubmit(onSubmit)}
			>
				<NameField
					register={register("storageName", nameOption)}
					subContent={
						<p className={styles.storageForm__subInfo}>
							{"< 50 символов"}
						</p>
					}
					isAvailable={!isReadOnly}
				/>

				<div className={styles.storageForm__capacityContainer}>
					<p className={styles.storageForm__capacityLabel}>
						Вместимость (в ячейках)
					</p>

					<p className={styles.storageForm__capacityValue}>
						{storage?.capacity ? storage.capacity : 0}
					</p>

					{!isReadOnly && (
						<div
							className={
								styles.storageForm__capacityButtonContainer
							}
						>
							<Input
								label="Увеличить +"
								type="number"
								register={register("increaseValue", {
									onChange: (e) => {
										const raw = Number(e.target.value);
										setValue(
											"increaseValue",
											raw < 0 ? 0 : raw,
											{
												shouldValidate: true,
												shouldDirty: true,
											},
										);
									},
								})}
							/>
							<Input
								label="Уменьшить -"
								type="number"
								register={register("decreaseValue", {
									onChange: (e) => {
										const raw = Number(e.target.value);
										setValue(
											"decreaseValue",
											raw < 0 ? 0 : raw,
											{
												shouldValidate: true,
												shouldDirty: true,
											},
										);
									},
								})}
							/>
						</div>
					)}
				</div>

				<div className={styles.storageForm__subInfoContainer}>
					<CityField
						register={register("city")}
						isAvailable={!isReadOnly}
						width="200px"
					/>
				</div>

				{!isReadOnly && (
					<div className={styles.storageForm__buttonsContainer}>
						<ButtonText
							textWeight={100}
							textSize="var(--normal-font-size)"
							textColor="var(--grey-color)"
							onClick={() => reset()}
						>
							Сбросить
						</ButtonText>

						<ButtonText
							textSize="var(--normal-font-size)"
							textColor="var(--green-color)"
							type="submit"
						>
							{mode === "save" ? "Сохранить" : "Создать"}
						</ButtonText>
					</div>
				)}
			</form>
		</>
	);
};

export default StorageForm;
