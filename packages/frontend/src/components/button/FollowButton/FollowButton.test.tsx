import React from "react";
import renderer from "react-test-renderer";
import { Story1, Story2, Story3 } from "./FollowButton.stories";

describe("snapshot", () => {
  test("snapshot1", () => {
    const tree = renderer.create(<Story1 />).toJSON();
    expect(tree).toMatchSnapshot();
  });
  test("snapshot2", () => {
    const tree = renderer.create(<Story2 />).toJSON();
    expect(tree).toMatchSnapshot();
  });
  test("snapshot3", () => {
    const tree = renderer.create(<Story3 />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
