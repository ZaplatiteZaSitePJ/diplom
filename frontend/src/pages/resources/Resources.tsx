import styles from "./Resources.module.scss";
import ResourcesSearch from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

const Resources = () => {
	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Ресурсы</h1>
			</div>

			<div className={styles.main__mainContainer}>
				<ResourcesSearch />
			</div>
		</main>
	);
};

export default Resources;
