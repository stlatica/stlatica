import { Button, Card, IconButton, Stack, Typography } from "@mui/material";
import React from "react";
import StarBorderIcon from "@mui/icons-material/StarBorder";
import RecyclingIcon from "@mui/icons-material/Recycling";
import ReplyIcon from "@mui/icons-material/Reply";

const X = () => {
  return (
    <Card style={{ backgroundColor: "skyblue" }} elevation={8}>
      <Stack>
        <Typography>ユーザー名</Typography>
        <Typography>ここに書き込みを表示する</Typography>
        <Stack direction={"row"}>
          <IconButton>
            <ReplyIcon></ReplyIcon>
          </IconButton>
          <IconButton>
            <RecyclingIcon></RecyclingIcon>
          </IconButton>
          <IconButton>
            <StarBorderIcon></StarBorderIcon>
          </IconButton>
        </Stack>
      </Stack>
    </Card>
  );
};

const TimeLine: React.FC = () => {
  const ar = new Array(100).fill(0);

  const list = ar.map((x, index) => {
    return (
      <div style={{ marginBottom: "0.5em" }}>
        <X></X>
      </div>
    );
  });

  return <div>{list}</div>;
};

export default function Home() {
  return (
    <main>
      <TimeLine></TimeLine>
    </main>
  );
}
