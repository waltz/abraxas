import * as dotenv from 'dotenv';

import { App } from "@slack/bolt";
import { GenericMessageEvent } from "@slack/bolt";

import { Message } from "./message";

dotenv.config();

const port = process.env.PORT || 3000;

const app = new App({
  token: process.env.SLACK_BOT_TOKEN,
  signingSecret: process.env.SLACK_SIGNING_SECRET,
});

(async () => {
  await app.start(port);

  app.event("message", async ({ payload, client, logger }) => {
    const event = payload as GenericMessageEvent;
    console.log("Got event", event);

    if (event.text) {
      const message = new Message(event.text);

      if (message.hasMedia) {
        console.log("Message had media. Reacting.");
        const result = await client.reactions.add({ channel: event.channel, timestamp: event.event_ts, name: "breedtv" });
        console.log("Reaction response: ", result);
      }
    } else {
      console.log("Message text was empty");
    }
  });
})();