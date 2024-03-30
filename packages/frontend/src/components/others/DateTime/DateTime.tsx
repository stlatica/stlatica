import React from "react";

type DateTimeProps = {
  readonly date: string;
  readonly formatDate: (date: Date) => string;
};

/**
 *
 */
export const DateTime: React.FC<DateTimeProps> = ({ date, formatDate }) => {
  const dateObj = new Date(date);
  return <div>{formatDate(dateObj)}</div>;
};
