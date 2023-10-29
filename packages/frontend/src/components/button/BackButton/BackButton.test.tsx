import React from "react";
import { render } from "@testing-library/react";
import { Story1 } from "./BackButton.stories";
import { describe, expect, test } from "vitest";

describe("snapshot", () => {
  test("snapshot", () => {
    const { asFragment } = render(<Story1 />);
    expect(asFragment()).toMatchSnapshot();
  });
});
