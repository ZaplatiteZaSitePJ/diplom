import {
	useGetDocssQuery,
	useGetMyDocsQuery,
} from "@app/api/items/docs/docsAPI";
import styles from "./ObjectList.module.scss";
import {
	getTransferLabel,
	makeLightID,
} from "@entities/Objects/types/baseObjects.type";
import type { DocsFilter, DocsItem } from "@entities/Objects/types/docs.type";
import { useNavigate } from "react-router-dom";

const DocsList = ({ filter, isMe }: { filter: DocsFilter; isMe?: boolean }) => {
	const navigate = useNavigate();

	const { data: allItems } = useGetDocssQuery(filter, {
		skip: isMe,
	});

	const { data: myItems } = useGetMyDocsQuery(undefined, {
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
							<th>трансфер</th>
							<th>последнее хранилище</th>
							<th>подписант</th>
						</tr>
					</thead>

					<tbody>
						{items.data.map((el: DocsItem) => (
							<tr
								key={el.id}
								onClick={() =>
									navigate(`/items/${el.id}?type=docs`)
								}
								className={styles.row}
							>
								<td>{makeLightID(el.id)}</td>
								<td>{el.universal_name}</td>
								<td>{getTransferLabel(el.transfer_status)}</td>
								<td>{el.last_storage || "—"}</td>
								<td>{el.responsible_worker_email || "—"}</td>
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

export default DocsList;
