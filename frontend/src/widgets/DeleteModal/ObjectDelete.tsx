import styles from "./DeleteModal.module.scss";
import type { FC } from "react";
import { ButtonText } from "@shared/ui/ui-kit";
import { useNavigate } from "react-router-dom";
import type { BaseObjectType } from "@entities/Objects/types/baseObjects.type";
import { useDeleteTechMutation } from "@app/api/items/tech/techAPI";
import { enqueueSnackbar } from "notistack";

type ObjectDeleteModalProps = {
	object: BaseObjectType;
	handleClose?: () => void;
};

const ItemDeleteModal: FC<ObjectDeleteModalProps> = ({
	object,
	handleClose,
}) => {
	const navigate = useNavigate();
	const [triggerDelete] = useDeleteTechMutation();

	const handleDelete = async () => {
		try {
			await triggerDelete(object.id as string).unwrap();
			enqueueSnackbar(`Объект ${object.id} Успешно удален!`, {
				variant: "success",
				autoHideDuration: 5000,
			});
			navigate(-1);
		} catch (error: any) {
			if (error?.status === "FETCH_ERROR") {
				enqueueSnackbar(
					`Ошибка! Объект ${object.id} не удалось удалить. Попробуйте позже.`,
					{
						variant: "error",
						autoHideDuration: 5000,
					},
				);
				console.log(error);
			}
		}
	};
	return (
		<div className={styles.deleteModule}>
			<p className={styles.deleteModule__info}>
				{object.universal_name} ({object.id})
			</p>

			<div className={styles.deleteModule__buttonContainer}>
				<ButtonText
					type="button"
					textSize="var(--normal-font-size)"
					textColor="var(--white-color)"
					onClick={handleClose}
				>
					Отмена
				</ButtonText>

				<ButtonText
					textWeight={100}
					textSize="var(--normal-font-size)"
					textColor="var(--white-color)"
					type="button"
					onClick={handleDelete}
				>
					Удалить
				</ButtonText>
			</div>
		</div>
	);
};

export default ItemDeleteModal;
