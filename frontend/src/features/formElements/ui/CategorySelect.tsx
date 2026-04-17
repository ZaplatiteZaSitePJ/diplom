import { CustomSelect } from "@shared/ui/ui-kit/select/ui/Select";
import type { CustomSelectProps } from "@shared/ui/ui-kit";
import type { FC } from "react";
import { useGetCategoriesQuery } from "@app/api/items/tech/techAPI";

const CategorySelect: FC<CustomSelectProps & { id: string }> = ({
	register,
	defaultValue,
	label = "Категория",
	isAvailable,
	id,
}) => {
	const data = useGetCategoriesQuery(id);
	const categories = data?.data?.data;
	const options = categories?.map((el) => ({ value: el, label: el }));

	return (
		<CustomSelect
			options={options}
			defaultValue={defaultValue}
			register={register}
			label={label}
			isAvailable={isAvailable}
		/>
	);
};

export default CategorySelect;
