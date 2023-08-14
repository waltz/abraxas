import test from 'ava';

import { Message } from './message';

test("can find media", t => {
  const davidBowie = "https://www.youtube.com/watch?v=3qrOvBuWJ-c";
  const message = new Message(davidBowie);
  t.is(message.hasMedia, true);
});