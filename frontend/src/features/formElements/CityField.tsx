import type { InputProps } from "@shared/ui/ui-kit";
import { Input } from "@shared/ui/ui-kit";
import type { FC } from "react";

const CityField: FC<InputProps> = ({
	register,
	sx,
	subContent,
	label = "Город",
	isAvailable,
	width,
}) => {
	return (
		<Input
			label={label}
			register={register}
			width={width}
			fullWidth={width ? false : true}
			sx={sx}
			subContent={subContent}
			isAvailable={isAvailable}
		/>
	);
};

export default CityField;
