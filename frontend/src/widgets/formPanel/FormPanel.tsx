import {
	cloneElement,
	isValidElement,
	useState,
	type FC,
	type ReactNode,
} from "react";
import styles from "./FormPanel.module.scss";
import cn from "classnames";

type FormPanelProps = {
	children: ReactNode;
	readOnlyProp?: boolean;
};

const FormPanel: FC<FormPanelProps> = ({ children, readOnlyProp }) => {
	const [isReadOnly, setIsReadOnly] = useState<boolean>(true);

	const switchMode = (mode: boolean) => {
		if (!readOnlyProp) {
			setIsReadOnly(mode);
		}
	};

	return (
		<div className={styles.formPanel}>
			<div className={styles.formPanel__buttonsContainer}>
				<div className={styles.formPanel__switch}>
					<button
						type="button"
						onClick={() => switchMode(true)}
						className={cn({
							[styles.active]: isReadOnly,
						})}
					>
						чтение
					</button>

					{!readOnlyProp && (
						<>
							<span>/</span>
							<button
								type="button"
								onClick={() => switchMode(false)}
								className={cn({
									[styles.active]: !isReadOnly,
								})}
							>
								изменение
							</button>
						</>
					)}
				</div>
			</div>
			<div className={styles.formPanel__formPlace}>
				{isValidElement<{
					isReadOnly?: boolean;
					setReadOnly?: typeof setIsReadOnly;
				}>(children)
					? cloneElement(children, {
							isReadOnly,
							setReadOnly: () => setIsReadOnly(true),
						})
					: children}
			</div>
		</div>
	);
};

export default FormPanel;
