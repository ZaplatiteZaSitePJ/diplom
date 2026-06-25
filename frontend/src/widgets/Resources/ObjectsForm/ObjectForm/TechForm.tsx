import { useForm, type SubmitHandler } from "react-hook-form";
import { type FC, useEffect, useState } from "react";

import styles from "./ObjectForm.module.scss";

import { ButtonText, Input } from "@shared/ui/ui-kit";

import type { TechItem } from "@entities/Objects/types/tech.type";
import type { StorageType } from "@entities/Storages/types/storages.type";

import {
	useCreateTechMutation,
	useUpdateTechMutation,
} from "@app/api/items/tech/techAPI";

import { toISODate } from "@features/utils/dateConverter";
import { techDataForPush } from "./dataForPush";

import TransferSelect from "@features/formElements/ui/TransferSelect";
import QualitySelect from "@features/formElements/ui/QualitySelect";
import CategorySelect from "@features/formElements/ui/CategorySelect";

import StorageSearch from "@widgets/Storages/StorageSearch/StorageSearch";

import { enqueueSnackbar } from "notistack";
import { useGetStorageByIdQuery } from "@app/api/storage/storageAPI";
import type { UserType } from "@entities/User/types/user.type";
import PersonalSearch from "@widgets/Resources/ResourcesSearch/PersonalSearch";

type TechFormProps = {
	object?: TechItem;
	mode: "save" | "create";
	isReadOnly?: boolean;
	setReadOnly?: () => void;
};

