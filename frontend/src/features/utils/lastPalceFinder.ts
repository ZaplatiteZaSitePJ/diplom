import type { TransferStatus } from "@entities/Objects/types/baseObjects.type";

export function lastPlaceFinder(
	status: TransferStatus,
	el: { last_worker_email: string; last_storage: string },
): string {
	if (status === "worker") {
		return ` ${el.last_worker_email}`;
	}

	if (status === "storage") {
		return ` ${el.last_storage}`;
	}

	return "-";
}
