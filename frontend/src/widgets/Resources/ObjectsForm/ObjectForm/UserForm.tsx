import { useForm, type SubmitHandler } from "react-hook-form";
import { type FC, useEffect } from "react";
import styles from "./ObjectForm.module.scss";
import { ButtonText, Input } from "@shared/ui/ui-kit";
import type { UserType } from "@entities/User/types/user.type";
import {
	useCreateUserMutation,
	useUpdateUserMutation,
} from "@app/api/users/usersAPI";
import { enqueueSnackbar } from "notistack";

type UserFormProps = {
	object?: UserType;
	mode: "save" | "create";
	isReadOnly?: boolean;
	setReadOnly?: () => void;
};

const UserForm: FC<UserFormProps> = ({
	object,
	mode,
	isReadOnly = false,
	setReadOnly,
}) => {
	const { register, handleSubmit, reset } = useForm<UserType>();

	const [triggerPost] = useCreateUserMutation();
	const [triggerPatch] = useUpdateUserMutation();

	useEffect(() => {
		if (object) {
			reset(object);
		}
	}, [object, reset]);

	const submitHandler: SubmitHandler<UserType> = async (data) => {
		const id = data.id;

		try {
			if (mode === "create") {
				await triggerPost(data).unwrap();

				enqueueSnackbar(
					`Пользователь ${data.lastname} ${data.name} (${id}) успешно создан`,
					{
						variant: "success",
						autoHideDuration: 5000,
					},
				);
			}

			if (mode === "save") {
				await triggerPatch({
					id,
					body: data,
				}).unwrap();

				enqueueSnackbar(
					`Пользователь ${data.lastname} ${data.name} (${id}) успешно обновлен`,
					{
						variant: "success",
						autoHideDuration: 5000,
					},
				);
			}

			setReadOnly?.();
		} catch (err: any) {
			console.error(err);

			if (err?.status === 400) {
				enqueueSnackbar("Ошибка! Введены некорректные данные", {
					variant: "error",
					autoHideDuration: 7000,
				});

				return;
			}

			if (err?.status === 409) {
				enqueueSnackbar("Ошибка! Пользователь уже существует", {
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
			{/* ========================= */}
			{/* Основная информация */}
			{/* ========================= */}
			<div className={styles.objectForm__additionalContainer}>
				<h2>Контактная информация</h2>

				<Input
					register={register("lastname")}
					label="Фамилия"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("name")}
					label="Имя"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("email")}
					label="Email"
					width="240px"
					isAvailable={!isReadOnly}
				/>
			</div>

			<div className={styles.objectForm__additionalContainer}>
				<h2>Работа</h2>

				<Input
					register={register("post")}
					label="Должность"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("grade")}
					label="Грейд"
					width="240px"
					isAvailable={!isReadOnly}
				/>

				<Input
					register={register("city")}
					label="Город"
					width="240px"
					isAvailable={!isReadOnly}
				/>
			</div>

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

export default UserForm;
