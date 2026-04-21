import PersonalSearch from "@widgets/Resources/ResourcesSearch/PersonalSearch";
import styles from "./Personal.module.scss";
import ResourcesSearch from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

const Personal = () => {
	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Персонал</h1>
			</div>

			<div className={styles.main__mainContainer}>
				<PersonalSearch />
			</div>
		</main>
	);
};

export default Personal;
