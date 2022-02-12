/* eslint-disable max-classes-per-file */
import BaseError from '../base-error';

export class HttpError extends BaseError {
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  constructor(public response: Response) {
    super('HttpError', 'Unexpected error occurred');
  }
}

export class UnprocessableImageError extends BaseError {
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  constructor(public response: Response) {
    super('UnprocessableImageError', 'Unexpected error occurred');
  }
}

class HttpService {
  async get<T>(path: string) {
    return this.request<T>('get', path);
  }

  async post<T>(path: string, body: FormData) {
    return this.request<T>('post', path, body);
  }

  async delete<T>(path: string, body: FormData) {
    return this.request<T>('delete', path, body);
  }

  async put<T>(path: string, body: FormData) {
    return this.request<T>('put', path, body);
  }

  /* eslint-disable-next-line class-methods-use-this */
  private async request<T>(method: string, path: string, body?: FormData) {
    const response = await fetch(`${process.env.REACT_APP_PUBLIC_SERVER_URL}${path}`, {
      method,
      body,
    });

    if (response.status === 422) {
      throw new UnprocessableImageError(response);
    }

    if (response.status < 200 || response.status >= 300) {
      throw new HttpError(response);
    }

    const responseBody: T = await response.json();

    return responseBody;
  }
}

export const httpService = new HttpService();
