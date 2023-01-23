import express from "express";

const port = process.env.PORT || 3000;
const app = express();

console.log(`Server started on port=${port}`);

app.get("/", (_, res) => {
  res.send("abraxas is alive");
});

app.listen(port);