import { useParams } from "react-router-dom";
import styles from "./UserUnit.module.scss";
import {
	useGetMeByIdQuery,
	useGetUserByIdQuery,
} from "@app/api/users/usersAPI";

import { UserLayout } from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";
import UserForm from "@widgets/Resources/ObjectsForm/ObjectForm/UserForm";
import ResourcesPanel from "@widgets/Resources/ResourcesSearch/ResourcesSearch";

type Props = {
	callPlace?: "user" | "me";
};

export default function UserUnit({ callPlace = "user" }: Props) {
	const { id } = useParams<{ id: string }>();

	const isMe = callPlace === "me";

	const { data: userData } = useGetUserByIdQuery(id || "", {
		skip: isMe,
	});

	const { data: meData } = useGetMeByIdQuery(undefined, {
		skip: !isMe,
	});

	const data = isMe ? meData : userData;

	if (!data) return null;

	const user = data.data;

	return (
		<div className={styles.page}>
			<UserLayout
				title={
					isMe
						? `Ваш профиль: ${user.name} ${user.lastname}`
						: `${user.name} ${user.lastname}`
				}
				subTitle={user.email}
				form={<UserForm object={user} mode="save" />}
				entitie={user}
				isMe={isMe}
				readOnly
			/>

			<div className={styles.page__resources}>
				<h2>Ресурсы в пользовании</h2>
				<ResourcesPanel
					callPlace={callPlace === "user" ? "worker" : "me"}
					name={user.email}
				/>
			</div>
		</div>
	);
}
