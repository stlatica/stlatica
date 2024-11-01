import type React from "react";

import { userIcon } from "./user-icon.css";

type UserIconProps = {
  /**
   * base64
   */
  readonly iconImage: string;

  readonly alt?: string;
};

/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = ({ iconImage, alt = "user icon" }) => {
  return <img alt={alt} class={userIcon} src={iconImage} />;
};
