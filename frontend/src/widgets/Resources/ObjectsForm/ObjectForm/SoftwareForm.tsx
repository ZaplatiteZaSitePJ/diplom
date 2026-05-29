import { useForm, type SubmitHandler } from "react-hook-form";
import { useEffect, type FC } from "react";
import styles from "./ObjectForm.module.scss";
import { ButtonText, Input } from "@shared/ui/ui-kit";

import type { SoftwareItemPublic } from "@entities/Objects/types/software.type";

import CategorySelect from "@features/formElements/ui/CategorySelect";
import { softwareDataForPush } from "./dataForPush";
import { toISODate } from "@features/utils/dateConverter";
import {
	useCreateSoftwareMutation,
	useUpdateSoftwareMutation,
} from "@app/api/items/software/softwareAPI";
import { enqueueSnackbar } from "notistack";

type SoftwareFormProps = {
	object?: SoftwareItemPublic;
	mode: "save" | "create";
	isReadOnly?: boolean;
	setReadOnly?: () => void;
};

const SoftwareForm: FC<SoftwareFormProps> = ({
	object,
	mode,
	isReadOnly = false,
	setReadOnly,
}) => {
	const { register, handleSubmit, reset } = useForm<SoftwareItemPublic>();

	const [triggerPost] = useCreateSoftwareMutation();
	const [triggerPatch] = useUpdateSoftwareMutation();

	useEffect(() => {
		if (object) {
			reset({
				...object,
				started_at: toISODate(object.started_at),
				expired_at: toISODate(object.expired_at),
			});
		}
	}, [object, reset]);

	const submitHandler: SubmitHandler<SoftwareItemPublic> = async (data) => {
		const formatted = softwareDataForPush(data);
		const id = formatted.id;

		try {
			if (mode === "create") {
				await triggerPost(formatted).unwrap();

				enqueueSnackbar(
					`${formatted.vendor} ${formatted.title} (${id}) успешно создано`,
					{
						variant: "success",
						autoHideDuration: 5000,
					},
				);
			}

			if (mode === "save") {
				await triggerPatch({
					id,
					body: formatted,
				}).unwrap();

				enqueueSnackbar(
					`${formatted.vendor} ${formatted.title} (${id}) успешно обновлено`,
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
					register={register("vendor")}
					label="Вендор"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("title")}
					label="Название"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("license_key")}
					label="Лицензия"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<CategorySelect
					id="2"
					register={register("category")}
					isAvailable={!isReadOnly}
					defaultValue={object?.category}
				/>

				<Input
					register={register("purchase_price")}
					label="Цена"
					width="240px"
					type="number"
					isAvailable={!isReadOnly}
				/>
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Принадлежность</h2>

				<Input
					register={register("last_worker_email")}
					label="Пользователь"
					width="240px"
					isAvailable={!isReadOnly}
				/>
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Срок лицензии</h2>

				<Input
					register={register("started_at")}
					label="Начало"
					type="date"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("expired_at")}
					label="Конец"
					type="date"
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

export default SoftwareForm;
