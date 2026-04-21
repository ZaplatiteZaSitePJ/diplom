import { useParams } from "react-router-dom";
import styles from "./UserUnit.module.scss";
import { useGetUserByIdQuery } from "@app/api/users/usersAPI";

import EntitiesLayout from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";
import UserForm from "@widgets/Resources/ObjectsForm/ObjectForm/UserForm";
import ResourcesPanel from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

export default function UserUnit() {
	const { id } = useParams<{ id: string }>();

	const { data: user } = useGetUserByIdQuery(id || "");

	if (!user) return null;

	return (
		<div className={styles.page}>
			<EntitiesLayout
				treeLink={`/tree/${user.data.id}`}
				title={`${user.data.name} ${user.data.lastname}`}
				subTitle={user.data.email}
				form={<UserForm object={user.data} mode="save" />}
				entitie={user.data}
			/>

			<div className={styles.page__resources}>
				<h2>Ресурсы в пользовании</h2>
				<ResourcesPanel callPlace="worker" name={user?.data?.email} />
			</div>
		</div>
	);
}
