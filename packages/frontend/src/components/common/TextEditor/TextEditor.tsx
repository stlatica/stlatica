"use client";

import React, { useRef } from "react";

import { CreateID } from "@/features/acl/CreateID";

type TextEditorProps = {
  readonly defaultValue: string;
  readonly maxlength: number;
  readonly title: string;
};

/**
 *
 */
export const TextEditor: React.FC<TextEditorProps> = ({ defaultValue, maxlength, title }) => {
  const componentId = useRef(CreateID());
  return (
    <div className="relative h-full">
      <textarea
        className="peer block h-full w-full appearance-none rounded-lg border border-white bg-transparent p-3 px-2.5 pb-2.5 pt-4 text-sm text-gray-800 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
        defaultValue={defaultValue}
        id={componentId.current}
        maxLength={maxlength}
        placeholder=" "
      />
      <label
        className="absolute start-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 bg-white px-2 text-sm text-gray-500 duration-300 peer-placeholder-shown:top-6 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 peer-focus:text-blue-600 rtl:peer-focus:left-auto rtl:peer-focus:translate-x-1/4 dark:bg-gray-800 dark:text-gray-400 peer-focus:dark:text-blue-500"
        htmlFor={componentId.current}
      >
        {title}
      </label>
    </div>
  );
};
