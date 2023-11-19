import { render } from "@testing-library/react";
import React from "react";
import { describe, expect, test } from "vitest";

import { Sample1 } from "./CancelButton.stories";

describe("snapshot", () => {
  test("snapshot", () => {
    const { asFragment } = render(<Sample1 />);
    expect(asFragment()).toMatchSnapshot();
  });
});
