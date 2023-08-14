export class Message {
  message: string;
  // urlRegex = new RegExp("https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]*\.[a-zA-Z0-9()]*\b([-a-zA-Z0-9()@:%_\+.~#?&=]*)", "g");
  urlRegex = new RegExp("youtube.com\/watch", "g");


  constructor(message: string) {
    this.message = message;
  }

  get hasMedia(): boolean {
    if (!this.message) {
      return false;
    }

    console.log('regex', this.urlRegex);

    const matches = [...this.message.matchAll(this.urlRegex)];
    console.log('Saw message matches', this.message, matches);

    return matches.length > 0;
  }
}