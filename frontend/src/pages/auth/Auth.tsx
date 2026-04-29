import styles from "./Auth.module.scss";
import LoginForm from "@widgets/Auth/LoginForm";
import Logo from "@shared/assets/site-logo.svg?react";

export default function Auth() {
	return (
		<main className={styles.main}>
			<Logo className={styles.main__logo} />

			<div className={styles.main__mainContainer}>
				<LoginForm />
			</div>
		</main>
	);
}
