import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, useEffect } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/ObjectList";
import type { MerchItem } from "@entities/Objects/types/merch.type";

const MerchSearch = () => {
	const { register, watch } = useForm<Partial<MerchItem>>();

	const [isWrapped, setWrapped] = useState<boolean>(false);
	const [request, setRequest] = useState<object[]>([]);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const formValues = watch();

	// useEffect(() => {
	// 	const fetchParsed = async () => {
	// 		try {
	// 			const parsed = await parseFormSearch(formValues);
	// 			setRequest(parsed);
	// 		} catch (error) {
	// 			console.error("parseFormSearch error:", error);
	// 		}
	// 	};

	// 	fetchParsed();
	// }, [JSON.stringify(formValues)]);

	return (
		<div className={styles.objectFormSearch}>
			<form
				className={styles.objectFormSearch__filter}
				style={
					isWrapped
						? { height: "600px", overflow: "auto" }
						: { height: "138px", overflow: "hidden" }
				}
			>
				<div className={styles.objectFormSearch__filterContainer}>
					<Input
						label="Артикул"
						register={register("id")}
						width="240px"
					/>
					<Input
						label="Название"
						register={register("title")}
						width="240px"
					/>
					<Input
						label="Хранилище"
						register={register("last_storage")}
						width="240px"
						type="string"
					/>
				</div>

				<div
					className={styles.objectFormSearch__filterContainer}
					style={
						isWrapped ? { display: "flex" } : { display: "none" }
					}
				>
					<Input
						label="Размер"
						register={register("size")}
						width="240px"
						type="number"
					/>

					<Input
						label="Цвет"
						register={register("color")}
						width="240px"
						type="string"
					/>

					<Input
						label="Местанахождение"
						register={register("transfer_status")}
						width="240px"
						type="string"
					/>
				</div>

				<div
					className={styles.objectFormSearch__additionalContainer}
					style={
						isWrapped ? { display: "block" } : { display: "none" }
					}
				>
					{" "}
				</div>
			</form>

			<button
				className={cn(styles.objectFormSearch__unwrapButton, {
					[styles.wrapped]: isWrapped,
				})}
				type="button"
				onClick={handleWrap}
			>
				↓
			</button>

			<div className={styles.objectFormSearch__objectListPlace}>
				{<ObjectList />}
			</div>
		</div>
	);
};

export default MerchSearch;
