const API_URL = 'http://localhost:8000';

export function post<T = any>(data: T): Promise<Response> {
  return fetch(API_URL + '/api/auth', {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });
}
