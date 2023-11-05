import React from "React";

import type { GlobalProvider } from "@ladle/react";
import "../src/app/global.css";

export const Provider: GlobalProvider = ({ children, globalState }) => <>{children}</>;
