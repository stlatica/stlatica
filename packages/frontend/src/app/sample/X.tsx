import { Card, IconButton, Stack, Typography } from "@mui/material";
import React from "react";
import StarBorderIcon from "@mui/icons-material/StarBorder";
import RecyclingIcon from "@mui/icons-material/Recycling";
import ReplyIcon from "@mui/icons-material/Reply";

export const X = () => {
  return (
    <Card style={{ backgroundColor: "skyblue" }} elevation={8}>
      <Stack>
        <Typography>ユーザー名</Typography>
        <Typography>ここに書き込みを表示する</Typography>
        <Stack direction="row">
          <IconButton>
            <ReplyIcon />
          </IconButton>
          <IconButton>
            <RecyclingIcon />
          </IconButton>
          <IconButton>
            <StarBorderIcon />
          </IconButton>
        </Stack>
      </Stack>
    </Card>
  );
};
