import React from "react";
import renderer from "react-test-renderer";
import { Story1 } from "./FollowButton.stories";

test("snapshot", () => {
  const tree = renderer.create(<Story1 />).toJSON();
  expect(tree).toMatchSnapshot();
});
