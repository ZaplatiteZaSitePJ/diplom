import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState, useEffect } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/ObjectList";
import type { SoftwareItem } from "@entities/Objects/types/software.type";

const SoftwareSearch = () => {
	const { register, watch } = useForm<Partial<SoftwareItem>>();

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
						label="Вендор"
						register={register("vendor")}
						width="240px"
					/>
					<Input
						label="Название ПО"
						register={register("title")}
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
						label="Ключ лицензии"
						register={register("license_key")}
						width="240px"
						type="string"
					/>

					<Input
						label="Дата активации"
						register={register("started_at")}
						width="240px"
						type="date"
					/>

					<Input
						label="Дата истечения"
						register={register("expired_at")}
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

export default SoftwareSearch;
