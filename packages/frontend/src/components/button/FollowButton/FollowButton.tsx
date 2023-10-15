// /* istanbul ignore file */
// import { styled } from "@mui/material/styles";

// /**
//  *
//  */
// export const FollowButton = styled("div")({});

import React, { useState, useEffect } from "react";
import Button from "@mui/material/Button";

type FollowButtonProps = {
  initialFollowState?: boolean;
};

const FollowButton: React.FC<FollowButtonProps> = ({ initialFollowState }) => {
  const [isFollow, setIsFollow] = useState(initialFollowState);

  const handleClick = () => {
    setIsFollow(!isFollow);
  };

  useEffect(() => {
    console.log(isFollow ? "Currently Following" : "Not Following");
  }, [isFollow]);

  return (
    <Button variant="contained" color={isFollow ? "secondary" : "primary"} onClick={handleClick}>
      {isFollow ? "フォロー中" : "フォロー"}
    </Button>
  );
};

export default FollowButton;
