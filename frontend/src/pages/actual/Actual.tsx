import ResourcesPanel from "@widgets/Resources/ResourcesSearch/ResourcesSearch";
import styles from "./Actual.module.scss";
import StorageSearch from "@widgets/Storages/StorageSearch/StorageSearch";
import TechSearch from "@widgets/Resources/ResourcesSearch/TechSearch";

const Actual = () => {
	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Актуальное</h1>
			</div>

			<div className={styles.main__mainContainer}>
				<h2 className={styles.main__subTitle}>Заполненные хранилища</h2>
				<div className={styles.main__content}>
					<StorageSearch isFull />
				</div>

				<h2 className={styles.main__subTitle}>Сломанная техника</h2>
				<TechSearch isBroken={true} />
			</div>
		</main>
	);
};

export default Actual;