const TechForm: FC<TechFormProps> = ({
	object,
	mode,
	isReadOnly = false,
	setReadOnly,
}) => {
	const { register, handleSubmit, reset, setValue, watch } =
		useForm<TechItem>();
	console.log(object);

	const [triggerPost] = useCreateTechMutation();
	const [triggerPatch] = useUpdateTechMutation();
	const { data: storage } = useGetStorageByIdQuery(
		object?.last_storage_id ?? "",
		{
			skip: !object?.last_storage_id,
		},
	);

	const transferStatus = watch("transfer_status");
	const [newPlace, setNewPlace] = useState<StorageType | undefined>();
	const [newPerson, setNewPerson] = useState<UserType | undefined>();

	const [openTreeStorage, setOpenTreeStorage] = useState<boolean>(false);
	const [openTreePersonal, setOpenTreePersonal] = useState<boolean>(false);

	// Состояния для "Откуда" и "Куда"
	const [movementFromStorage, setMovementFromStorage] = useState<
		StorageType | undefined
	>();
	const [movementFromPerson, setMovementFromPerson] = useState<
		UserType | undefined
	>();
	const [movementToStorage, setMovementToStorage] = useState<
		StorageType | undefined
	>();
	const [movementToPerson, setMovementToPerson] = useState<
		UserType | undefined
	>();

	const [openMovementFrom, setOpenMovementFrom] = useState<boolean>(false);
	const [openMovementTo, setOpenMovementTo] = useState<boolean>(false);

	// Определяем типы для "Откуда" и "Куда" на основе transfer_status
	const isTransferingToStorage = transferStatus === "transfering_to_storage";
	const isTransferingToWorker = transferStatus === "transfering_to_worker";

	// "Откуда": если перемещаем на склад — значит было у работника (и наоборот)
	const movementFromType: "storage" | "worker" | null = isTransferingToStorage
		? "worker"
		: isTransferingToWorker
			? "storage"
			: null;

	// "Куда": совпадает с направлением transfer_status
	const movementToType: "storage" | "worker" | null = isTransferingToStorage
		? "storage"
		: isTransferingToWorker
			? "worker"
			: null;

	useEffect(() => {
		if (object) {
			reset({
				...object,
				warranty_started_at: toISODate(object.warranty_started_at),
				warranty_end_at: toISODate(object.warranty_end_at),
				sended_at: toISODate(object.sended_at) || undefined,
				arrived_at: toISODate(object.arrived_at) || undefined,
			});

			if (object.last_storage) {
				setNewPlace({
					id: "",
					storageName: object.last_storage,
					capacity: 0,
					occupied_cells: 0,
				} as StorageType);
			}

			// Предзаполнение movement_from / movement_to из существующих данных
			if (object.movement_from) {
				// Пытаемся определить тип по текущему transfer_status
				if (object.transfer_status === "transfering_to_storage") {
					setMovementFromPerson({
						email: object.movement_from,
					} as UserType);
					setValue("movement_from", object.movement_from);
				} else if (object.transfer_status === "transfering_to_worker") {
					setMovementFromStorage({
						id: "",
						storageName: object.movement_from,
						capacity: 0,
						occupied_cells: 0,
					} as StorageType);
					setValue("movement_from", object.movement_from);
				}
			}

			if (object.movement_to) {
				if (object.transfer_status === "transfering_to_storage") {
					setMovementToStorage({
						id: "",
						storageName: object.movement_to,
						capacity: 0,
						occupied_cells: 0,
					} as StorageType);
					setValue("movement_to", object.movement_to);
				} else if (object.transfer_status === "transfering_to_worker") {
					setMovementToPerson({
						email: object.movement_to,
					} as UserType);
					setValue("movement_to", object.movement_to);
				}
			}
		}
	}, [object, reset, setValue]);

	// Сбрасываем выбранные значения "Откуда"/"Куда" при смене transfer_status
	useEffect(() => {
		setMovementFromStorage(undefined);
		setMovementFromPerson(undefined);
		setMovementToStorage(undefined);
		setMovementToPerson(undefined);
		setValue("movement_from", "");
		setValue("movement_to", "");
		setOpenMovementFrom(false);
		setOpenMovementTo(false);
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [transferStatus, setValue]);

	const handleSelectStorage = (storage: StorageType) => {
		setNewPlace(storage);
		setValue("last_storage", storage.storageName);
		setOpenTreeStorage(false);

		enqueueSnackbar(`Выбрано хранилище ${storage.storageName}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const handleSelectPerson = (person: UserType) => {
		setNewPerson(person);
		setValue("last_worker_email", person.email);
		setOpenTreePersonal(false);

		enqueueSnackbar(`Выбран пользователь ${person.email}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const handleSelectMovementFromStorage = (storage: StorageType) => {
		setMovementFromStorage(storage);
		setValue("movement_from", storage.storageName);
		setOpenMovementFrom(false);

		enqueueSnackbar(`Откуда: ${storage.storageName}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const handleSelectMovementFromPerson = (person: UserType) => {
		setMovementFromPerson(person);
		setValue("movement_from", person.email);
		setOpenMovementFrom(false);

		enqueueSnackbar(`Откуда: ${person.email}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const handleSelectMovementToStorage = (storage: StorageType) => {
		setMovementToStorage(storage);
		setValue("movement_to", storage.storageName);
		setOpenMovementTo(false);

		enqueueSnackbar(`Куда: ${storage.storageName}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const handleSelectMovementToPerson = (person: UserType) => {
		setMovementToPerson(person);
		setValue("movement_to", person.email);
		setOpenMovementTo(false);

		enqueueSnackbar(`Куда: ${person.email}`, {
			variant: "info",
			autoHideDuration: 3000,
		});
	};

	const submitHandler: SubmitHandler<TechItem> = async (data) => {
		const formattedData = techDataForPush(data);
		console.log(formattedData);

		const id = formattedData.id;

		try {
			if (mode === "create") {
				await triggerPost(formattedData).unwrap();

				enqueueSnackbar(
					`Объект ${formattedData.brand} ${formattedData.model} (${id}) успешно создан`,
					{
						variant: "success",
						autoHideDuration: 5000,
					},
				);
			}

			if (mode === "save") {
				await triggerPatch({
					id,
					body: formattedData,
				}).unwrap();

				enqueueSnackbar(
					`Объект ${formattedData.brand} ${formattedData.model} (${id}) успешно обновлен`,
					{
						variant: "success",
						autoHideDuration: 5000,
					},
				);
			}

			setReadOnly?.();
		} catch (err: any) {
			console.error(err);

			if (err?.status === 400 || err?.status === 422) {
				enqueueSnackbar("Ошибка! Введены некорректные данные", {
					variant: "error",
					autoHideDuration: 7000,
				});

				return;
			}

			enqueueSnackbar("Ошибка! Попробуйте позже", {
				variant: "error",
				autoHideDuration: 7000,
			});
		}
	};

	// Отображаемое значение для "Откуда"
	const movementFromDisplayValue =
		movementFromType === "storage"
			? movementFromStorage?.storageName
			: movementFromType === "worker"
				? movementFromPerson?.email
				: "";

	// Отображаемое значение для "Куда"
	const movementToDisplayValue =
		movementToType === "storage"
			? movementToStorage?.storageName
			: movementToType === "worker"
				? movementToPerson?.email
				: "";

	return (
		<form
			className={styles.objectForm}
			onSubmit={handleSubmit(submitHandler)}
		>
			<div className={styles.objectForm__additionalContainer}>
				<h2>Основная информация</h2>

				<Input
					register={register("brand")}
					label="Бренд"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("model")}
					label="Модель"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<CategorySelect
					id={"0"}
					register={register("category")}
					isAvailable={!isReadOnly}
					defaultValue={object?.category}
				/>

				<Input
					register={register("occupied_cells")}
					label="Занимаемое место"
					width="240px"
					isAvailable={!isReadOnly}
					type="number"
				/>

				<Input
					register={register("purchase_price")}
					label="Цена закупки"
					width="240px"
					isAvailable={!isReadOnly}
					type="number"
				/>
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Местонахождение</h2>

				<TransferSelect
					isAvailable={!isReadOnly}
					register={register("transfer_status")}
					defaultValue={object?.transfer_status}
				/>
				<div>
					<Input
						register={register("last_worker_email")}
						type="string"
						label="Последний пользователь"
						width="240px"
						isAvailable={false}
					/>
					{!isReadOnly && (
						<ButtonText
							type="button"
							textSize="var(--smallest-font-size)"
							textWeight={100}
							onClick={() => {
								setOpenTreePersonal((prev) => !prev);
								setOpenTreeStorage(false);
							}}
						>
							{!openTreePersonal
								? "Изменить работника ↓"
								: "Закрыть ↑"}
						</ButtonText>
					)}
				</div>

				<div>
					<Input
						register={register("last_storage")}
						type="string"
						label="Последнее хранилище"
						width="240px"
						isAvailable={false}
					/>
					{!isReadOnly && (
						<ButtonText
							type="button"
							textSize="var(--smallest-font-size)"
							textWeight={100}
							onClick={() => {
								setOpenTreeStorage((prev) => !prev);
								setOpenTreePersonal(false);
							}}
						>
							{!openTreeStorage
								? "Изменить хранилище ↓"
								: "Закрыть ↑"}
						</ButtonText>
					)}
				</div>

				<div className={styles.objectForm__storageContainer}>
					{openTreeStorage && (
						<StorageSearch
							mode="replace"
							onSelect={handleSelectStorage}
							storage={storage?.data}
						/>
					)}

					{openTreePersonal && (
						<PersonalSearch
							mode="replace"
							onSelect={handleSelectPerson}
						/>
					)}
				</div>

				{(isTransferingToStorage || isTransferingToWorker) && (
					<div
						className={`${styles.objectForm__additionalContainer} ${styles.objectForm__additionalContainer_replace}`}
					>
						{/* Откуда */}
						<div>
							<Input
								register={register("movement_from")}
								type="string"
								label={
									movementFromType === "storage"
										? "Откуда (хранилище)"
										: "Откуда (работник)"
								}
								width="240px"
								isAvailable={false}
								value={
									object?.movement_from ||
									movementFromDisplayValue ||
									""
								}
							/>
							{!isReadOnly && movementFromType && (
								<ButtonText
									type="button"
									textSize="var(--smallest-font-size)"
									textWeight={100}
									onClick={() => {
										setOpenMovementFrom((prev) => !prev);
										setOpenMovementTo(false);
									}}
								>
									{!openMovementFrom
										? `Выбрать ${movementFromType === "storage" ? "хранилище" : "работника"} ↓`
										: "Закрыть ↑"}
								</ButtonText>
							)}
						</div>

						{/* Куда */}
						<div>
							<Input
								register={register("movement_to")}
								type="string"
								label={
									movementToType === "storage"
										? "Куда (хранилище)"
										: "Куда (работник)"
								}
								width="240px"
								isAvailable={false}
								value={
									object?.movement_to ||
									movementToDisplayValue ||
									""
								}
							/>
							{!isReadOnly && movementToType && (
								<ButtonText
									type="button"
									textSize="var(--smallest-font-size)"
									textWeight={100}
									onClick={() => {
										setOpenMovementTo((prev) => !prev);
										setOpenMovementFrom(false);
									}}
								>
									{!openMovementTo
										? `Выбрать ${movementToType === "storage" ? "хранилище" : "работника"} ↓`
										: "Закрыть ↑"}
								</ButtonText>
							)}
						</div>

						{/* Поисковые блоки для movement_from и movement_to */}
						<div className={styles.objectForm__storageContainer}>
							{openMovementFrom &&
								movementFromType === "storage" && (
									<StorageSearch
										mode="replace"
										onSelect={
											handleSelectMovementFromStorage
										}
									/>
								)}
							{openMovementFrom &&
								movementFromType === "worker" && (
									<PersonalSearch
										mode="replace"
										onSelect={
											handleSelectMovementFromPerson
										}
									/>
								)}
							{openMovementTo && movementToType === "storage" && (
								<StorageSearch
									mode="replace"
									onSelect={handleSelectMovementToStorage}
								/>
							)}
							{openMovementTo && movementToType === "worker" && (
								<PersonalSearch
									mode="replace"
									onSelect={handleSelectMovementToPerson}
								/>
							)}
						</div>

						<Input
							register={register("post_number")}
							type="string"
							label="Номер посылки"
							width="240px"
							isAvailable={!isReadOnly}
						/>

						<Input
							register={register("sended_at")}
							type="date"
							label="Отправлено"
							width="240px"
							isAvailable={!isReadOnly}
						/>

						<Input
							register={register("arrived_at")}
							type="date"
							label="Прибыло"
							width="240px"
							isAvailable={!isReadOnly}
						/>
					</div>
				)}
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Гарантия</h2>

				<QualitySelect
					register={register("quality_status")}
					defaultValue={object?.quality_status}
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("warranty_started_at")}
					type="date"
					label="Начало гарантии"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("warranty_end_at")}
					type="date"
					label="Конец гарантии"
					width="240px"
					isAvailable={!isReadOnly}
				/>
			</div>

			{/* Кнопки */}
			{!isReadOnly && (
				<div className={styles.objectForm__buttonsContainer}>
					<ButtonText
						type="button"
						textWeight={100}
						textSize="var(--normal-font-size)"
						textColor="var(--grey-color)"
						onClick={() => {
							reset();

							setNewPlace(undefined);
							setNewPerson(undefined);
							setMovementFromStorage(undefined);
							setMovementFromPerson(undefined);
							setMovementToStorage(undefined);
							setMovementToPerson(undefined);

							setOpenTreeStorage(false);
							setOpenTreePersonal(false);
							setOpenMovementFrom(false);
							setOpenMovementTo(false);
						}}
					>
						Сбросить
					</ButtonText>

					<ButtonText
						type="submit"
						textSize="var(--normal-font-size)"
						textColor="var(--green-color)"
					>
						{mode === "save" ? "Сохранить" : "Создать"}
					</ButtonText>
				</div>
			)}
		</form>
	);
};

export default TechForm;
