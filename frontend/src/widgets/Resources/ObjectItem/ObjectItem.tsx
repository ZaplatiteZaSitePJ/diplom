import { type FC } from "react";
import styles from "./ObjectItem.module.scss";
import type { BaseObjectType } from "@entities/Objects/types/baseObjects.type";
import { Link } from "react-router-dom";

type ObjectItemProps = {
	object: BaseObjectType;
};

const ObjectItem: FC<ObjectItemProps> = ({ object }) => {
	return (
		<div className={styles.objectItem}>
			<Link to={`/objects/${object.id}`}>
				<div
					style={{
						width: "240px",
						height: "120px",
						borderRadius: "16px",
					}}
				></div>
				<p>{object.universal_name ?? `${object.universal_name}`}</p>
			</Link>
		</div>
	);
};

export default ObjectItem;
