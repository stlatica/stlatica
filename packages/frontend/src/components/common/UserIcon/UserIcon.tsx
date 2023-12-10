import React from "react";

type UserIconProps = {
  // children: React.ReactNode;
  /**
   * bese64
   */
  readonly iconImage: string;
};

/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = (props) => {
  const { iconImage } = props;
  return <img alt="icon" className="aspect-square h-full w-full rounded-full" src={iconImage} />;
};
