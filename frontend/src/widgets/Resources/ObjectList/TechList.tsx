import styles from "./ObjectList.module.scss";
import { useGetTechsQuery } from "@app/api/items/tech/techAPI";
import {
	getTransferLabel,
	makeLightID,
} from "@entities/Objects/types/baseObjects.type";
import type { TechFilter } from "@entities/Objects/types/tech.type";
import { useNavigate } from "react-router-dom";

const TechList = ({ filter }: { filter: TechFilter }) => {
	const { data: items } = useGetTechsQuery(filter);
	const navigate = useNavigate();

	return (
		<div className={styles.objectList}>
			{items && items.data.length > 0 ? (
				<table className={styles.table}>
					<thead>
						<tr>
							<th>id</th>
							<th>название</th>
							<th>категория</th>
							<th>трансфер</th>
							<th>последнее хранилище</th>
							<th>последний владелец</th>
						</tr>
					</thead>

					<tbody>
						{items.data.map((el) => (
							<tr
								key={el.id}
								onClick={() =>
									navigate(`/items/${el.id}?type=tech`)
								}
								className={styles.row}
							>
								<td>{makeLightID(el.id)}</td>
								<td>{el.universal_name}</td>
								<td>{el.category || "—"}</td>
								<td>{getTransferLabel(el.transfer_status)}</td>
								<td>{el.last_storage || "—"}</td>
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

export default TechList;
