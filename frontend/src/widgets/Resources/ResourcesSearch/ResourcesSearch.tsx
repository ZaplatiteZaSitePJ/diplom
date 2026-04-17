import type {
	CategoriesList,
	TransferStatus,
} from "@entities/Objects/types/baseObjects.type";
import TechSearch from "./TechSearch";
import DocsSearch from "./DocsSearch";
import SoftwareSearch from "./SoftwareSearch";
import type { FC } from "react";
import { useState } from "react";
import CategoriesChips from "@widgets/Resources/CategoriesChips/CategoriesChips";
import TechForm from "../ObjectsForm/ObjectForm/TechForm";
import DocsForm from "../ObjectsForm/ObjectForm/DocsForm";
import SoftwareForm from "../ObjectsForm/ObjectForm/SoftwareForm";

type ResourcesSearchProps = {
	callPlace?: Extract<TransferStatus, "worker" | "storage">;
	mode?: "search" | "create";
	name?: string;
};

const ResourcesPanel: FC<ResourcesSearchProps> = ({
	callPlace,
	name,
	mode = "search",
}) => {
	const [currentCategorie, setCurrentCategorie] =
		useState<CategoriesList>("Техника");

	const renderSearch = () => {
		switch (currentCategorie) {
			case "Техника":
				return <TechSearch callPlace={callPlace} name={name} />;

			case "Документы":
				return <DocsSearch />;

			case "Программное обеспечение":
				return <SoftwareSearch />;

			default:
				return null;
		}
	};

	const renderCreate = () => {
		switch (currentCategorie) {
			case "Техника":
				return <TechForm mode="create" />;

			case "Документы":
				return <DocsForm mode="create" />;

			case "Программное обеспечение":
				return <SoftwareForm mode="create" />;

			default:
				return null;
		}
	};

	return (
		<>
			<CategoriesChips
				currentCategorie={currentCategorie}
				setCurrentCategorie={setCurrentCategorie}
			/>

			{mode == "search" && renderSearch()}
			{mode == "create" && renderCreate()}
		</>
	);
};

export default ResourcesPanel;
