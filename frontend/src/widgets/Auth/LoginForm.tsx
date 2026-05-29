import { useForm } from "react-hook-form";
import EmailField from "@features/formElements/ui/AuthElements/EmailField";
import PasswordField from "@features/formElements/ui/AuthElements/PasswordField";
import styles from "./AuthForm.module.scss";
import { ButtonText } from "@shared/ui/ui-kit";
import cn from "classnames";
import { useNavigate } from "react-router-dom";
import type { UserType } from "@entities/User/types/user.type";
import { useLoginMutation } from "@app/api/auth/authAPI";
import { enqueueSnackbar } from "notistack";

export default function LoginForm() {
	const navigate = useNavigate();

	const { register, handleSubmit, reset, getValues } =
		useForm<Pick<UserType, "email" | "password">>();
	const [trigger, { isLoading, isError }] = useLoginMutation();

	const onLogin = async () => {
		if (
			getValues().email.trim() == "" ||
			getValues().password.trim() == ""
		) {
			enqueueSnackbar(
				"Ошибка! Поля почты и пароля не должны быть пустыми",
				{
					variant: "error",
					autoHideDuration: 7000,
				},
			);
			return;
		}

		try {
			const response = await trigger(getValues()).unwrap();

			const access = response?.data?.access;

			if (access) {
				localStorage.setItem("access", access);
				localStorage.setItem("role", response?.data?.role);

				enqueueSnackbar("Успешный вход!", {
					variant: "success",
					autoHideDuration: 3000,
				});

				reset();

				navigate("/profile", { replace: true });
			}
		} catch (err: any) {
			const statusCode = err?.status;

			if (statusCode === 404 || statusCode === 401) {
				enqueueSnackbar("Ошибка! Неправильный логин или пароль", {
					variant: "error",
					autoHideDuration: 7000,
				});
			} else {
				enqueueSnackbar(
					"Ошибка подкдлючения! \nПроблемы с интернетом или сервер временно недоступен",
					{
						variant: "error",
						autoHideDuration: 7000,
					},
				);
			}

			console.error(err);
		}
	};

	return (
		<>
			<form className={styles.authForm} onSubmit={handleSubmit(onLogin)}>
				<h1>Вход в систему</h1>

				<EmailField
					register={register("email")}
					// subContent={<p>Обязательное поле!</p>}
				/>
				<PasswordField
					register={register("password")}
					// subContent={<p>Обязательное поле!</p>}
				/>

				<div className={styles.authForm__buttonsContainer}>
					<ButtonText textColor="var(--white-color)" type="submit">
						Войти
					</ButtonText>
				</div>
			</form>

			<div
				className={cn(styles.authForm__animationContainer, {
					[styles.authForm__isLoading]: isLoading,
				})}
			></div>

			{/* {isError && (
				<p className={styles.errorResponse}>
					Ошибка! Неверная почта или пароль
				</p>
			)} */}
		</>
	);
}
