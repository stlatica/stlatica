import React from "react";

type UserIconProps = {
  // children: React.ReactNode;
  /**
   * bese64
   */
  readonly iconImage: string;
};

const cropStyle = {
  borderRadius: "50%",
};
/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = (props) => {
  const { iconImage } = props;
  return (
    <div>
      <img alt="icon" src={iconImage} style={cropStyle} />
    </div>
  );
};
