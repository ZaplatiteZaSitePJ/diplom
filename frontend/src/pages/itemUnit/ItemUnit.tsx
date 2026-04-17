import { useEffect } from "react";
import { useParams, useSearchParams } from "react-router-dom";

import { useLazyGetTechByIdQuery } from "@app/api/items/tech/techAPI";
import { useLazyGetDocsByIdQuery } from "@app/api/items/docs/docsAPI";
import { useLazyGetSoftwareByIdQuery } from "@app/api/items/software/softwareAPI";

import EntitiesLayout from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";

import TechForm from "@widgets/Resources/ObjectsForm/ObjectForm/TechForm";
import DocsForm from "@widgets/Resources/ObjectsForm/ObjectForm/DocsForm";
import SoftwareForm from "@widgets/Resources/ObjectsForm/ObjectForm/SoftwareForm";

export default function ItemUnit() {
	const { id } = useParams<{ id: string }>();
	const [searchParams] = useSearchParams();
	const type = searchParams.get("type");

	const [triggerTech, { data: tech }] = useLazyGetTechByIdQuery();
	const [triggerDocs, { data: docs }] = useLazyGetDocsByIdQuery();
	const [triggerSoftware, { data: software }] = useLazyGetSoftwareByIdQuery();

	useEffect(() => {
		if (!id) return;

		if (type === "tech") {
			triggerTech(id);
		}

		if (type === "docs") {
			triggerDocs(id);
		}

		if (type === "software") {
			triggerSoftware(id);
		}
	}, [id, type, triggerTech, triggerDocs, triggerSoftware]);

	if (!type) return null;

	// TECH
	if (type === "tech" && tech) {
		return (
			<EntitiesLayout
				treeLink={`/tree/${tech.data.id}`}
				title={tech.data.universal_name}
				subTitle={tech.data.id}
				form={<TechForm object={tech.data} mode="save" />}
				entitie={tech.data}
			/>
		);
	}

	// DOCS
	if (type === "docs" && docs) {
		return (
			<EntitiesLayout
				treeLink={`/tree/${docs.data.id}`}
				title={docs.data.universal_name}
				subTitle={docs.data.id}
				form={<DocsForm object={docs.data} mode="save" />}
				entitie={docs.data}
			/>
		);
	}

	// SOFTWARE
	if (type === "software" && software) {
		return (
			<EntitiesLayout
				treeLink={`/tree/${software.data.id}`}
				title={software.data.universal_name}
				subTitle={software.data.id}
				form={<SoftwareForm object={software.data} mode="save" />}
				entitie={software.data}
			/>
		);
	}

	return null;
}
