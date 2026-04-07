import { useForm } from "react-hook-form";
import EmailField from "@features/formElements/ui/AuthElements/EmailField";
import PasswordField from "@features/formElements/ui/AuthElements/PasswordField";
import styles from "./AuthForm.module.scss";
import { ButtonText } from "@shared/ui/ui-kit";
import cn from "classnames";
import { useNavigate } from "react-router-dom";
import type { UserType } from "@entities/User/types/user.type";
import { useLoginMutation } from "@app/api/auth/authAPI";

export default function LoginForm() {
	const navigate = useNavigate();

	const { register, handleSubmit, reset, getValues } =
		useForm<Pick<UserType, "email" | "password">>();
	const [trigger, { isLoading, isError }] = useLoginMutation();

	const onLogin = async () => {
		try {
			const { data } = await trigger(getValues());
			localStorage.setItem("access", data?.data?.access || "");
			reset();
			navigate("/storages", { replace: true });
		} catch {
			console.log("Ошибка");
		}
	};

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
					[styles.authForm__isLoading]: isLoading,
				})}
			></div>

			{isError && (
				<p className={styles.errorResponse}>
					Ошибка! Неверная почта или пароль
				</p>
			)}
		</>
	);
}
