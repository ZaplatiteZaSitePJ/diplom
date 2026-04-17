import { CustomSelect } from "@shared/ui/ui-kit/select/ui/Select";
import type { CustomSelectProps } from "@shared/ui/ui-kit";
import type { FC } from "react";
import { TransferStatusList } from "@entities/Objects/types/baseObjects.type";

const options = TransferStatusList;

const TransferSelect: FC<CustomSelectProps> = ({
	register,
	defaultValue,
	label = "Трансфер",
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

export default TransferSelect;
