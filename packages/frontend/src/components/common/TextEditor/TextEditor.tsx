import TextField, { TextFieldProps } from "@mui/material/TextField";
import React from "react";

type TextEditorProps = TextFieldProps & { readonly maxlength?: number };

/**
 *
 */
export const TextEditor: React.FC<TextEditorProps> = ({ maxlength, inputProps, ...others }) => {
  return <TextField {...others} inputProps={{ ...inputProps, maxlength }} />;
};
