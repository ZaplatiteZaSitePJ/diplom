import CategoriesChips from "@widgets/Resources/CategoriesChips/CategoriesChips";
import styles from "./Resources.module.scss";
import { useState } from "react";
import type { CategoriesList } from "@entities/Objects/types/baseObjects.type";
import ResourcesSearch from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

export default function Categories() {
	const [currentCategorie, setCurrentCategorie] =
		useState<CategoriesList>("Техника");

	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Ресурсы</h1>
			</div>

			<div className={styles.main__mainContainer}>
				<CategoriesChips
					currentCategorie={currentCategorie}
					setCurrentCategorie={setCurrentCategorie}
				/>
				<ResourcesSearch currentCategorie={currentCategorie} />
			</div>
		</main>
	);
}
