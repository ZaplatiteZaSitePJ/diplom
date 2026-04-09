import styles from "./ObjectList.module.scss";
import ObjectItem from "../ObjectItem/ObjectItem";
import { useGetTechsQuery } from "@app/api/items/tech/techAPI";

const ObjectList = () => {
	const { data: items } = useGetTechsQuery();
	console.log(items?.data);
	return (
		<div className={styles.objectList}>
			{items && items.data.length > 0 ? (
				items.data.map((el) => <ObjectItem object={el} />)
			) : (
				<p>Выберите способ фильрации объектов</p>
			)}
		</div>
	);
};

export default ObjectList;
