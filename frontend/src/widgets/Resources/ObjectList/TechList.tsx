import styles from "./ObjectList.module.scss";
import {
	useGetMyTechsQuery,
	useGetTechsQuery,
} from "@app/api/items/tech/techAPI";
import {
	getQualityLabel,
	getTransferLabel,
	makeLightID,
} from "@entities/Objects/types/baseObjects.type";
import type { TechFilter, TechItem } from "@entities/Objects/types/tech.type";
import { lastPlaceFinder } from "@features/utils/lastPalceFinder";
import { useNavigate } from "react-router-dom";

const TechList = ({ filter, isMe }: { filter: TechFilter; isMe?: boolean }) => {
	const { data: allItems } = useGetTechsQuery(filter, {
		skip: isMe, // 👈 пропускаем если me
	});

	const { data: myItems } = useGetMyTechsQuery(undefined, {
		skip: !isMe, // 👈 пропускаем если НЕ me
	});

	const items = isMe ? myItems : allItems;
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
							<th>качество</th>
							<th>трансфер</th>
							<th>местонахождение</th>
						</tr>
					</thead>

					<tbody>
						{items.data.map((el: TechItem) => (
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
								<td>
									{getQualityLabel(el.quality_status) || "—"}
								</td>
								<td>{getTransferLabel(el.transfer_status)}</td>
								<td>
									{lastPlaceFinder(el.transfer_status, el)}
								</td>
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
