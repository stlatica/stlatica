import React from "react";

type UserIconProps = {
  // children: React.ReactNode;

  /**
   * bese64
   */
  readonly iconImage: string;

  readonly alt?: string;
};

/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = ({ iconImage, alt = "user icon image." }) => {
  return <img alt={alt} className="aspect-square size-full rounded-full" src={iconImage} />;
};
