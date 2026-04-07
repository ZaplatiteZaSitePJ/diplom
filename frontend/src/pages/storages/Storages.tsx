import ButtonText from "@shared/ui/ui-kit/buttons/ButtonText";
import styles from "./Storages.module.scss";
import Modal from "@features/modal/Modal";
import { useState } from "react";
import { useRevalidator } from "react-router-dom";
import StorageSearch from "@widgets/Storages/StorageSearch/StorageSearch";

export default function Storages() {
	const [isCreation, setIsCreation] = useState(false);
	const { revalidate } = useRevalidator();

	const handleCreateStorage = () => {
		setIsCreation(true);
	};

	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Хранилища</h1>
				<ButtonText
					textColor="var(--green-color)"
					textSize="36px"
					onClick={handleCreateStorage}
				>
					+
				</ButtonText>
			</div>

			<div className={styles.main__mainContainer}>
				<StorageSearch />
			</div>

			{isCreation && (
				<Modal
					title="Создание Хранилища"
					onClose={() => setIsCreation(false)}
				>
					<></>
					{/* <StorageForm
						mode="create"
						handleClose={() => {
							setIsCreation(false);
							revalidate();
						}}
					/> */}
				</Modal>
			)}
		</main>
	);
}
