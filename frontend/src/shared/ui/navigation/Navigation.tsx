import { NavLink } from "react-router-dom";
import styles from "./Navigation.module.scss";
import Logo from "@shared/assets/site-logo.svg?react";
import cn from "classnames";

export default function Navigation() {
	return (
		<div className={styles.container}>
			<Logo className={styles.logo} />

			<nav className={styles.navigation}>
				<NavLink
					to={"/categories"}
					className={({ isActive }) =>
						cn(styles.navigation__pages, {
							[styles.navigation__active]: isActive,
						})
					}
				>
					Категории
				</NavLink>

				<NavLink
					to={"/storages"}
					className={({ isActive }) =>
						cn(styles.navigation__pages, {
							[styles.navigation__active]: isActive,
						})
					}
				>
					Хранилища
				</NavLink>

				<NavLink
					to={"/personal"}
					className={({ isActive }) =>
						cn(styles.navigation__pages, {
							[styles.navigation__active]: isActive,
						})
					}
				>
					Персонал
				</NavLink>

				<NavLink
					to={"/actual"}
					className={({ isActive }) =>
						cn(styles.navigation__pages, {
							[styles.navigation__active]: isActive,
						})
					}
				>
					Актуальное
				</NavLink>

				<NavLink
					to={"/messanger"}
					className={({ isActive }) =>
						cn(styles.navigation__pages, {
							[styles.navigation__active]: isActive,
						})
					}
				>
					Сообщения
				</NavLink>
			</nav>
		</div>
	);
}
