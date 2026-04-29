import { useState, type FC, type ReactNode } from "react";
import styles from "./EntitiesLayout.module.scss";
import type { EntitiesLayoutProps } from "../types/entitiesLayout.type";
import { ButtonFilled, ButtonBorderred } from "@shared/ui/ui-kit";
import FormPanel from "@widgets/formPanel/FormPanel";
import { useNavigate } from "react-router-dom";
import Modal from "@features/modal/Modal";
import ResourcesPanel from "@widgets/Resources/ResourcesSearch/ResourcesSearch";
import ItemDeleteModal from "@widgets/DeleteModal/ObjectDelete";
import StorageDeleteModal from "@widgets/DeleteModal/StorageDelete";
import { useLogoutMutation } from "@app/api/auth/authAPI";

//
// 🔹 BaseLayout (общая обёртка)
//
type BaseLayoutProps = {
	title: string;
	subTitle?: string;
	treeLink?: string;
	children: ReactNode;
	actions?: ReactNode;
	statistic?: React.ReactNode;
};

const BaseLayout: FC<BaseLayoutProps> = ({
	title,
	subTitle,
	treeLink,
	children,
	actions,
}) => {
	const navigate = useNavigate();

	return (
		<main className={styles.main}>
			<div className={styles.main__topContainer}>
				<div className={styles.main__titlePlace}>
					<h1 className={styles.main__title}>{title}</h1>

					{subTitle && (
						<button
							type="button"
							className={styles.main__article}
							onClick={() => treeLink && navigate(treeLink)}
						>
							{subTitle}
						</button>
					)}
				</div>

				<div className={styles.main__buttonsBlock}>{actions}</div>
			</div>

			<div className={styles.main__mainContainer}>{children}</div>
		</main>
	);
};

//
// 👤 UserLayout
//
export const UserLayout: FC<EntitiesLayoutProps> = ({
	title,
	subTitle,
	form,
	isMe,
}) => {
	const [logout] = useLogoutMutation();
	const navigate = useNavigate();

	const [isAddModalOpen, setIsAddModalOpen] = useState(false);

	const handleLogout = async () => {
		if (!isMe) return;

		try {
			await logout().unwrap();
			localStorage.removeItem("access");
			navigate("/auth");
		} catch (e) {
			console.error(e);
		}
	};

	return (
		<>
			<BaseLayout
				title={title || ""}
				subTitle={subTitle}
				actions={
					<>
						<ButtonFilled onClick={() => setIsAddModalOpen(true)}>
							Записать объект
						</ButtonFilled>

						{isMe && (
							<ButtonBorderred
								borderColor="var(--red-color)"
								onClick={handleLogout}
								disabled={!isMe}
							>
								Выйти
							</ButtonBorderred>
						)}
					</>
				}
			>
				<FormPanel>{form}</FormPanel>
			</BaseLayout>

			{isAddModalOpen && (
				<Modal
					title="Добавить объект"
					onClose={() => setIsAddModalOpen(false)}
				>
					<ResourcesPanel mode="create" />
				</Modal>
			)}
		</>
	);
};

//
// 📦 StorageLayout
//
export const StorageLayout: FC<EntitiesLayoutProps> = ({
	title,
	subTitle,
	form,
	statistic,
	entitie,
}) => {
	const [isDeleting, setIsDeleting] = useState(false);
	const [isAddModalOpen, setIsAddModalOpen] = useState(false);

	return (
		<>
			<BaseLayout
				title={title || ""}
				subTitle={subTitle}
				actions={
					<>
						<ButtonFilled onClick={() => setIsAddModalOpen(true)}>
							Добавить объект
						</ButtonFilled>

						<ButtonBorderred
							borderColor="var(--red-color)"
							onClick={() => setIsDeleting(true)}
						>
							Удалить
						</ButtonBorderred>
					</>
				}
			>
				<>
					<FormPanel>{form}</FormPanel>

					<div className={styles.main__statisticPlace}>
						{statistic}
					</div>
				</>
			</BaseLayout>

			{isAddModalOpen && (
				<Modal
					title="Добавить объект"
					onClose={() => setIsAddModalOpen(false)}
				>
					<ResourcesPanel mode="create" />
				</Modal>
			)}

			{isDeleting && entitie && (
				<Modal
					title="Удаление"
					onClose={() => setIsDeleting(false)}
					bgColor="var(--red-color)"
				>
					<StorageDeleteModal
						storage={entitie}
						handleClose={() => setIsDeleting(false)}
					/>
				</Modal>
			)}
		</>
	);
};
//
// 🧱 ResourcesLayout
//
export const ResourcesLayout: FC<EntitiesLayoutProps> = ({
	title,
	subTitle,
	form,
	entitie,
}) => {
	const [isDeleting, setIsDeleting] = useState(false);

	return (
		<>
			<BaseLayout
				title={title || ""}
				subTitle={subTitle}
				actions={
					<ButtonBorderred
						borderColor="var(--red-color)"
						onClick={() => setIsDeleting(true)}
					>
						Удалить
					</ButtonBorderred>
				}
			>
				<FormPanel>{form}</FormPanel>
			</BaseLayout>

			{isDeleting && entitie && (
				<Modal
					title="Удаление"
					onClose={() => setIsDeleting(false)}
					bgColor="var(--red-color)"
				>
					<ItemDeleteModal
						object={entitie}
						handleClose={() => setIsDeleting(false)}
					/>
				</Modal>
			)}
		</>
	);
};
