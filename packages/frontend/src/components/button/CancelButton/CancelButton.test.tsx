import React from "react";
import { render } from "@testing-library/react";
import { Sample1 } from "./sample";
import { describe, expect, test } from "vitest";

test("snapshot", () => {
  const { asFragment } = render(<Sample1 />);
  expect(asFragment()).toMatchSnapshot();
});
