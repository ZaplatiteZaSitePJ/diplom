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
import type { TechItem } from "@entities/Objects/types/tech.type";
import type { DocsItem } from "@entities/Objects/types/docs.type";

type ResourcesSearchProps = {
	callPlace?: Extract<TransferStatus, "worker" | "storage"> | "me";
	mode?: "search" | "create";
	name?: string;
	storageName?: string;
};

const ResourcesPanel: FC<ResourcesSearchProps> = ({
	callPlace,
	name,
	mode = "search",
	storageName = undefined,
}) => {
	const [currentCategorie, setCurrentCategorie] =
		useState<CategoriesList>("Техника");

	// 🔥 формируем constFilter ОДИН раз
	const constFilter = {
		last_worker_email: callPlace === "worker" ? name : undefined,
		last_worker: callPlace === "worker" ? name : undefined,
		last_storage: callPlace === "storage" ? name : undefined,
		transfer_status:
			callPlace === "worker"
				? "worker"
				: callPlace === "storage"
					? "storage"
					: undefined,
	};

	const renderSearch = () => {
		switch (currentCategorie) {
			case "Техника":
				return (
					<TechSearch
						constFilter={constFilter}
						isMe={callPlace === "me"}
					/>
				);

			case "Документы":
				return (
					<DocsSearch
						constFilter={constFilter}
						isMe={callPlace === "me"}
					/>
				);

			case "Программное обеспечение":
				return (
					<SoftwareSearch
						constFilter={constFilter}
						isMe={callPlace === "me"}
					/>
				);

			default:
				return null;
		}
	};

	const renderCreate = () => {
		switch (currentCategorie) {
			case "Техника":
				return (
					<TechForm
						mode="create"
						object={
							callPlace == "storage"
								? ({ last_storage: storageName } as TechItem)
								: undefined
						}
					/>
				);

			case "Документы":
				return (
					<DocsForm
						mode="create"
						object={
							callPlace == "storage"
								? ({ last_storage: storageName } as DocsItem)
								: undefined
						}
					/>
				);

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

			{mode === "search" && renderSearch()}
			{mode === "create" && renderCreate()}
		</>
	);
};

export default ResourcesPanel;
