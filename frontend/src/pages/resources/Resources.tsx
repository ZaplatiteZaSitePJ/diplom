import { useState } from "react";

import ButtonText from "@shared/ui/ui-kit/buttons/ButtonText";
import Modal from "@features/modal/Modal";

import ResourcesSearch from "@widgets/Resources/ResourcesSearch/ResourcesSearch";
import ResourcesPanel from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

import styles from "./Resources.module.scss";

const Resources = () => {
	const [isAddModalOpen, setIsAddModalOpen] = useState(false);

	const handleOpenAddModal = () => {
		setIsAddModalOpen(true);
	};

	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<h1>Ресурсы</h1>

				<ButtonText
					textColor="var(--green-color)"
					textSize="36px"
					onClick={handleOpenAddModal}
				>
					+
				</ButtonText>
			</div>

			<div className={styles.main__mainContainer}>
				<ResourcesSearch />
			</div>

			{isAddModalOpen && (
				<Modal
					title="Добавить объект"
					onClose={() => setIsAddModalOpen(false)}
				>
					<ResourcesPanel mode="create" />
				</Modal>
			)}
		</main>
	);
};

export default Resources;
