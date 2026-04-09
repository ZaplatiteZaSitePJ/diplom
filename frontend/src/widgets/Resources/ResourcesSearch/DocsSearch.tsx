import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, useEffect } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/ObjectList";
import type { DocsItem } from "@entities/Objects/types/docs.type";

const DocsSearch = () => {
	const { register, watch } = useForm<Partial<DocsItem>>();

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
						label="Вид"
						register={register("doc_type")}
						width="240px"
					/>
					<Input
						label="Номер"
						register={register("doc_number")}
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
						label="Местонахождение"
						register={register("transfer_status")}
						width="240px"
						type="string"
					/>

					<Input
						label="Подписант"
						register={register("responsible_worker")}
						width="240px"
						type="string"
					/>

					<Input
						label="Дата подписи"
						register={register("full_signed_at")}
						width="240px"
						type="date"
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

export default DocsSearch;
