import { render } from "@testing-library/react";
import React from "react";
import { describe, expect, test } from "vitest";

import { Story1 } from "./BackButton.stories";

describe("snapshot", () => {
  test("snapshot", () => {
    const { asFragment } = render(<Story1 />);
    expect(asFragment()).toMatchSnapshot();
  });
});
