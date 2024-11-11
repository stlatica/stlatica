import { TextInput, type TextInputProps } from "@mantine/core";
import type React from "react";

type TextEditorProps = TextInputProps;

/**
 *
 */
export const TextEditor: React.FC<TextEditorProps> = ({ ...others }) => {
  return <TextInput {...others} />;
};
