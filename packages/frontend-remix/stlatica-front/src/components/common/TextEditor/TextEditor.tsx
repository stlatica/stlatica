import { TextInput, TextInputProps } from "@mantine/core";
import React from "react";

type TextEditorProps = TextInputProps;

/**
 *
 */
export const TextEditor: React.FC<TextEditorProps> = ({ ...others }) => {
  return <TextInput {...others} />;
};
