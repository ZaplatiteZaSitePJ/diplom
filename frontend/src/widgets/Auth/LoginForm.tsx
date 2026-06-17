import { useForm } from "react-hook-form";
import EmailField from "@features/formElements/ui/AuthElements/EmailField";
import PasswordField from "@features/formElements/ui/AuthElements/PasswordField";
import styles from "./AuthForm.module.scss";
import { ButtonText } from "@shared/ui/ui-kit";
import cn from "classnames";
import { useNavigate } from "react-router-dom";
import type { UserType } from "@entities/User/types/user.type";
import { useLoginMutation } from "@app/api/auth/authAPI";
import { closeSnackbar, enqueueSnackbar } from "notistack";
import { useEffect, useRef } from "react";

export default function LoginForm() {
	const navigate = useNavigate();

	const { register, handleSubmit, reset, getValues } =
		useForm<Pick<UserType, "email" | "password">>();

	const [trigger, { isLoading, isError }] = useLoginMutation();

	const retryIntervalRef = useRef<NodeJS.Timeout | null>(null);
	const isRetryingRef = useRef(false);

	const stopRetry = () => {
		if (retryIntervalRef.current) {
			clearInterval(retryIntervalRef.current);
			retryIntervalRef.current = null;
		}
		isRetryingRef.current = false;
	};

	const loginRequest = async (values: {
		email: string;
		password: string;
	}) => {
		const response = await trigger(values).unwrap();

		const access = response?.data?.access;

		if (access) {
			closeSnackbar();
			localStorage.setItem("access", access);
			localStorage.setItem("role", response?.data?.role);

			enqueueSnackbar("Успешный вход!", {
				variant: "success",
				autoHideDuration: 3000,
			});

			stopRetry();
			reset();

			navigate("/profile", { replace: true });
		}
	};

	const startRetry = (fixedValues: { email: string; password: string }) => {
		if (isRetryingRef.current) return;

		isRetryingRef.current = true;

		retryIntervalRef.current = setInterval(async () => {
			try {
				await loginRequest(fixedValues);
			} catch (err: any) {
				const statusCode = err?.status;

				if (statusCode !== 422) {
					stopRetry();
				}
			}
		}, 10000);
	};

	const onLogin = async () => {
		const values = getValues();

		if (values.email.trim() === "" || values.password.trim() === "") {
			enqueueSnackbar(
				"Ошибка! Поля почты и пароля не должны быть пустыми",
				{ variant: "error", autoHideDuration: 7000 },
			);
			return;
		}

		try {
			await loginRequest(values);
		} catch (err: any) {
			const statusCode = err?.status;

			if (statusCode === 422) {
				closeSnackbar();
				enqueueSnackbar(
					`Первый вход. Следуйте инструкциям, отправленным на электронную почту ${values.email}`,
					{
						variant: "info",
						autoHideDuration: 70000,
					},
				);

				startRetry(values);
				return;
			}

			if (statusCode === 404 || statusCode === 401) {
				enqueueSnackbar("Ошибка! Неправильный логин или пароль", {
					variant: "error",
					autoHideDuration: 10000,
				});
			} else {
				enqueueSnackbar(
					"Ошибка подключения! Проблемы с интернетом или сервер недоступен",
					{ variant: "error", autoHideDuration: 7000 },
				);
			}
		}
	};

	useEffect(() => {
		return () => stopRetry();
	}, []);

	return (
		<>
			<form className={styles.authForm} onSubmit={handleSubmit(onLogin)}>
				<h1>Вход в систему</h1>

				<EmailField register={register("email")} />
				<PasswordField register={register("password")} />

				<div className={styles.authForm__buttonsContainer}>
					<ButtonText textColor="var(--white-color)" type="submit">
						Войти
					</ButtonText>
				</div>
			</form>

			<div
				className={cn(styles.authForm__animationContainer, {
					[styles.authForm__isLoading]:
						isLoading || isRetryingRef.current,
				})}
			/>
		</>
	);
}
