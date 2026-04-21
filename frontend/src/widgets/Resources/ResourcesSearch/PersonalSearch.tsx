import styles from "./FormSearch.module.scss";
import { useForm } from "react-hook-form";
import { Input } from "@shared/ui/ui-kit";
import { useState } from "react";
import cn from "classnames";
import ObjectList from "../ObjectList/PersonalList";
import type { UserType } from "@entities/User/types/user.type";

type UserFilter = Partial<UserType>;

const PersonalSearch = () => {
	const { register, watch } = useForm<UserFilter>();

	const [isWrapped, setWrapped] = useState<boolean>(false);

	const handleWrap = () => {
		setWrapped((prev) => !prev);
	};

	const rawValues = watch();

	const filter: UserFilter = {
		id: rawValues.id ?? undefined,
		name: rawValues.name ?? undefined,
		lastname: rawValues.lastname ?? undefined,
		email: rawValues.email ?? undefined,
		post: rawValues.post ?? undefined,
		grade: rawValues.grade ?? undefined,
		city: rawValues.city ?? undefined,
	};

	console.log("Form changed:", filter);

	return (
		<div className={styles.objectFormSearch}>
			<form
				className={styles.objectFormSearch__filter}
				style={
					isWrapped
						? {
								height: "368px",
								overflow: "auto",
								boxShadow: "0 12px 30px rgba(0, 0, 0, 0.75)",
							}
						: { height: "138px", overflow: "hidden" }
				}
			>
				<div className={styles.objectFormSearch__filterContainer}>
					<Input label="ID" register={register("id")} width="240px" />
					<Input
						label="Имя"
						register={register("name")}
						width="240px"
					/>
					<Input
						label="Фамилия"
						register={register("lastname")}
						width="240px"
					/>
				</div>

				<div
					className={styles.objectFormSearch__filterContainer}
					style={
						isWrapped ? { display: "flex" } : { display: "none" }
					}
				>
					<Input
						label="Email"
						register={register("email")}
						width="240px"
					/>
					<Input
						label="Должность"
						register={register("post")}
						width="240px"
					/>
					<Input
						label="Грейд"
						register={register("grade")}
						width="240px"
					/>
					<Input
						label="Город"
						register={register("city")}
						width="240px"
					/>
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
				<ObjectList filter={filter} />
			</div>
		</div>
	);
};

export default PersonalSearch;
