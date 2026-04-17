import { CustomSelect } from "@shared/ui/ui-kit/select/ui/Select";
import type { CustomSelectProps } from "@shared/ui/ui-kit";
import type { FC } from "react";
import { QualityStatusList } from "@entities/Objects/types/baseObjects.type";

const options = QualityStatusList;

const QualitySelect: FC<CustomSelectProps> = ({
	register,
	defaultValue,
	label = "Состояние",
	isAvailable,
}) => {
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

export default QualitySelect;
