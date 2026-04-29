import { useEffect } from "react";
import { useParams, useSearchParams } from "react-router-dom";

import { useLazyGetTechByIdQuery } from "@app/api/items/tech/techAPI";
import { useLazyGetDocsByIdQuery } from "@app/api/items/docs/docsAPI";
import { useLazyGetSoftwareByIdQuery } from "@app/api/items/software/softwareAPI";

import { ResourcesLayout } from "@shared/layouts/entietiesLayout/ui/EntitiesLayout";

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
		if (!id || !type) return;

		if (type === "tech") triggerTech(id);
		if (type === "docs") triggerDocs(id);
		if (type === "software") triggerSoftware(id);
	}, [id, type, triggerTech, triggerDocs, triggerSoftware]);

	if (!type) return null;

	// 🔧 TECH
	if (type === "tech" && tech) {
		const item = tech.data;

		return (
			<ResourcesLayout
				title={item.universal_name}
				subTitle={item.id}
				form={<TechForm object={item} mode="save" />}
				entitie={item}
			/>
		);
	}

	// 📄 DOCS
	if (type === "docs" && docs) {
		const item = docs.data;

		return (
			<ResourcesLayout
				title={item.universal_name}
				subTitle={item.id}
				form={<DocsForm object={item} mode="save" />}
				entitie={item}
			/>
		);
	}

	// 💻 SOFTWARE
	if (type === "software" && software) {
		const item = software.data;

		return (
			<ResourcesLayout
				title={item.universal_name}
				subTitle={item.id}
				form={<SoftwareForm object={item} mode="save" />}
				entitie={item}
			/>
		);
	}

	return null;
}
