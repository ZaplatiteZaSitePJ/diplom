import type { CategoriesList } from "@entities/Objects/types/baseObjects.type";
import TechSearch from "./TechSearch";
import DocsSearch from "./DocsSearch";
import SoftwareSearch from "./SoftwareSearch";
import ConsumablesSearch from "./ConsumablesSearch";
import MerchSearch from "./MerchSearch";

const ResourcesSearch = ({
	currentCategorie,
}: {
	currentCategorie: CategoriesList;
}) => {
	switch (currentCategorie) {
		case "Техника":
			return <TechSearch />;

		case "Документы":
			return <DocsSearch />;

		case "Программное обеспечение":
			return <SoftwareSearch />;

		case "Расходные ресурсы":
			return <ConsumablesSearch />;

		case "Брендовые вещи":
			return <MerchSearch />;

		default:
			return null;
	}
};

export default ResourcesSearch;
