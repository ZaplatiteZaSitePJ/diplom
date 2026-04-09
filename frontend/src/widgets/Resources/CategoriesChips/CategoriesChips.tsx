import type { CategoriesList } from "@entities/Objects/types/baseObjects.type";
import styles from "./CategoriesChips.module.scss";
import cn from "classnames";

const CategoriesChips = ({
	currentCategorie,
	setCurrentCategorie,
}: {
	currentCategorie: CategoriesList;
	setCurrentCategorie: (newValue: CategoriesList) => void;
}) => {
	const categories: CategoriesList[] = [
		"Техника",
		"Программное обеспечение",
		"Документы",
	];

	return (
		<ul className={styles.chips}>
			{categories.map((category) => {
				const isActive = currentCategorie === category;

				return (
					<li key={category} className={styles.chips__element}>
						<button
							className={cn(styles.chips__button, {
								[styles.chips__active]: isActive,
							})}
							onClick={() => setCurrentCategorie(category)}
						>
							{category}
						</button>
					</li>
				);
			})}
		</ul>
	);
};

export default CategoriesChips;
