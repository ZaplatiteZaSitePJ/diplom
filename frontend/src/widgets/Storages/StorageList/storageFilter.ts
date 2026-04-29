import type { StorageType } from "@entities/Storages/types/storages.type";

const nameFilter = (array: StorageType[], name: string | undefined) => {
	if (!array || !name) return array;

	return array.filter((element) => {
		const elementName = element.storageName.toLowerCase();
		const nameSort = name.toLowerCase();

		return elementName.includes(nameSort);
	});
};

const cityFilter = (array: StorageType[], city: string | undefined) => {
	if (!array || !city) return array;

	return array.filter((element) => {
		const elementCity = element.city.toLowerCase();
		const citySort = city.toLowerCase();

		return elementCity.includes(citySort);
	});
};

const capacityFilter = (
	array: StorageType[],
	requiredCells: number | undefined,
) => {
	if (!array || requiredCells === undefined) return array;

	return array.filter((element) => {
		const freeCells = element.capacity - element.occupied_cells;
		return freeCells >= requiredCells;
	});
};

const fullFilter = (array: StorageType[], isFull: boolean | undefined) => {
	if (!array || isFull === false) return array;

	return array.filter(
		(element) => 95 <= (element.occupied_cells / element.capacity) * 100,
	);
};

export default function storageFiltration(
	name: string | undefined,
	city: string | undefined,
	array: StorageType[],
	requiredCells?: number,
	isFull?: boolean,
) {
	const filteredByName = nameFilter(array, name);
	const filteredByCity = cityFilter(filteredByName, city);
	const filteredByCapacity = capacityFilter(filteredByCity, requiredCells);
	const filteredByFullness = fullFilter(filteredByCapacity, isFull);

	return filteredByFullness;
}
