import type { StorageType } from "@entities/Storages/types/storages.type";

const nameFilter = (array: StorageType[], name: string | undefined) => {
	if (array === undefined || name === undefined) {
		return array;
	}

	return [...array].filter((element) => {
		const elementName = element.storageName.toLowerCase();
		const nameSort = name.toLowerCase();

		return elementName.includes(nameSort);
	});
};

const cityFilter = (array: StorageType[], city: string | undefined) => {
	if (array === undefined || city === undefined) {
		return array;
	}

	return [...array].filter((element) => {
		const elementName = element.city.toLowerCase();
		const nameSort = city.toLowerCase();

		return elementName.includes(nameSort);
	});
};

export default function storageFiltration(
	name: string | undefined,
	city: string | undefined,
	array: StorageType[],
) {
	const filtredByName = nameFilter(array, name);
	const filtredByCity = cityFilter(filtredByName, city);
	return filtredByCity;
}
