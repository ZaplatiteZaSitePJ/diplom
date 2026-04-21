import styles from "./ObjectList.module.scss";
import { useGetUsersQuery } from "@app/api/users/usersAPI";
import type { UserType } from "@entities/User/types/user.type";
import { useNavigate } from "react-router-dom";

const UserList = ({ filter }: { filter: Partial<UserType> }) => {
	const { data: users } = useGetUsersQuery(filter);
	const navigate = useNavigate();

	return (
		<div className={styles.objectList}>
			{users && users.data.length > 0 ? (
				<table className={styles.table}>
					<thead>
						<tr>
							<th>фамилия</th>
							<th>имя</th>
							<th>email</th>
							<th>должность</th>
							<th>грейд</th>
							<th>город</th>
						</tr>
					</thead>

					<tbody>
						{users.data.map((el: UserType) => (
							<tr
								key={el.id}
								onClick={() => navigate(`/personal/${el.id}`)}
								className={styles.row}
							>
								<td>{el.lastname}</td>
								<td>{el.name}</td>
								<td>{el.email}</td>
								<td>{el.post}</td>
								<td>{el.grade}</td>
								<td>{el.city}</td>
							</tr>
						))}
					</tbody>
				</table>
			) : (
				<p className={styles.notFound}>
					Подходящих пользователей не найдено
				</p>
			)}
		</div>
	);
};

export default UserList;
