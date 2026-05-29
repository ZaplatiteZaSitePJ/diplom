import type { BaseObjectType } from "@entities/Objects/types/baseObjects.type";
import type { StorageType } from "@entities/Storages/types/storages.type";
import type { ReactNode } from "react";

export type EntitiesLayoutProps = {
	title?: string;
	subTitle?: string;
	deleteAction?: () => void;
	treeLink?: string;
	form?: ReactNode;
	statistic?: ReactNode;
	entitie?: StorageType | ObjectType;
	callPlace?: "storage" | "object" | "user";
	isMe?: boolean;
	readOnly?: boolean;
};
