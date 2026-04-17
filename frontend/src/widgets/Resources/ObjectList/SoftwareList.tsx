import { useGetSoftwaresQuery } from "@app/api/items/software/softwareAPI";
import styles from "./ObjectList.module.scss";
import { makeLightID } from "@entities/Objects/types/baseObjects.type";
import type { SoftwareFilter } from "@entities/Objects/types/software.type";
import { useNavigate } from "react-router-dom";

const SoftwareList = ({ filter }: { filter: SoftwareFilter }) => {
	const { data: items } = useGetSoftwaresQuery(filter);
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
							<th>лицензия</th>
							<th>последний владелец</th>
						</tr>
					</thead>

					<tbody>
						{items.data.map((el) => (
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
