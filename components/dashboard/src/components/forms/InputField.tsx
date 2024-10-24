/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import classNames from "classnames";
import { FunctionComponent, memo, ReactNode } from "react";
import { InputFieldHint } from "./InputFieldHint";

type Props = {
    label?: ReactNode;
    id?: string;
    hint?: ReactNode;
    error?: ReactNode;
    className?: string;
};

export const InputField: FunctionComponent<Props> = memo(({ label, id, hint, error, className, children }) => {
    return (
        <div className={classNames("mt-4 flex flex-col space-y-2", className)}>
            {label && (
                <label
                    className={classNames(
                        "text-md font-semibold dark:text-gray-400",
                        error ? "text-red-600" : "text-gray-600",
                    )}
                    htmlFor={id}
                >
                    {label}
                </label>
            )}
            {children}
            {error && <span className="text-red-500 text-sm">{error}</span>}
            {hint && <InputFieldHint>{hint}</InputFieldHint>}
        </div>
    );
});
