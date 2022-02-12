/* eslint no-unused-vars: 0 */

export default class BaseError {
  public stack?: string;

  constructor(public name: string, public message: string) {
    this.stack = new Error().stack;
  }
}
