import {
	useGetMySoftwaresQuery,
	useGetSoftwaresQuery,
} from "@app/api/items/software/softwareAPI";
import styles from "./ObjectList.module.scss";
import { makeLightID } from "@entities/Objects/types/baseObjects.type";
import type {
	SoftwareFilter,
	SoftwareItemPublic,
} from "@entities/Objects/types/software.type";
import { useNavigate } from "react-router-dom";

const SoftwareList = ({
	filter,
	isMe,
}: {
	filter: SoftwareFilter;
	isMe?: boolean;
}) => {
	const navigate = useNavigate();

	const { data: allItems } = useGetSoftwaresQuery(filter, {
		skip: isMe,
	});

	const { data: myItems } = useGetMySoftwaresQuery(undefined, {
		skip: !isMe,
	});

	const items = isMe ? myItems : allItems;

	return (
		<div className={styles.objectList}>
			{items && items.data.length > 0 ? (
				<table className={styles.table}>
					<thead>
						<tr>
							<th>id</th>
							<th>название</th>
							<th>категория</th>
							<th>лицензия</th>
							<th>последний владелец</th>
						</tr>
					</thead>

					<tbody>
						{items.data.map((el: SoftwareItemPublic) => (
							<tr
								key={el.id}
								onClick={() =>
									navigate(`/items/${el.id}?type=software`)
								}
								className={styles.row}
							>
								<td>{makeLightID(el.id)}</td>
								<td>{el.universal_name}</td>
								<td>{el.category || "—"}</td>
								<td>{el.license_key}</td>
								<td>{el.last_worker_email || "—"}</td>
							</tr>
						))}
					</tbody>
				</table>
			) : (
				<p className={styles.notFound}>
					Подходящих объектов не найдено
				</p>
			)}
		</div>
	);
};

export default SoftwareList;
