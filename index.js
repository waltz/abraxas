const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (_, res) => {
  console.log('Serving request');
  res.send('Thank you for talking to Abraxas');
});

app.listen(port, () => {
  console.log(`Abraxas is listening on port ${port}`);
});
