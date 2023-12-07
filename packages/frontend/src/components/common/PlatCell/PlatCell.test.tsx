import { render } from "@testing-library/react";
import React from "react";
import { describe, expect, test } from "vitest";

import { Story1, Story2, Story3, Story4, Story5 } from "./PlatCell.stories";

describe("snapshot", () => {
  test("snapshot1", () => {
    const { asFragment } = render(<Story1 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot2", () => {
    const { asFragment } = render(<Story2 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot3", () => {
    const { asFragment } = render(<Story3 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot4", () => {
    const { asFragment } = render(<Story4 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot5", () => {
    const { asFragment } = render(<Story5 />);
    expect(asFragment()).toMatchSnapshot();
  });
});
