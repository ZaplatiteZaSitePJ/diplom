import { useForm, type SubmitHandler } from "react-hook-form";
import { type FC, useEffect } from "react";
import styles from "./ObjectForm.module.scss";
import { ButtonText, Input } from "@shared/ui/ui-kit";
// import StatusSelect from "@features/formElements/ui/StatusSelect";
import type { TechItem } from "@entities/Objects/types/tech.type";
import {
	useCreateTechMutation,
	useUpdateTechMutation,
} from "@app/api/items/tech/techAPI";
import { toISODate } from "@features/utils/dateConverter";
import { techDataForPush } from "./dataForPush";
import TransferSelect from "@features/formElements/ui/TransferSelect";
import QualitySelect from "@features/formElements/ui/QualitySelect";
import CategorySelect from "@features/formElements/ui/CategorySelect";
import { enqueueSnackbar } from "notistack";

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
	const { register, handleSubmit, reset } = useForm<TechItem>();
	const [triggerPost] = useCreateTechMutation();
	const [triggerPatch] = useUpdateTechMutation();

	useEffect(() => {
		if (object) {
			reset({
				...object,
				warranty_started_at: toISODate(object.warranty_started_at),
				warranty_end_at: toISODate(object.warranty_end_at),
			});
		}
	}, [object, reset]);

	const submitHandler: SubmitHandler<TechItem> = async (data) => {
		const formattedData = techDataForPush(data);
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

				<Input
					register={register("last_worker_email")}
					type="string"
					label="Последний пользователь"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("last_storage")}
					type="string"
					label="Последнее хранилище"
					width="240px"
					isAvailable={!isReadOnly}
				/>
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
						textWeight={100}
						textSize="var(--normal-font-size)"
						textColor="var(--grey-color)"
						onClick={() => reset()}
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
