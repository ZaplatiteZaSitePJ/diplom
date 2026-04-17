import { useForm, type SubmitHandler } from "react-hook-form";
import { type FC, useEffect } from "react";
import styles from "./ObjectForm.module.scss";
import { ButtonText, Input } from "@shared/ui/ui-kit";
import { docsDataForPush } from "./dataForPush";
import TransferSelect from "@features/formElements/ui/TransferSelect";
import type { DocsItem } from "@entities/Objects/types/docs.type";
import {
	useCreateDocsMutation,
	useUpdateDocsMutation,
} from "@app/api/items/docs/docsAPI";
import CategorySelect from "@features/formElements/ui/CategorySelect";

type DocsFormProps = {
	object?: DocsItem;
	mode: "save" | "create";
	isReadOnly?: boolean;
	setReadOnly?: () => void;
};

const DocsForm: FC<DocsFormProps> = ({
	object,
	mode,
	isReadOnly = false,
	setReadOnly,
}) => {
	const { register, handleSubmit, reset } = useForm<DocsItem>();
	const [triggerPost] = useCreateDocsMutation();
	const [triggerPatch] = useUpdateDocsMutation();

	useEffect(() => {
		if (object) {
			reset({
				...object,
			});
		}
	}, [object, reset]);

	const submitHandler: SubmitHandler<DocsItem> = async (data) => {
		const formattedData = docsDataForPush(data);
		console.log(formattedData);
		const id = formattedData.id;

		try {
			if (mode == "create") {
				await triggerPost(formattedData);
			}

			if (mode == "save") {
				await triggerPatch({ id: id, body: formattedData });
			}

			setReadOnly?.();
		} catch {
			console.log("не отправилось");
		}
	};

	return (
		<form
			className={styles.objectForm}
			onSubmit={handleSubmit(submitHandler)}
		>
			<div className={styles.objectForm__additionalContainer}>
				<h2>Основная информация</h2>

				<CategorySelect
					id={"1"}
					register={register("category")}
					isAvailable={!isReadOnly}
					defaultValue={object?.category}
				/>

				<Input
					register={register("doc_number")}
					label="Номер документа"
					width="240px"
					isAvailable={!isReadOnly}
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
					register={register("last_storage")}
					label="Последее хранилище"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("responsible_worker_email")}
					label="Ответственный работник (email)"
					width="240px"
					isAvailable={!isReadOnly}
				/>
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Подписание</h2>
				<Input
					register={register("needed_signs")}
					type="number"
					label="Необходимо подписей"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("quality_status")}
					type="string"
					label="Получено подписей"
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

export default DocsForm;
