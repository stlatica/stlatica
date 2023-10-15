import React from "react";
import { render } from "@testing-library/react";
import { describe, expect, test } from "vitest";
import { Story1, Story2, Story3 } from "./FollowButton.stories";

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
});
